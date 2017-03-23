package v1

import (
	"wx_server_go/constants"
	"wx_server_go/models"
)

type VipClsController struct {
	BaseController
}

// @router /valid [get]
func (this *VipClsController) GetValid() {
	req := new(models.SeaVipcls)
	req.Status = "aa"
	req.Comid = CusId

	if rs, err := req.GetAll(); err == nil {
		this.Data["json"] = ResData(constants.Success, rs)
	} else {
		this.Data["json"] = ResData(constants.DataNull, rs)
	}
	this.ServeJSON()
}

//// @router /valid [get]
//func (this *VipClsController) GetValid() {
//	query := make(map[string]string)
//
//	query["status"] = "aa"
//	query["comid"] = CusId
//	if rs, err := models.GetVipCls(query); err == nil {
//		this.Data["json"] = ResData(constants.Success, rs)
//	} else {
//		this.Data["json"] = ResData(constants.DataNull, rs)
//	}
//	this.ServeJSON()
//}
