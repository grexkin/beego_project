package controllers

import (
	"github.com/astaxie/beego"
	"html/template"
)

type TestXsrfController struct {
	beego.Controller
}

//关闭这个controller的xsrf
func (t *TestXsrfController)Prepare()  {
	t.EnableXSRF = false
}
func (t *TestXsrfController)Get()  {
	t.Data["xsrfdata"] = template.HTML(t.XSRFFormHTML())
	t.TplName = "test_xsrf.html"
}

func (t *TestXsrfController)Post()  {
	t.TplName = "success.html"
}
