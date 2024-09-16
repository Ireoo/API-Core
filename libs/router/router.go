package router

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gookit/color"
	"io"
	"net/url"

	"github.com/Ireoo/API-Core/libs/conf"
	"github.com/Ireoo/API-Core/libs/debug"
	iJson "github.com/Ireoo/API-Core/libs/json"
	mongo "github.com/Ireoo/API-Core/libs/mongodb"
	simpleJson "github.com/bitly/go-simplejson"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

func parseInput(c *gin.Context) (*conf.Input, *conf.Other, error) {
	Input := new(conf.Input)
	buf := make([]byte, c.Request.ContentLength)
	if c.Request.Body == nil {
		return nil, nil, errors.New("request body is nil")
	}
	_, err := c.Request.Body.Read(buf)
	if err != nil && err != io.EOF {
		return nil, nil, err
	}

	res, err := simpleJson.NewJson(buf)
	if err != nil {
		return nil, nil, err
	}

	Input.Where = iJson.Format(res.Get("where"))
	Input.Data = iJson.Format(res.Get("data"))
	if andArray, err := res.Get("where").Get("$and").Array(); err == nil {
		if whereMap, ok := Input.Where.(map[string]interface{}); ok {
			whereMap["$and"] = andArray
			Input.Where = whereMap
		}
	}
	if orArray, err := res.Get("where").Get("$or").Array(); err == nil {
		if whereMap, ok := Input.Where.(map[string]interface{}); ok {
			whereMap["$or"] = orArray
			Input.Where = whereMap
		}
	}

	other := new(conf.Other)

	limit, err := res.Get("other").Get("limit").Int64()
	if err != nil || limit == 0 {
		limit = 20
	}
	other.Limit = limit
	page, err := res.Get("other").Get("page").Int64()
	if err != nil {
		page = 0
	}
	other.Page = page * limit
	// skip := page * limit // <- 删除未使用的变量
	other.Show = iJson.Format(res.Get("other").Get("show"))
	other.Distinct = iJson.Format(res.Get("other").Get("distinct"))
	other.Sort = iJson.Format(res.Get("other").Get("sort"))
	other.Indexes, _ = res.Get("other").Get("indexes").StringArray()
	other.Upsert, _ = res.Get("other").Get("upsert").Bool()
	other.Multi, _ = res.Get("other").Get("multi").Bool()

	Input.App, _ = res.Get("app").String()
	Input.Table = c.Param("table")
	Input.Mode = c.Param("mode")
	Input.Auth = c.Request.Header.Get("Authorization")
	Input.Other = other

	return Input, other, nil
}

func handleAuth(Input *conf.Input, secret string, other *conf.Other) (string, string, error) {
	app := "api"
	user := ""
	if Input.Auth != secret {
		AppInfo := new(conf.AppInfo)
		err := mongo.FindOne(app, "apps", bson.M{"secret": Input.Auth}, other, &AppInfo)
		if err != nil {
			return "", "", errors.New("the authorization verification information does not exist. please verify")
		}
		app = AppInfo.Id
		user = AppInfo.Uuid
	}
	return app, user, nil
}

func Table(c *gin.Context, secret string, Debug bool) {
	debug.SetDebug(Debug)

	Input, other, err := parseInput(c)
	if err != nil {
		output(c, nil, err)
		return
	}

	if debug.GetDebug() {
		jsonStr, _ := json.Marshal(Input)
		debug.Info("[INPUT]" + " " + string(jsonStr))
	}

	if Input.Auth == "" {
		output(c, nil, errors.New("no authorization"))
		return
	}

	app, user, err := handleAuth(Input, secret, other)
	if err != nil {
		output(c, nil, err)
		return
	}

	where := Input.Where
	data := Input.Data

	switch Input.Mode {
	case "once", "findOne":
		var result bson.M
		err := mongo.FindOne(app, Input.Table, where, other, &result)
		output(c, result, err)

	case "findAll":
		result, err := mongo.FindAll(app, Input.Table, where, other)
		output(c, result, err)

	case "find":
		result, err := mongo.FindPage(app, Input.Table, other, where)
		output(c, result, err)

	case "insert":
		insert, err := mongo.Insert(app, Input.Table, data)
		if err != nil {
			output(c, nil, err)
			return
		}
		var result bson.M
		debug.Info(bson.M{"_id": insert.InsertedID})
		err = mongo.FindOne(app, Input.Table, bson.M{"_id": insert.InsertedID}, other, &result)
		output(c, result, err)

	case "update":
		err := mongo.Update(app, Input.Table, where, data, other)
		if err != nil {
			output(c, nil, err)
			return
		}
		if other.Multi {
			result, err := mongo.FindAll(app, Input.Table, where, other)
			output(c, result, err)
		} else {
			var result bson.M
			err = mongo.FindOne(app, Input.Table, where, other, &result)
			output(c, result, err)
		}

	case "remove":
		err := mongo.Remove(app, Input.Table, where)
		output(c, where, err)

	case "removeAll":
		err := mongo.RemoveAll(app, Input.Table, where)
		output(c, where, err)

	case "count":
		count, err := mongo.Count(app, Input.Table, where)
		output(c, count, err)

	case "isEmpty":
		var r string
		if ex := mongo.IsEmpty(app, Input.Table); ex {
			r = "true"
		} else {
			r = "false"
		}
		output(c, r, nil)

	case "setIndex":
		var r string
		if ex := mongo.EnsureIndex(app, Input.Table, other.Indexes); ex == nil {
			r = "true"
		} else {
			r = "false"
		}
		output(c, r, nil)

	case "getIndexes":
		r := mongo.Indexes(app, Input.Table)
		output(c, r, nil)

	case "listCollections", "drop":
		if Input.App == "" {
			output(c, nil, errors.New("no app id"))
			return
		}
		_app := app
		if _app == "api" {
			_app = Input.App
		}

		appInfo := new(conf.AppInfo)
		_id, err := primitive.ObjectIDFromHex(_app)
		if err != nil {
			output(c, nil, err)
			return
		}
		err = mongo.FindOne(app, "users", bson.M{"_id": _id}, other, &appInfo)
		if err != nil {
			output(c, nil, err)
			return
		}
		if appInfo.Uuid != user {
			output(c, nil, errors.New("unauthorized operation"))
			return
		}
		if Input.Mode == "listCollections" {
			result, err := mongo.CollectionNames(_app)
			output(c, result, err)
		} else {
			err := mongo.DropDatabase(_app)
			output(c, nil, err)
		}

	default:
		output(c, nil, fmt.Errorf("operating mode in existence: %v", Input.Mode))
	}
}

func TableGet(c *gin.Context, secret string, Debug bool) {
	debug.SetDebug(Debug)

	Input := new(conf.Input)
	other := new(conf.Other)

	Input.Table = c.Param("table")
	Input.Mode = c.Param("mode")

	decoded, _ := url.QueryUnescape(c.Query("where"))
	_ = bson.UnmarshalExtJSON([]byte(decoded), true, &Input.Where)
	decoded, _ = url.QueryUnescape(c.Query("data"))
	_ = bson.UnmarshalExtJSON([]byte(decoded), true, &Input.Data)
	decoded, _ = url.QueryUnescape(c.Query("other"))
	_ = bson.UnmarshalExtJSON([]byte(decoded), true, &other)
	decoded, _ = url.QueryUnescape(c.Query("app"))
	_ = bson.UnmarshalExtJSON([]byte(decoded), true, &Input.App)
	decoded, _ = url.QueryUnescape(c.Query("auth"))
	if decoded != "" {
		Input.Auth = decoded
	} else {
		Input.Auth = c.Request.Header.Get("Authorization")
	}
	Input.Other = other

	if debug.GetDebug() {
		jsonStr, _ := json.Marshal(Input)
		debug.Info("[INPUT]" + " " + string(jsonStr))
	}

	if Input.Auth == "" {
		output(c, nil, fmt.Errorf("no authorization"))
		return
	}

	app, user, err := handleAuth(Input, secret, other)
	if err != nil {
		output(c, nil, err)
		return
	}

	where := Input.Where

	switch Input.Mode {
	case "once", "findOne":
		var result bson.M
		err := mongo.FindOne(app, Input.Table, where, other, &result)
		output(c, result, err)

	case "findAll":
		result, err := mongo.FindAll(app, Input.Table, where, other)
		output(c, result, err)

	case "find":
		result, err := mongo.FindPage(app, Input.Table, other, where)
		output(c, result, err)

	case "count":
		count, err := mongo.Count(app, Input.Table, where)
		output(c, count, err)

	case "isEmpty":
		var r string
		if ex := mongo.IsEmpty(app, Input.Table); ex {
			r = "true"
		} else {
			r = "false"
		}
		output(c, r, nil)

	case "getIndexes":
		r := mongo.Indexes(app, Input.Table)
		output(c, r, nil)

	case "listCollections":
		if Input.App == "" {
			output(c, nil, errors.New("no app id"))
			return
		}
		_app := app
		if _app == "api" {
			_app = Input.App
		}

		appInfo := new(conf.AppInfo)
		_id, err := primitive.ObjectIDFromHex(_app)
		if err != nil {
			output(c, nil, err)
			return
		}
		err = mongo.FindOne(app, "users", bson.M{"_id": _id}, other, &appInfo)
		if err != nil {
			output(c, nil, err)
			return
		}
		if appInfo.Uuid != user {
			output(c, nil, errors.New("unauthorized operation"))
			return
		}
		result, err := mongo.CollectionNames(_app)
		output(c, result, err)

	default:
		output(c, nil, fmt.Errorf("operating mode in existence: %v", Input.Mode))
	}
}

func output(c *gin.Context, r interface{}, e error) {
	if e != nil {
		debug.Error("[ERROR]" + " " + color.FgDefault.Render(e))
		out := &conf.Result{
			Success: false,
			Data:    e.Error(),
		}
		c.JSON(http.StatusOK, out)
		return
	}
	if r != nil {
		if debug.GetDebug() {
			jsonStr, _ := json.Marshal(r)
			debug.Trace("[OUTPUT]" + " " + string(jsonStr))
		}
		out := &conf.Result{
			Success: true,
			Data:    r,
		}
		c.JSON(http.StatusOK, out)
	}
}
