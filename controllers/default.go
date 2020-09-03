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

type UserStruct struct {
	ID       int
	UserName string
	Age      int
}

func (u *UserController) Get() {
	//结构体渲染
	user := UserStruct{ID:1,UserName:"zhangsan",Age:23}
	u.Data["user"] = user
	//数组渲染
	arrays := [5]int{1,2,3,4,5}
	u.Data["arrays"] = arrays
	//数组结构体渲染
	arr_struct := [5]UserStruct{
		{ID:1,UserName:"zhangsan1",Age:23},
		{ID:2,UserName:"zhangsan2",Age:21},
		{ID:3,UserName:"zhangsan3",Age:33},
		{ID:4,UserName:"zhangsan4",Age:19},
		{ID:5,UserName:"zhangsan5",Age:40},
	}
	u.Data["arr_struct"] = arr_struct
	u.TplName = "user.html"
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
