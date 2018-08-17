package main

import (
	"flag"
	"os"
	"strings"

	_ "./routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/xyzj/mxgo"
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
	conf     = flag.String("conf", "conf/app.conf", "set config file path")
	usehttps = flag.Bool("usehttps", false, "set if enable https")
	debug    = flag.Bool("debug", false, "set if log debug msessage")
	ver      = flag.Bool("version", false, "print version info and exit.")
)

func main() {
	os.MkdirAll("./log", 0775)
	flag.Parse()
	if *ver {
		println(mxgo.VersionInfo(programName, version, goVersion, platform, buildDate, author))
		os.Exit(1)
	}
	if len(strings.TrimSpace(*conf)) == 0 || !mxgo.IsExist(*conf) {
		flag.PrintDefaults()
		os.Exit(1)
	}
	beego.LoadAppConfig("ini", *conf)
	logs.SetLogger(logs.AdapterConsole)
	// logs.SetLogger(logs.AdapterFile, `{"filename":"log/mywebtools.log"}`)
	if !*debug {
		beego.BConfig.RunMode = "prod"
		logs.SetLevel(logs.LevelWarn)
	} else {
		beego.BConfig.RunMode = "dev"
		logs.SetLevel(logs.LevelDebug)
		logs.EnableFuncCallDepth(true)
	}
	logs.Async(10)
	beego.BConfig.WebConfig.AutoRender = false
	if *usehttps {
		beego.BConfig.Listen.EnableHTTPS = true
		p, _ := beego.AppConfig.Int("httpport")
		beego.BConfig.Listen.HTTPSAddr = beego.AppConfig.String("httpaddr")
		beego.BConfig.Listen.HTTPSPort = p + 1
		beego.BConfig.Listen.HTTPSCertFile = "./ca/server.crt"
		beego.BConfig.Listen.HTTPSKeyFile = "./ca/server.key"
	}
	beego.Run()
}
