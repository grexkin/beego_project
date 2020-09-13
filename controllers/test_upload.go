package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"strconv"
	"time"
)

type UploadController struct {
	beego.Controller
}

func (u *UploadController)Prepare()  {
	u.EnableXSRF = false
}

func (u *UploadController)Get()  {
	u.TplName = "test_upload.html"
}

func (u *UploadController)Post()  {
	//获取要上传的文件
	f,h,err := u.GetFile("upload_file")
	defer func() {
		f.Close()
	}()
	//生成时间戳
	time_unix_int := time.Now().UnixNano()
	time_unix_str := strconv.FormatInt(time_unix_int,10)
	//获取文件名
	filename := h.Filename
	if err != nil {
		return
	}
	fmt.Println(filename,time_unix_str)
	//保存获取的文件
	u.SaveToFile("upload_file","static/upload/"+time_unix_str+filename)
	//u.TplName = "success.html"

	//ajax文件上传
	u.Data["json"] = map[string]string{"code":"200","message":"上传成功"}
	u.ServeJSONP()
}