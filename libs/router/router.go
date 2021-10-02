package router

import (
	"fmt"
	iJson "github.com/Ireoo/API-Core/libs/json"

	"log"
	"net/http"

	"github.com/Ireoo/API-Core/libs/conf"
	mongo "github.com/Ireoo/API-Core/libs/mongodb"
	"github.com/gin-gonic/gin"
	"github.com/gookit/color"

	"go.mongodb.org/mongo-driver/bson"

	simplejson "github.com/bitly/go-simplejson"
)

func Table(c *gin.Context, secret string) {
	Input := new(conf.Input)

	buf := make([]byte, c.Request.ContentLength)
	_, _ = c.Request.Body.Read(buf)
	fmt.Println(string(buf))

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

	// trace(Input)

	if Input.Auth == "" {
		c.String(http.StatusNonAuthoritativeInfo, "Not Authorization!")
	}

	app := "api"
	if Input.Auth != secret {
		AppInfo := new(conf.AppInfo)
		error := mongo.FindOne(app, "apps", bson.M{"secret": Input.Auth}, other, &AppInfo)
		if error != nil {
			miss(error)
			c.String(http.StatusNonAuthoritativeInfo, "The authorization verification information does not exist. Please verify.")
		}
		app = AppInfo.Id
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
		}
		var result bson.M
		miss(bson.M{"_id": insert.InsertedID})
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
		trace(r)
		c.String(http.StatusOK, r)

	case "listCollections":
		result, error := mongo.CollectionNames(Input.App)
		output(c, result, error)

	default:
		c.String(http.StatusNotFound, "不存在的操作模式："+Input.Mode)
	}
}

func Mode(c *gin.Context, secret string) {
	Input := new(conf.Input)

	buf := make([]byte, c.Request.ContentLength)
	_, _ = c.Request.Body.Read(buf)
	fmt.Println(string(buf))

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

	// trace(Input)

	if Input.Auth == "" {
		c.String(http.StatusNonAuthoritativeInfo, "Not Authorization!")
	}

	app := "api"
	if Input.Auth != secret {
		//var result bson.M
		AppInfo := new(conf.AppInfo)
		error := mongo.FindOne(app, "apps", bson.M{"secret": Input.Auth}, other, &AppInfo)
		if error != nil {
			miss(error)
			c.String(http.StatusNonAuthoritativeInfo, "The authorization verification information does not exist. Please verify.")
		}
		app = AppInfo.Id
		//fmt.Println(app)
	}

	switch Input.Mode {
	case "collectionNames":
		names, error := mongo.CollectionNames(app)
		if error != nil {
			miss(error)
			c.String(http.StatusNotFound, error.Error())
		}
		trace(names)
		c.JSON(http.StatusOK, names)

	default:
		c.String(http.StatusNotFound, "不存在的操作模式："+Input.Mode)
	}
}

func trace(message interface{}) {
	log.Println(color.FgGreen.Render(message))
}

func miss(message interface{}) {
	log.Println(color.FgRed.Render(message))
}

func output(c *gin.Context, r interface{}, e error) {
	if e != nil {
		miss(e)
		out := &conf.Result{
			Success: false,
			Data:    e.Error(),
		}
		c.JSON(http.StatusOK, out)
	}
	if r != nil {
		trace(r)
		out := &conf.Result{
			Success: true,
			Data:    r,
		}
		c.JSON(http.StatusOK, out)
	}
}
