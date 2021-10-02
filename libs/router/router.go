package router

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"reflect"

	"log"
	"net/http"

	"github.com/Ireoo/API-Core/libs/conf"
	"github.com/Ireoo/API-Core/libs/mongodb"
	"github.com/gookit/color"
	"github.com/labstack/echo/v4"
	"gopkg.in/mgo.v2/bson"

	gojsonq "github.com/thedevsaddam/gojsonq/v2"
	"github.com/tidwall/sjson"
)

func Test(c echo.Context, secret, auth string) (err error) {
	Input := new(conf.Input) //new(conf.Input)
	if error := c.Bind(Input); error != nil {
		miss(error)
	} else {
		trace(Input)
	}

	buf := make([]byte, c.Request().ContentLength)
	c.Request().Body.Read(buf)
	trace(string(buf))

	trace(`[INPUT][WHERE]: `)
	trace(reflect.TypeOf(Input.Where))
	trace(Input.Where)
	mjson, _ := json.Marshal(Input.Where)

	where := gojsonq.New().FromString(string(mjson))
	trace(reflect.TypeOf(where))
	trace(where)
	id := gojsonq.New().FromString(string(mjson)).Find("_id")
	trace(reflect.TypeOf(id))
	trace(id)

	value, _ := sjson.Set(string(mjson), "_id", id)
	trace(value)
	// Input.Where._id = bson.ObjectIdHex(string(id.([]byte)))
	return
}

func Table(c echo.Context, secret, auth string) (err error) {
	Input := new(conf.Input)
	if error := c.Bind(Input); error != nil {
		miss(error)
	} else {
		trace(Input)
	}

	Input.Table = c.Param("table")
	Input.Mode = c.Param("mode")
	Input.Auth = auth

	if Input.Auth == "" {
		return c.String(http.StatusNonAuthoritativeInfo, "Not Authorization!")
	}

	app := "api"
	if Input.Auth != secret {
		//var result bson.M
		AppInfo := new(conf.AppInfo)
		error := mongo.FindOne(app, "apps", bson.M{"secret": Input.Auth}, bson.M{}, &AppInfo)
		if error != nil {
			miss(error)
			return c.String(http.StatusNonAuthoritativeInfo, "The authorization verification information does not exist. Please verify.")
		}
		mjson, _ := json.Marshal(AppInfo)
		id := gojsonq.New().FromString(string(mjson)).Find("Id")
		app = string(id.([]byte)) //md5V(hex.EncodeToString([]byte(AppInfo.Id)))
		//fmt.Println(app)
	}

	where := Input.Where
	data := Input.Data

	switch Input.Mode {
	case "once":
		var result bson.M
		error := mongo.FindOne(app, Input.Table, where, bson.M{}, &result)
		return output(c, result, error)

	case "findOne":
		var result bson.M
		error := mongo.FindOne(app, Input.Table, where, bson.M{}, &result)
		return output(c, result, error)

	case "findAll":
		result, error := mongo.FindAll(app, Input.Table, where, bson.M{})
		return output(c, result, error)

	case "find":
		result, error := mongo.FindPage(app, Input.Table, Input.Other.Page, Input.Other.Limit, where, bson.M{})
		return output(c, result, error)

	case "insert":
		insert, error := mongo.Insert(app, Input.Table, data)
		return output(c, insert, error)

	case "update":
		error := mongo.Update(app, Input.Table, bson.M{"$set": data}, where)
		return output(c, data, error)

	case "updateAll":
		error := mongo.UpdateAll(app, Input.Table, bson.M{"$set": data}, where)
		return output(c, data, error)

	case "upsert":
		error := mongo.Upsert(app, Input.Table, bson.M{"$set": data}, where)
		return output(c, where, error)

	case "remove":
		error := mongo.Remove(app, Input.Table, where)
		return output(c, where, error)

	case "removeAll":
		error := mongo.RemoveAll(app, Input.Table, where)
		return output(c, where, error)

	case "count":
		count, error := mongo.Count(app, Input.Table, where)
		return output(c, count, error)
		// return c.String(http.StatusOK, strconv.Itoa(int(count)))

	case "isEmpty":
		var r string
		if ex := mongo.IsEmpty(app, Input.Table); ex {
			r = "true"
		} else {
			r = "false"
		}
		trace(r)
		return c.String(http.StatusOK, r)

	case "listCollections":
		result, error := mongo.CollectionNames(Input.App)
		return output(c, result, error)

	default:
		return c.String(http.StatusNotFound, "不存在的操作模式："+Input.Mode)
	}
}

func trace(message interface{}) {
	log.Println(color.FgGreen.Render(message))
}

func miss(message interface{}) {
	log.Println(color.FgRed.Render(message))
}

func output(c echo.Context, r interface{}, e error) (err error) {
	if e != nil {
		miss(e)
		out := &conf.Result{
			Success: false,
			Data:    e.Error(),
		}
		return c.JSON(http.StatusOK, out)
	}
	trace(r)
	out := &conf.Result{
		Success: true,
		Data:    r,
	}
	return c.JSON(http.StatusOK, out)
}

func md5V(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
