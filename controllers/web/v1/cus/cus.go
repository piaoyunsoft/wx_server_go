package cus

import (
	"wx_server_go/constants"
	"wx_server_go/controllers/web/v1"
	"wx_server_go/models/cus"
)

type CusController struct {
	v1.BaseController
}

// @router / [get]
func (this *CusController) GetAll() {
	query := make(map[string]string)

	if rs, err := cus.GetPlatCus(query); err == nil {
		this.Data["json"] = v1.ResData(constants.Success, rs)
	} else {
		this.Data["json"] = v1.ResData(constants.DataNull, rs)
	}
	this.ServeJSON()
}

// @router /my [get]
func (this *CusController) GetMyCus() {
	query := make(map[string]string)
	query["cusID"] = v1.CusId
	if rs, err := cus.GetPlatCus(query); err == nil {
		this.Data["json"] = v1.ResData(constants.Success, rs)
	} else {
		this.Data["json"] = v1.ResData(constants.DataNull, rs)
	}
	this.ServeJSON()
}
