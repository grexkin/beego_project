package routers

import (
	"beego_project/controllers"
	"github.com/astaxie/beego"
)

func init() {
		beego.Router("/", &controllers.MainController{})
		beego.Router("/user", &controllers.UserController{})
		beego.Router("/test_static", &controllers.StaticController{})
		//beego.Router("/params/?:id", &controllers.ParamsController{})
		beego.Router("/params",&controllers.ParamsController{})
		beego.Router("/other_type_data",&controllers.OtherTypeDataController{})
		beego.Router("/flash_data",&controllers.FlashController{})
}
