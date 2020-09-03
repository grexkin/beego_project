package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

type UserController struct {
	beego.Controller
}

func (u *UserController) Get() {
	u.Data["Username"] = "wanghui"
	u.Data["Age"] = "12"
	u.TplName = "user.html"
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
