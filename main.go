package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gookit/color"

	Info "github.com/Ireoo/API-Core/info"

	mongo "github.com/Ireoo/API-Core/libs/mongodb"
	Router "github.com/Ireoo/API-Core/libs/router"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	ver         bool
	ssl         bool
	Debug       bool
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
	flag.BoolVar(&Debug, "debug", false, "是否开启Debug功能,默认不开启")
	flag.Parse()
}

func main() {

	// 输出版本信息
	if ver {
		fmt.Printf("API-Core version: %s\nbuild time: %s\n", Info.Version, Info.BuildTime)
		return
	}
	fmt.Printf("API-Core version: %s\nbuild time: %s\n", Info.Version, Info.BuildTime)
	fmt.Println("")
	fmt.Println("")

	// 建立与数据库的连接
	_ = mongo.New(command_uri)

	// 是否开启debug
	if !Debug {
		Debug = os.Getenv("DEBUG") == "true"
		if string(Info.Version) != "unknown version" {
			if Debug {
				gin.SetMode(gin.DebugMode)
			} else {
				gin.SetMode(gin.ReleaseMode)
			}
		} else {
			Debug = true
			gin.SetMode(gin.DebugMode)
		}
	}

	// 初始化网页服务器
	router := gin.Default()

	// 设置 CORS
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
		Router.Table(c, secret, Debug)
	})

	router.GET("/:table/:mode", func(c *gin.Context) {
		Router.TableGet(c, secret, Debug)
	})

	// 获取启动服务绑定的端口
	_port := os.Getenv("PORT")
	if _port != "" {
		port = _port
	}

	// 启动服务
	fmt.Println("")
	log.Println(color.FgGreen.Sprintf(color.Bold.Sprintf(" ✔ ")) + " Listening " + color.FgGreen.Sprintf("http(s)://0.0.0.0:%s", port))
	fmt.Println("")
	router.Run(":" + port)
}
