package common

import (
	"github.com/astaxie/beego"
	"github.com/ddliao/go-lib/slog"
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
		slog.Error(err.Error())
		this.Data["json"] = map[string]interface{}{"isSuc": false}
		this.ServeJSON()
		return
	}
	path := "/static/upload/" + h.Filename
	fullPath := "." + path
	//	f.Close()
	err = this.SaveToFile("file", fullPath)
	if err != nil {
		slog.Error(err.Error())
		this.Data["json"] = map[string]interface{}{"isSuc": false}
		this.ServeJSON()
		return
	}
	this.Data["json"] = map[string]interface{}{"isSuc": true, "url": path}
	this.ServeJSON()
}
