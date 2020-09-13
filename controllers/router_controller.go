package controllers

import (
	"github.com/astaxie/beego"
)

type RouterController struct {
	beego.Controller
}

func (r *RouterController)Get()  {
	/*
	id := r.Ctx.Input.Param("id")
	fmt.Println("id====>",id)
	 */
	if r.Ctx.Request.Method == "POST" {
		r.StopRun()   //终止下面的流程
		r.TplName = "test_router_post.html"
	} else if r.Ctx.Request.Method == "GET"{
		r.TplName = "test_router.html"
	}
}