package main

import (
	"encoding/hex"

	"flag"
	"fmt"
	"net/http"
	"os"

	Info "github.com/Ireoo/API-Core/info"
	"github.com/Ireoo/API-Core/libs/conf"
	"github.com/Ireoo/API-Core/libs/mongodb"
	Router "github.com/Ireoo/API-Core/libs/router"

	"golang.org/x/crypto/acme/autocert"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	Logger "github.com/labstack/gommon/log"
	"gopkg.in/mgo.v2/bson"
)

// var ver = flag.Bool("v", false, "版本信息")
// var port = flag.String("p", "2019", "端口地址")
// var ssl = flag.Bool("ssl", false, "是否开启SSL功能,默认不开启")
// var secret = flag.String("secret", "94f3eee0-218f-41fc-9318-94cf5430fc7f", "管理权限密钥")

var (
	ver         bool
	ssl         bool
	port        string
	secret      string
	command_uri string
)

func init() {
	flag.StringVar(&port, "p", "2019", "端口地址")
	flag.StringVar(&secret, "secret", "94f3eee0-218f-41fc-9318-94cf5430fc7f", "管理权限密钥")
	flag.StringVar(&command_uri, "mongodb", "", "MongoDB connect uri")
	flag.BoolVar(&ver, "v", false, "版本信息")
	flag.BoolVar(&ssl, "ssl", false, "是否开启SSL功能,默认不开启")
	flag.Parse()
}

//var _host = flag.String("host", "", "设置绑定域名,默认不绑定,需要绑定请设置绑定的域名,如: x.domain.com")

var auth = ""

func main() {
	// flag.Parse()

	if ver {
		//fmt.Printf(`API-Core version: %s`, version)
		fmt.Printf("API-Core version: %s\nbuild time: %s\n", Info.Version, Info.BuildTime)
		return
	}

	fmt.Printf("API-Core version: %s\nbuild time: %s\n", Info.Version, Info.BuildTime)
	fmt.Println("")
	fmt.Println("")

	_ = mongo.New(command_uri)

	e := echo.New()

	e.HideBanner = true

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.Gzip())
	// 主要用于拦截panic错误并且在控制台打印错误日志，避免echo程序直接崩溃。
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.Logger.SetLevel(Logger.DEBUG)

	// e.Logger.Print(os.Args)

	// 设置静态文件
	e.Use(middleware.Static("./static"))
	e.Static("/static", "static/static")
	e.File("/favicon.ico", "static/favicon.ico")
	e.File("/", "static/index.html")
	e.File("/admin", "static/admin.html")

	// 	e.GET("/", func(c echo.Context) error {
	// 		return c.HTML(http.StatusOK, `<h1 style="text-align: center;">欢迎使用 iData API 数据中心!</h1>
	// <h3 style="text-align: center;">Welcome to the iData API Data Center!</h3>`)
	// 	})

	// 程序核心部分
	e.POST("/:table/:mode", func(c echo.Context) (err error) {
		return Router.Table(c, secret, auth)
	}, func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			auth = c.Request().Header.Get(echo.HeaderAuthorization)
			return next(c)
		}
	})

	e.POST("/:mode", func(c echo.Context) (err error) {
		Input := new(conf.Input)
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
		if Input.Auth != secret {
			//var result bson.M
			AppInfo := new(conf.AppInfo)
			error := mongo.FindOne(app, "apps", bson.M{"secret": Input.Auth}, bson.M{}, &AppInfo)
			if error != nil {
				e.Logger.Print(error)
				return c.String(http.StatusNonAuthoritativeInfo, "The authorization verification information does not exist. Please verify.")
			}
			app = hex.EncodeToString([]byte(AppInfo.Id))
			//fmt.Println(app)
		}

		switch Input.Mode {
		case "collectionNames":
			names, error := mongo.CollectionNames(app)
			if error != nil {
				e.Logger.Print(error)
				return c.String(http.StatusNotFound, error.Error())
			}
			e.Logger.Print(names)
			return c.JSON(http.StatusOK, names)

		default:
			return c.String(http.StatusNotFound, "不存在的操作模式："+Input.Mode)
		}

	}, func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			auth = c.Request().Header.Get(echo.HeaderAuthorization)
			return next(c)
		}
	})

	_port := os.Getenv("PORT")

	if _port != "" {
		port = _port
	}

	if !ssl {
		// 使用 port 设置的端口启动服务
		fmt.Println("")
		e.Logger.Fatal(e.StartServer(&http.Server{Addr: ":" + port}))
		fmt.Println("")
		fmt.Println("")
	} else {
		// 设置ssl协议缓存地址
		e.AutoTLSManager.HostPolicy = autocert.HostWhitelist("localhost", "ireoo.com")
		// Cache certificates
		e.AutoTLSManager.Cache = autocert.DirCache(".cache")

		// 重定向到https不带www
		e.Pre(middleware.HTTPSRedirect())

		// use ssl for 443
		fmt.Println("")
		e.Logger.Fatal(e.StartAutoTLS(":443"))
		fmt.Println("")
		fmt.Println("")
	}
}
