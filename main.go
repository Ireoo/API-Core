package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
)

type input struct {
	Where interface{} `json:"where" form:"where" query:"where"`
	Data  interface{} `json:"data" form:"data" query:"data"`
	Other interface{} `json:"other" form:"other" query:"other"`
}

const version = "1.0.0"

var ver = flag.Bool("v", false, "版本信息")
var port = flag.String("p", "", "端口地址")

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
		}
		e.Logger.Print(Input)
		return c.JSON(http.StatusOK, Input)
	})

	// 使用 2019 端口启动服务
	e.Logger.Fatal(e.StartServer(&http.Server{Addr: ":2019"}))

	// 设置ssl协议缓存地址
	// e.AutoTLSManager.HostPolicy = autocert.HostWhitelist("<DOMAIN>")
	// Cache certificates
	// e.AutoTLSManager.Cache = autocert.DirCache("~/.cache")

	// 重定向到https不带www
	// e.Pre(middleware.HTTPSNonWWWRedirect())

	// use ssl for 443
	// e.Logger.Fatal(e.StartAutoTLS(":443"))
}
