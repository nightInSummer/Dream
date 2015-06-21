package routers

import (
	"Dream/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.IndexController{})
	beego.Router("/contect.html", &controllers.ContectController{})
	beego.Router("/jinru.html", &controllers.InformationController{}, "get:Enter")
	beego.Router("/jinru.html", &controllers.InformationController{}, "post:Contect")
	beego.Router("/reg.html", &controllers.CheckController{}, "get:Reg;post:PostReg")
	beego.Router("/index.html", &controllers.CheckController{}, "get:Out")
	beego.Router("/index.html", &controllers.CheckController{}, "get:Login;post:PostLogin")
}
