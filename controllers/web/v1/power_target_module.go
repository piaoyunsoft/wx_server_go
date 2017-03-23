package v1

import (
	"wx_server_go/constants"
	"wx_server_go/models"
)

type ModuleController struct {
	BaseController
}

// @Title 获取模块信息
// @Description
// @Success 200 {object} models.PowerModule
// @router / [get]
func (this *ModuleController) GetAll() {
	req := new(models.SeaPowertargetModule)
	if rs, err := req.GetAll(); err == nil {
		this.Data["json"] = ResData(constants.Success, rs)
	} else {
		this.Data["json"] = ResData(constants.DataNull, rs)
	}
	this.ServeJSON()
}

// @router /valid [get]
func (this *ModuleController) GetValidModule() {
	req := new(models.SeaPowertargetModule)
	req.Status = "aa"
	if rs, err := req.GetAll(); err == nil {
		this.Data["json"] = ResData(constants.Success, rs)
	} else {
		this.Data["json"] = ResData(constants.DataNull, rs)
	}
	this.ServeJSON()
}
