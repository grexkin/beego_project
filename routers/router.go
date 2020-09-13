package routers

import (
	"beego_project/controllers"
	"github.com/astaxie/beego"
)

func init() {
		beego.Router("/", &controllers.MainController{})
		beego.Router("/user", &controllers.UserController{})
		beego.Router("/test_static", &controllers.StaticController{})
		//beego.Router("/params/:id", &controllers.ParamsController{})
		//beego.Router("/params/?:id:int", &controllers.ParamsController{})
		//beego.Router("/params/?:id:string", &controllers.ParamsController{})
		//beego.Router("/params/:id([0-9]+)", &controllers.ParamsController{})
		//beego.Router("/params/?:id)", &controllers.ParamsController{})
		//beego.AutoRouter(&controllers.ParamsController{}) //自动路由http://127.0.0.1:8080/params/get
		//自定义路由
		//beego.Router("/params/?:id)", &controllers.ParamsController{},"get:Get,post:Post")
		beego.Router("/params",&controllers.ParamsController{})
		beego.Router("/other_type_data",&controllers.OtherTypeDataController{})
		beego.Router("/flash_data",&controllers.FlashController{})
		beego.Router("/flash_ajax",&controllers.FlashAjaxController{})
		//路由配置方式
		//beego.Router("/test_router",&controllers.RouterController{})  //固定路由
		//beego.Router("/test_router/:id:int",&controllers.RouterController{}) //正则路由
		//beego.Router("/test_router/?:id:int",&controllers.RouterController{}) //正则路由
		//beego.AutoRouter(&controllers.RouterController{}) //自动路由
		//自定义路由
		beego.Router("/test_router/?:id:int",&controllers.RouterController{},"get,post:Get")
		//xsrf
		beego.Router("/test_xsrf",&controllers.TestXsrfController{})
		beego.Router("/test_upload",&controllers.UploadController{})
}