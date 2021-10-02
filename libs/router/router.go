package router

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"

	"log"
	"net/http"

	"github.com/Ireoo/API-Core/libs/conf"
	"github.com/Ireoo/API-Core/libs/json"
	mongo "github.com/Ireoo/API-Core/libs/mongodb"
	"github.com/gin-gonic/gin"
	"github.com/gookit/color"

	"go.mongodb.org/mongo-driver/bson"

	simplejson "github.com/bitly/go-simplejson"
	// gojsonq "github.com/thedevsaddam/gojsonq"
)

//func Test(c gin.Context, secret) {
//	Input := new(conf.Input) //new(conf.Input)
//	if error := c.Bind(Input); error != nil {
//		miss(error)
//	} else {
//		trace(Input)
//	}
//
//	buf := make([]byte, c.Request().ContentLength)
//	c.Request().Body.Read(buf)
//	trace(string(buf))
//
//	trace(`[INPUT][WHERE]: `)
//	trace(reflect.TypeOf(Input.Where))
//	trace(Input.Where)
//	mjson, _ := json.Marshal(Input.Where)
//
//	where := gojsonq.New().FromString(string(mjson))
//	trace(reflect.TypeOf(where))
//	trace(where)
//	id := gojsonq.New().FromString(string(mjson)).Find("_id")
//	trace(reflect.TypeOf(id))
//	trace(id)
//
//	value, _ := sjson.Set(string(mjson), "_id", id)
//	trace(value)
//	// Input.Where._id = bson.ObjectIdHex(string(id.([]byte)))
//	return
//}

func Table(c *gin.Context, secret string) {
	Input := new(conf.Input)

	buf := make([]byte, c.Request.ContentLength)
	_, _ = c.Request.Body.Read(buf)
	fmt.Println(string(buf))

	res, err := simplejson.NewJson(buf)
	if err != nil {
		fmt.Printf("1 - %v\n", err)
		return
	}

	Input.Where = ijson.Format(res.Get("where"))
	Input.Data = ijson.Format(res.Get("data"))
	Input.App, _ = res.Get("app").String()
	// trace(where)
	// trace(data)
	// trace(_app)

	Input.Table = c.Param("table")
	Input.Mode = c.Param("mode")
	Input.Auth = c.Request.Header.Get("Authorization")

	trace(Input)

	if Input.Auth == "" {
		c.String(http.StatusNonAuthoritativeInfo, "Not Authorization!")
	}

	app := "api"
	if Input.Auth != secret {
		//var result bson.M
		AppInfo := new(conf.AppInfo)
		error := mongo.FindOne(app, "apps", bson.M{"secret": Input.Auth}, bson.M{}, &AppInfo)
		if error != nil {
			miss(error)
			c.String(http.StatusNonAuthoritativeInfo, "The authorization verification information does not exist. Please verify.")
		}
		app = AppInfo.Id
		// mjson, _ := json.Marshal(AppInfo)
		// id := gojsonq.New().FromString(string(mjson)).Find("Id")
		// app = string(id.([]byte)) //md5V(hex.EncodeToString([]byte(AppInfo.Id)))
		//fmt.Println(app)
	}

	where := Input.Where
	data := Input.Data

	switch Input.Mode {
	case "once":
		var result bson.M
		error := mongo.FindOne(app, Input.Table, where, bson.M{}, &result)
		output(c, result, error)

	case "findOne":
		var result bson.M
		error := mongo.FindOne(app, Input.Table, where, bson.M{}, &result)
		output(c, result, error)

	case "findAll":
		result, error := mongo.FindAll(app, Input.Table, where, bson.M{})
		output(c, result, error)

	case "find":
		result, error := mongo.FindPage(app, Input.Table, Input.Other.Page, Input.Other.Limit, where, bson.M{})
		output(c, result, error)

	case "insert":
		insert, error := mongo.Insert(app, Input.Table, data)
		if error != nil {
			output(c, nil, error)
		}
		var result bson.M
		miss(bson.M{"_id": insert.InsertedID})
		error = mongo.FindOne(app, Input.Table, bson.M{"_id": insert.InsertedID}, bson.M{}, &result)
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
	if error := c.Bind(Input); error != nil {
		miss(error)
	} else {
		trace(Input)
	}

	Input.Table = c.Param("table")
	Input.Mode = c.Param("mode")
	Input.Auth = c.Request.Header.Get("Authorization")

	if Input.Auth == "" {
		c.String(http.StatusNonAuthoritativeInfo, "Not Authorization!")
	}

	app := "api"
	if Input.Auth != secret {
		//var result bson.M
		AppInfo := new(conf.AppInfo)
		error := mongo.FindOne(app, "apps", bson.M{"secret": Input.Auth}, bson.M{}, &AppInfo)
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

func md5V(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
