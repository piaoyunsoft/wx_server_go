package v1

import (
	"wx_server_go/constants"
	"wx_server_go/models"
)

type CusController struct {
	BaseController
}

// @router / [get]
func (this *CusController) GetAll() {
	req := new(models.SeaPlatcus)

	if rs, err := req.GetAll(); err == nil {
		this.Data["json"] = ResData(constants.Success, rs)
	} else {
		this.Data["json"] = ResData(constants.DataNull, rs)
	}
	this.ServeJSON()
}

// @router /my [get]
func (this *CusController) GetMyCus() {
	req := new(models.SeaPlatcus)
	req.Cusid = this.CusId

	if rs, err := req.GetAll(); err == nil {
		this.Data["json"] = ResData(constants.Success, rs)
	} else {
		this.Data["json"] = ResData(constants.DataNull, rs)
	}
	this.ServeJSON()
}
