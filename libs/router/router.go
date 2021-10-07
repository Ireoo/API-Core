package router

import (
	"errors"
	"fmt"

	iJson "github.com/Ireoo/API-Core/libs/json"

	"net/http"

	"github.com/Ireoo/API-Core/libs/conf"
	"github.com/Ireoo/API-Core/libs/debug"
	mongo "github.com/Ireoo/API-Core/libs/mongodb"
	"github.com/gin-gonic/gin"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	simplejson "github.com/bitly/go-simplejson"
)

func Table(c *gin.Context, secret string, Debug bool) {

	debug.SetDebug(Debug)

	Input := new(conf.Input)

	buf := make([]byte, c.Request.ContentLength)
	_, _ = c.Request.Body.Read(buf)

	res, err := simplejson.NewJson(buf)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	Input.Where = iJson.Format(res.Get("where"))
	Input.Data = iJson.Format(res.Get("data"))

	other := new(conf.Other)

	limit, err := res.Get("other").Get("page").Int64()
	if err != nil {
		limit = 20
	}
	if limit == 0 {
		limit = 20
	}
	other.Limit = limit
	page, err := res.Get("other").Get("page").Int64()
	if err != nil {
		page = 0
	}
	other.Page = page * limit
	other.Show = iJson.Format(res.Get("other").Get("show"))
	other.Distinct = iJson.Format(res.Get("other").Get("distinct"))
	other.Sort = iJson.Format(res.Get("other").Get("sort"))

	Input.App, _ = res.Get("app").String()

	Input.Table = c.Param("table")
	Input.Mode = c.Param("mode")
	Input.Auth = c.Request.Header.Get("Authorization")

	debug.Info("[INPUT]" + " " + string(buf))

	if Input.Auth == "" {
		c.String(http.StatusNonAuthoritativeInfo, "Not Authorization!")
		return
	}

	app := "api"
	user := ""
	if Input.Auth != secret {
		AppInfo := new(conf.AppInfo)
		error := mongo.FindOne(app, "apps", bson.M{"secret": Input.Auth}, other, &AppInfo)
		if error != nil {
			debug.Error(error)
			c.String(http.StatusNonAuthoritativeInfo, "The authorization verification information does not exist. Please verify.")
			return
		}
		app = AppInfo.Id
		user = AppInfo.Uuid
	}

	where := Input.Where
	data := Input.Data

	switch Input.Mode {
	case "once":
		var result bson.M
		error := mongo.FindOne(app, Input.Table, where, other, &result)
		output(c, result, error)

	case "findOne":
		var result bson.M
		error := mongo.FindOne(app, Input.Table, where, other, &result)
		output(c, result, error)

	case "findAll":
		result, error := mongo.FindAll(app, Input.Table, where, other)
		output(c, result, error)

	case "find":
		result, error := mongo.FindPage(app, Input.Table, other, where)
		output(c, result, error)

	case "insert":
		insert, error := mongo.Insert(app, Input.Table, data)
		if error != nil {
			output(c, nil, error)
			return
		}
		var result bson.M
		debug.Info(bson.M{"_id": insert.InsertedID})
		error = mongo.FindOne(app, Input.Table, bson.M{"_id": insert.InsertedID}, other, &result)
		output(c, result, error)

	case "update":
		error := mongo.Update(app, Input.Table, bson.M{"$set": data}, where)
		output(c, data, error)

	case "updateAll":
		error := mongo.UpdateAll(app, Input.Table, bson.M{"$set": data}, where)
		output(c, data, error)

	case "upsert":
		error := mongo.Upsert(app, Input.Table, bson.M{"$set": data}, where)
		output(c, where, error)

	case "remove":
		error := mongo.Remove(app, Input.Table, where)
		output(c, where, error)

	case "removeAll":
		error := mongo.RemoveAll(app, Input.Table, where)
		output(c, where, error)

	case "count":
		count, error := mongo.Count(app, Input.Table, where)
		output(c, count, error)
		// return c.String(http.StatusOK, strconv.Itoa(int(count)))

	case "isEmpty":
		var r string
		if ex := mongo.IsEmpty(app, Input.Table); ex {
			r = "true"
		} else {
			r = "false"
		}
		debug.Trace(r)
		c.String(http.StatusOK, r)

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
		_id, error := primitive.ObjectIDFromHex(_app)
		if error != nil {
			output(c, nil, error)
			return
		}
		error = mongo.FindOne(app, "users", bson.M{"_id": _id}, other, &appInfo)
		if error != nil {
			output(c, nil, error)
			return
		}
		if appInfo.Uuid != user {
			output(c, nil, errors.New("unauthorized operation"))
			return
		}
		result, error := mongo.CollectionNames(_app)
		output(c, result, error)

	default:
		output(c, nil, fmt.Errorf("operating mode in existence: %v", Input.Mode))
	}
}

func output(c *gin.Context, r interface{}, e error) {
	if e != nil {
		debug.Error(e)
		out := &conf.Result{
			Success: false,
			Data:    e.Error(),
		}
		c.JSON(http.StatusOK, out)
	}
	if r != nil {
		debug.Trace(r)
		out := &conf.Result{
			Success: true,
			Data:    r,
		}
		c.JSON(http.StatusOK, out)
	}
}
