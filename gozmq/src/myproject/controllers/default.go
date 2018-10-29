package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}
type EchoController struct {
	beego.Controller
}
type jsonA struct {
	Arr []int `json:"arr"`
}
type jsonB struct {
	Color string `json:"color"`
	Type string `json:"type"`
}
type jsonC struct {
	Data jsonB `json:"data"`
}

type jsonD struct {
	Name string `json:"name"`
	Age int `json:"age"`
	jsonA `json:""`
	jsonC `json:""`
}
type TplController struct {
	beego.Controller
}

func (c *TplController) Get() {
	c.Data["A"]="This is a Test temple file."
	//c.TplName="tmp.tpl"
	c.TplName="a.html"
}
func (c *EchoController) Get() {
	jsonData:=jsonD{
		Name:"sdsdsds",
		Age:12,
		jsonA:jsonA{[]int{23,2,3,}},
		jsonC:jsonC{jsonB{"red","server"}},
	}
	c.Data["json"]=&jsonData
	c.Ctx.Output.JSON(c.Data["json"],true,true)
	//c.ServeJSON()
	//c.Data["a"]=1
	//c.Data["b"]=2
	//c.Ctx.Output.Body([]byte("Hello World!"))
}
func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
