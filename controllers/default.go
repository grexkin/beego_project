package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

type UserController struct {
	beego.Controller
}

//静态文件的controller
type StaticController struct {
	beego.Controller
}

//从前端传参到后端
type ParamsController struct {
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
	//map 渲染
	test_map := map[string]string{
		"name": "隔壁老王",
		"age": "貌美如花",
	}
	u.Data["maps"] = test_map
	//map_struct 渲染
	test_map_struct := make(map[int]UserStruct)
	test_map_struct[0] = UserStruct{ID:0,UserName:"张三",Age:29}
	test_map_struct[1] = UserStruct{ID:1,UserName:"张三",Age:18}
	test_map_struct[2] = UserStruct{ID:2,UserName:"张三",Age:28}
	u.Data["map_struct"] = test_map_struct
	//slice 切片渲染
	slices := []int{2,4,6,8,10}
	u.Data["slices"] = slices
	u.TplName = "user.html"
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func (s *StaticController)Get()  {
	s.TplName = "static_test.html"
}

func (p *ParamsController)Get()  {
	//方式一： http://127.0.0.1:8080/params?name=zhangsan
	//url: /params
	name := p.GetString("name")
	fmt.Println("========",name)
	name1 := p.Input().Get("name")
	fmt.Println("===",name1)
	name2 := p.Ctx.Input.Param(":name")  //不能获取
	fmt.Println("=====name2",name2)
	//方式二：http://127.0.0.1:8080/params/daaa
	// /params/?:id
	// /params/?:id int    指定id是int类型
	id := p.Ctx.Input.Param(":id")  //用于/xxx/value 这种形式的参数获取
	fmt.Println("=====id",id)
	id1 := p.Input().Get(":id")   //不能获取
	fmt.Println("====id1",id1)
	id2 := p.GetString(":id") 
	fmt.Println("====id2",id2)
	p.TplName = "param_test.html"
}