package main

import (
	"API-Core/libs/mongo"
	"flag"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"gopkg.in/mgo.v2/bson"
)

type input struct {
	Where bson.M `json:"where"`
	Data  bson.M `json:"data"`
	Other other  `json:"other"`
	Table string `json:"table"`
	Mode  string `json:"mode"`
	Auth  string `json:"auth"`
}

type other struct {
	page  int
	limit int
}

const version = "1.0.0"

var ver = flag.Bool("v", false, "版本信息")
var port = flag.String("p", "2019", "端口地址")

var auth = ""

func main() {
	flag.Parse()

	if *ver {
		fmt.Printf(`API-Core version: %s`, version)
		return
	}

	// var result bson.M
	// err := mongo.FindOne("api", "user", bson.M{"username": "18551410359"}, bson.M{"_id": 0}, &result)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(result)

	e := echo.New()
	e.Logger.SetLevel(log.DEBUG)

	// e.Logger.Print(os.Args)

	// 设置静态文件
	e.Static("/", "static")
	e.File("/favicon.ico", "static/favicon.ico")
	e.File("/", "static/index.html")
	e.File("/admin", "static/admin.html")

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	// 	e.GET("/", func(c echo.Context) error {
	// 		return c.HTML(http.StatusOK, `<h1 style="text-align: center;">欢迎使用 iData API 数据中心!</h1>
	// <h3 style="text-align: center;">Welcome to the iData API Data Center!</h3>`)
	// 	})

	// 程序核心部分
	e.POST("/:table/:mode", func(c echo.Context) (err error) {
		Input := new(input)
		if err := c.Bind(Input); err != nil {
			e.Logger.Print(err)
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
			app = "api"
		}

		switch c.Param("mode") {
		case "findOne":
			var result bson.M
			err := mongo.FindOne(app, Input.Table, Input.Where, bson.M{}, &result)
			if err != nil {
				e.Logger.Print(err)
				return c.String(http.StatusNotFound, err.Error())
			} else {
				e.Logger.Print(result)
				return c.JSON(http.StatusOK, result)
			}
			break

		case "findAll":
			var result []bson.M
			err := mongo.FindAll(app, Input.Table, Input.Where, bson.M{}, &result)
			if err != nil {
				fmt.Println(err)
				return c.String(http.StatusNotFound, err.Error())
			} else {
				e.Logger.Print(result)
				return c.JSON(http.StatusOK, result)
			}
			break

		case "findPage":
			var result []bson.M
			err := mongo.FindPage(app, Input.Table, Input.Other.page, Input.Other.limit, Input.Where, bson.M{}, &result)
			if err != nil {
				e.Logger.Print(err)
				return c.String(http.StatusNotFound, err.Error())
			} else {
				e.Logger.Print(result)
				return c.JSON(http.StatusOK, result)
			}
			break

		case "insert":
			err := mongo.Insert(app, Input.Table, Input.Data)
			if err != nil {
				e.Logger.Print(err)
				return c.String(http.StatusNotFound, err.Error())
			} else {
				return c.JSON(http.StatusOK, Input.Data)
			}
			break

		case "update":
			err = mongo.Update(app, Input.Table, Input.Data, Input.Where)
			if err != nil {
				e.Logger.Print(err)
				return c.String(http.StatusNotFound, err.Error())
			} else {
				return c.JSON(http.StatusOK, Input.Data)
			}
			break

		case "updateAll":
			err = mongo.UpdateAll(app, Input.Table, Input.Data, Input.Where)
			if err != nil {
				e.Logger.Print(err)
				return c.String(http.StatusNotFound, err.Error())
			} else {
				return c.JSON(http.StatusOK, Input.Data)
			}
			break

		case "upsert":
			err = mongo.Upsert(app, Input.Table, Input.Data, Input.Where)
			if err != nil {
				e.Logger.Print(err)
				return c.String(http.StatusNotFound, err.Error())
			} else {
				return c.JSON(http.StatusOK, Input.Where)
			}
			break

		case "remove":
			err = mongo.Remove(app, Input.Table, Input.Where)
			if err != nil {
				e.Logger.Print(err)
				return c.String(http.StatusNotFound, err.Error())
			} else {
				return c.JSON(http.StatusOK, Input.Where)
			}
			break

		case "removeAll":
			err = mongo.RemoveAll(app, Input.Table, Input.Where)
			if err != nil {
				e.Logger.Print(err)
				return c.String(http.StatusNotFound, err.Error())
			} else {
				return c.JSON(http.StatusOK, Input.Where)
			}
			break

		case "count":
			count, err := mongo.Count(app, Input.Table, Input.Where)
			if err != nil {
				e.Logger.Print(err)
				return c.String(http.StatusNotFound, err.Error())
			} else {
				e.Logger.Print(count)
				return c.String(http.StatusOK, strconv.Itoa(count))
			}
			break

		case "isEmpty":
			var r string
			if ex := mongo.IsEmpty(app, Input.Table); ex {
				r = "true"
			} else {
				r = "false"
			}
			e.Logger.Print(r)
			return c.String(http.StatusOK, r)
			break

		default:
			return c.String(http.StatusNotFound, "不存在的操作模式："+Input.Mode)
			break
		}

		return c.String(http.StatusNoContent, "")
	}, func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			auth = c.Request().Header.Get(echo.HeaderAuthorization)
			return next(c)
		}
	})

	// 使用 port 设置的端口启动服务
	e.Logger.Fatal(e.StartServer(&http.Server{Addr: ":" + *port}))

	// 设置ssl协议缓存地址
	// e.AutoTLSManager.HostPolicy = autocert.HostWhitelist("<DOMAIN>")
	// Cache certificates
	// e.AutoTLSManager.Cache = autocert.DirCache("~/.cache")

	// 重定向到https不带www
	// e.Pre(middleware.HTTPSNonWWWRedirect())

	// use ssl for 443
	// e.Logger.Fatal(e.StartAutoTLS(":443"))
}
