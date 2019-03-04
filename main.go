package main

import (
	"API-Core/libs/mongo"
	"encoding/hex"
	"flag"
	"fmt"
	"net/http"
	"strconv"

	"golang.org/x/crypto/acme/autocert"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"gopkg.in/mgo.v2/bson"
)

type input struct {
	Where bson.M `json:"where"`
	Data  bson.M `json:"data"`
	Other struct {
		page  int
		limit int
	} `json:"other"`
	Table string `json:"table"`
	Mode  string `json:"mode"`
	Auth  string `json:"auth"`
}

type appInfo struct {
	Id bson.ObjectId `bson:"_id"`
}

const version = "1.0.0"

var ver = flag.Bool("v", false, "版本信息")
var port = flag.String("p", "2019", "端口地址,默认: 2019")
var ssl = flag.Bool("ssl", false, "是否开启SSL功能,默认不开启")

//var _host = flag.String("host", "", "设置绑定域名,默认不绑定,需要绑定请设置绑定的域名,如: x.domain.com")

var auth = ""

func main() {
	flag.Parse()

	if *ver {
		fmt.Printf(`API-Core version: %s`, version)
		return
	}

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

	}, func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			auth = c.Request().Header.Get(echo.HeaderAuthorization)
			return next(c)
		}
	})

	if !*ssl {
		// 使用 port 设置的端口启动服务
		e.Logger.Fatal(e.StartServer(&http.Server{Addr: ":" + *port}))
	} else {
		// 设置ssl协议缓存地址
		// e.AutoTLSManager.HostPolicy = autocert.HostWhitelist("<DOMAIN>")
		// Cache certificates
		e.AutoTLSManager.Cache = autocert.DirCache("~/.cache")

		// 重定向到https不带www
		e.Pre(middleware.HTTPSRedirect())

		// use ssl for 443
		e.Logger.Fatal(e.StartAutoTLS(":443"))
	}
}
