package controllers

import (
	"encoding/json"
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

//其他数据格式
type OtherTypeDataController struct {
	beego.Controller
}

//Flash数据
type FlashController struct {
	beego.Controller
}

type UserStruct struct {
	ID       int
	UserName string   `form:"username"`
	Age      int	  `form:"age"`
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

func (p *ParamsController)Post()  {
	//获取前端返回的数据
	//第一种方式
	/*
	username := p.GetString("username")
	usernames := p.GetStrings("username")  //获取到所有的username，组成slice
	fmt.Println(username)
	fmt.Println(usernames)
	age := p.GetString("age")
	fmt.Println(username,age)
	//第二种方式
	username1 := p.Input().Get("username")
	age1 := p.Input().Get("age")
	fmt.Println(username1,age1)
	//第三种方式: url需要配置，所以不可能获取到数据
	//p.Ctx.Input.Param("username")
	//获取其他数据类型
	age2,_ := p.GetInt64("age")
	price,_ := p.GetFloat("price")
	is_true,_ := p.GetBool("is_true")
	fmt.Println(username,age2,price,is_true)
	 */
	//第三种获取form表单的方式
	/*
	user := UserStruct{}
	err := p.ParseForm(&user)
	fmt.Println(err)

	if err == nil {
		fmt.Println(&user)
	}
	p.TplName = "success.html"
	 */
	//获取ajax数据
	body := p.Ctx.Input.RequestBody    //二进制数据
	user := UserStruct{}
	json.Unmarshal(body,&user)
	fmt.Println(user)
	result := map[string]string{"code":"200","message":"处理成功"}
	p.Data["json"] = result
	p.ServeJSON()   //返回json格式
}

func (o *OtherTypeDataController)Get()  {
	user := UserStruct{ID: 1,UserName: "abcd",Age: 18}
	// json 格式传输
	/*
	o.Data["json"] = &user
	o.ServeJSON()
	 */
	/* xml 格式
	o.Data["xml"] = &user
	o.ServeXML()
	 */
	/* jsonp 格式
	o.Data["jsonp"] = &user
	o.ServeJSONP()
	 */
	o.Data["yaml"] = &user
	o.ServeYAML()
}

//flash controller
func (f *FlashController)Get()  {
	flash := beego.ReadFromRequest(&f.Controller)
	notice := flash.Data["notice"]
	err := flash.Data["error"]
	fmt.Println("===")
	fmt.Println(err)
	if len(notice) != 0 {}
	if len(err) != 0 {
		f.TplName = "error.html"
	}else if len(notice) != 0{
		f.TplName = "success.html"
	} else {
		f.TplName = "flash.html"
	}
}

func (f *FlashController)Post()  {
	flash := beego.NewFlash()   //初始化flash
	username := f.GetString("username")
	pwd := f.GetString("pwd")
	fmt.Println(username)
	fmt.Println(pwd)
	if len(username) == 0 {
		fmt.Println("用户名不能为空")
		flash.Error("用户名不能为空！！")
		flash.Store(&f.Controller)
		f.Redirect("/flash_data",302)
	}else if pwd != "123456"{
		fmt.Println("密码错误")
		flash.Error("密码错误")
		flash.Store(&f.Controller)
		f.Redirect("/flash_data",302)
	}else {
		flash.Notice("成功")
		flash.Store(&f.Controller)
		f.Redirect("/flash_data",302)
	}

}