package main

import (
	"flag"
	"fmt"

	"os"
	// "strconv"

	Info "github.com/Ireoo/API-Core/info"

	mongo "github.com/Ireoo/API-Core/libs/mongodb"
	Router "github.com/Ireoo/API-Core/libs/router"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	// "github.com/gookit/color"
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

	//gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	corsConf := cors.DefaultConfig()
	corsConf.AddAllowHeaders("Authorization")
	corsConf.AllowAllOrigins = true
	router.Use(cors.New(corsConf))

	// e.Logger.Print(os.Args)

	// 设置静态文件
	//e.Use(middleware.Static("./static"))
	//e.Static("/static", "static/static")
	//e.File("/favicon.ico", "static/favicon.ico")
	//e.File("/", "static/index.html")
	//e.File("/admin", "static/admin.html")

	// 	e.GET("/", func(c echo.Context) error {
	// 		return c.HTML(http.StatusOK, `<h1 style="text-align: center;">欢迎使用 iData API 数据中心!</h1>
	// <h3 style="text-align: center;">Welcome to the iData API Data Center!</h3>`)
	// 	})

	// 程序核心部分
	router.POST("/:table/:mode", func(c *gin.Context) {
		Router.Table(c, secret)
	})

	// router.POST("/:mode", func(c *gin.Context) {
	// 	Router.Table(c, secret)
	// })

	_port := os.Getenv("PORT")

	if _port != "" {
		port = _port
	}

	if !ssl {
		// 使用 port 设置的端口启动服务
		//fmt.Println("")
		//e.Logger.Fatal(e.StartServer(&http.Server{Addr: ":" + port}))
		//fmt.Println("")
		//fmt.Println("")
		router.Run(":" + port)
	} else {
		// 设置ssl协议缓存地址
		//e.AutoTLSManager.HostPolicy = autocert.HostWhitelist("localhost", "ireoo.com")
		//// Cache certificates
		//e.AutoTLSManager.Cache = autocert.DirCache(".cache")
		//
		//// 重定向到https不带www
		//e.Pre(middleware.HTTPSRedirect())
		//
		//// use ssl for 443
		//fmt.Println("")
		//e.Logger.Fatal(e.StartAutoTLS(":443"))
		//fmt.Println("")
		//fmt.Println("")
	}
}
