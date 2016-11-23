package common

import (
	"wx_server_go/utils"

	"github.com/astaxie/beego"
)

type FileController struct {
	beego.Controller
}

// @Title 文件上传
// @Description
// @Success 200 {object} models.Platcuswxtask
// @router / [post]
func (this *FileController) UploadFile() {
	f, h, err := this.GetFile("file") //获取上传的文件
	defer f.Close()
	if err == nil {
		//		f.Close()
	} else {
		utils.Error(err.Error())
		this.Data["json"] = map[string]interface{}{"isSuc": false}
		this.ServeJSON()
		return
	}
	path := "/static/upload/" + h.Filename
	fullPath := "." + path
	//	f.Close()
	err = this.SaveToFile("file", fullPath)
	if err != nil {
		utils.Error(err.Error())
		this.Data["json"] = map[string]interface{}{"isSuc": false}
		this.ServeJSON()
		return
	}
	this.Data["json"] = map[string]interface{}{"isSuc": true, "url": path}
	this.ServeJSON()
}
