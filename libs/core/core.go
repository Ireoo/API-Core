package core

import (
	"libs/mongo"
	"encoding/hex"
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

func Core(c echo.Context) (err error) {
	Input := new(input)
	if error := c.Bind(Input); error != nil {
		e.Logger.Print(error)
	} else {
		e.Logger.Print(Input)
	}

	Input.Table = c.Param("table")
	Input.Mode = c.Param("mode")
	Input.Auth = auth

	if Input.Auth == "" {
		return c.String(http.StatusNonAuthoritativeInfo, "Not Authorization!")
	}

	app := "api"
	if Input.Auth != "94f3eee0-218f-41fc-9318-94cf5430fc7f" {
		//var result bson.M
		AppInfo := new(appInfo)
		error := mongo.FindOne(app, "apps", bson.M{"secret": Input.Auth}, bson.M{}, &AppInfo)
		if error != nil {
			e.Logger.Print(error)
			return c.String(http.StatusNonAuthoritativeInfo, "The authorization verification information does not exist. Please verify.")
		}
		app = hex.EncodeToString([]byte(AppInfo.Id))
		//fmt.Println(app)
	}

	switch Input.Mode {
	case "findOne":
		var result bson.M
		error := mongo.FindOne(app, Input.Table, Input.Where, bson.M{}, &result)
		if error != nil {
			e.Logger.Print(error)
			return c.String(http.StatusNotFound, error.Error())
		}
		e.Logger.Print(result)
		return c.JSON(http.StatusOK, result)

	case "findAll":
		var result []bson.M
		error := mongo.FindAll(app, Input.Table, Input.Where, bson.M{}, &result)
		if error != nil {
			fmt.Println(error)
			return c.String(http.StatusNotFound, error.Error())
		}
		e.Logger.Print(result)
		return c.JSON(http.StatusOK, result)

	case "findPage":
		var result []bson.M
		error := mongo.FindPage(app, Input.Table, Input.Other.page, Input.Other.limit, Input.Where, bson.M{}, &result)
		if error != nil {
			e.Logger.Print(error)
			return c.String(http.StatusNotFound, error.Error())
		}
		e.Logger.Print(result)
		return c.JSON(http.StatusOK, result)

	case "insert":
		error := mongo.Insert(app, Input.Table, Input.Data)
		if error != nil {
			e.Logger.Print(error)
			return c.String(http.StatusNotFound, error.Error())
		}
		return c.JSON(http.StatusOK, Input.Data)

	case "update":
		error := mongo.Update(app, Input.Table, bson.M{"$set": Input.Data}, Input.Where)
		if error != nil {
			e.Logger.Print(error)
			return c.String(http.StatusNotFound, error.Error())
		}
		return c.JSON(http.StatusOK, Input.Data)

	case "updateAll":
		error := mongo.UpdateAll(app, Input.Table, bson.M{"$set": Input.Data}, Input.Where)
		if error != nil {
			e.Logger.Print(error)
			return c.String(http.StatusNotFound, error.Error())
		}
		return c.JSON(http.StatusOK, Input.Data)

	case "upsert":
		error := mongo.Upsert(app, Input.Table, bson.M{"$set": Input.Data}, Input.Where)
		if error != nil {
			e.Logger.Print(error)
			return c.String(http.StatusNotFound, error.Error())
		}
		return c.JSON(http.StatusOK, Input.Where)

	case "remove":
		error := mongo.Remove(app, Input.Table, Input.Where)
		if error != nil {
			e.Logger.Print(error)
			return c.String(http.StatusNotFound, error.Error())
		}
		return c.JSON(http.StatusOK, Input.Where)

	case "removeAll":
		error := mongo.RemoveAll(app, Input.Table, Input.Where)
		if error != nil {
			e.Logger.Print(error)
			return c.String(http.StatusNotFound, error.Error())
		}
		return c.JSON(http.StatusOK, Input.Where)

	case "count":
		count, error := mongo.Count(app, Input.Table, Input.Where)
		if error != nil {
			e.Logger.Print(error)
			return c.String(http.StatusNotFound, error.Error())
		}
		e.Logger.Print(count)
		return c.String(http.StatusOK, strconv.Itoa(count))

	case "isEmpty":
		var r string
		if ex := mongo.IsEmpty(app, Input.Table); ex {
			r = "true"
		} else {
			r = "false"
		}
		e.Logger.Print(r)
		return c.String(http.StatusOK, r)

	default:
		return c.String(http.StatusNotFound, "不存在的操作模式："+Input.Mode)
	}

}
