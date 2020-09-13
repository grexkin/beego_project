package main

import (
	_ "beego_project/routers"
	"github.com/astaxie/beego"
)

func main() {
	//beego.SetViewsPath("front")
	//beego.SetStaticPath("static","front")
	/* 配置文件读取
	httpport := beego.AppConfig.String("httpport")
	httpport1,_ := beego.AppConfig.Int("httpport")
	fmt.Println(httpport)
	fmt.Printf("%T,%T\n",httpport,httpport1)
	username := beego.AppConfig.String("username")
	passwd := beego.AppConfig.String("passwd")
	fmt.Println(username,passwd)
	 */
	beego.BConfig.WebConfig.EnableXSRF = true
	beego.BConfig.WebConfig.XSRFKey = "MUDepLTLbGXPcTvQgO27Xu7G69UUm8SeIM8gv4o"
	beego.BConfig.WebConfig.XSRFExpire = 3600   //单位是秒

	beego.Run()
}