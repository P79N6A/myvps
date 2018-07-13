package routers

import (
	"../controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{}, "get:Get;post:Post")
	beego.Router("/webtools", &controllers.MainController{}, "get:WebTools")
	ns := beego.NewNamespace("/wt",
		beego.NSRouter("/", &controllers.MainController{}, "get:IPInfo"),
		beego.NSRouter("ssh", &controllers.MainController{}, "get:Ssh"),
		beego.NSRouter("kod", &controllers.MainController{}, "get:Kod"),
		beego.NSRouter("deluge", &controllers.MainController{}, "get:Deluge"),
		beego.NSRouter("mldonkey", &controllers.MainController{}, "get:Mldonkey"),
	)
	beego.AddNamespace(ns)
	ns = beego.NewNamespace("/vps",
		beego.NSRouter("/info", &controllers.VpsController{}, "get:VpsInfo"))
	beego.AddNamespace(ns)
}
