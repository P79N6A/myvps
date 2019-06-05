package main

import (
	"flag"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"./lib"
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"github.com/xyzj/gopsu"
	ginmiddleware "github.com/xyzj/gopsu/gin-middleware"
)

var (
	version     = "0.0.0"
	goVersion   = ""
	buildDate   = ""
	platform    = ""
	author      = "Xu Yuan"
	programName = "my web tools"
)

var (
	web   = flag.Int("http", 6819, "set port to bind")
	debug = flag.Bool("debug", false, "set if log debug msessage")
	ver   = flag.Bool("version", false, "print version info and exit.")
)

const (
	httpCertFile = "./ca/http.pem"
	httpKeyFile  = "./ca/http-key.pem"
)

// multiRender 预置模板
func multiRender() multitemplate.Renderer {
	r := multitemplate.NewRenderer()
	r.AddFromString("vpsinfo", lib.TPLVpsinfo)
	r.AddFromString("404", ginmiddleware.Template404)
	return r
}
func main() {
	flag.Parse()
	if !*debug {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()
	// 中间件
	// 日志
	// r.Use(ginmiddleware.LoggerWithRolling("", "", 0, loglevel, !*debug, *debug))
	// 错误恢复
	r.Use(gin.Recovery())
	// 渲染模板
	r.HTMLRender = multiRender()
	// 基础路由
	r.GET("/", lib.RemoteIP)
	r.POST("/", lib.RemoteIP)
	gwt := r.Group("/wt")
	gwt.GET("/", lib.IPCache)
	gwt.GET("/kod", lib.Kod)
	gwt.GET("/deluge", lib.Deluge)
	gwt.GET("/mldonkey", lib.Mldonkey)
	gwt.GET("/ssh", lib.Ssh)
	goffice := r.Group("/soho")
	goffice.GET("/", lib.SohoCache)
	goffice.GET("/kod", lib.SohoKod)
	gvps := r.Group("/vps")
	gvps.GET("/", lib.Vps)
	gvps.GET("/v4info", lib.VpsInfo)
	// 错误路由
	r.NoMethod(ginmiddleware.Page405)
	r.NoRoute(ginmiddleware.Page404)

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", *web),
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	fmt.Fprintln(gin.DefaultWriter, "Success start HTTP(S) server at :"+strconv.Itoa(*web))
	var err error
	if gopsu.IsExist(httpCertFile) && gopsu.IsExist(httpKeyFile) {
		err = s.ListenAndServeTLS(httpCertFile, httpKeyFile)
	} else {
		err = s.ListenAndServe()
	}
	if err != nil {
		fmt.Fprintln(gin.DefaultWriter, "Failed start HTTP(S) server at :"+strconv.Itoa(*web)+"|"+err.Error())
	}
}
