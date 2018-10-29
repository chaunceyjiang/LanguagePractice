package routers

import (
	"myproject/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/echo/",&controllers.EchoController{})
    beego.Router("/tpl/",&controllers.TplController{})
}
