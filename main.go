package main

import (
	"flag"
	"fmt"
	"net/http"

	"os"

	Info "github.com/Ireoo/API-Core/info"

	mongo "github.com/Ireoo/API-Core/libs/mongodb"
	Router "github.com/Ireoo/API-Core/libs/router"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

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

func main() {
	if ver {
		fmt.Printf("API-Core version: %s\nbuild time: %s\n", Info.Version, Info.BuildTime)
		return
	}

	fmt.Printf("API-Core version: %s\nbuild time: %s\n", Info.Version, Info.BuildTime)
	fmt.Println("")
	fmt.Println("")

	_ = mongo.New(command_uri)

	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	corsConf := cors.DefaultConfig()
	corsConf.AddAllowHeaders("Authorization")
	corsConf.AllowAllOrigins = true
	router.Use(cors.New(corsConf))

	// 设置静态文件
	router.StaticFS("/static/css", http.Dir("./static/static/css"))
	router.StaticFS("/static/js", http.Dir("./static/static/js"))
	router.StaticFS("/static/fonts", http.Dir("./static/static/fonts"))
	router.StaticFile("/", "./static/index.html")
	router.StaticFile("/admin", "./static/admin.html")
	router.StaticFile("/favicon.ico", "./static/favicon.ico")

	// 程序核心部分
	router.POST("/:table/:mode", func(c *gin.Context) {
		Router.Table(c, secret)
	})

	//router.POST("/:mode", func(c *gin.Context) {
	//	Router.Mode(c, secret)
	//})

	_port := os.Getenv("PORT")

	if _port != "" {
		port = _port
	}

	router.Run(":" + port)
}
