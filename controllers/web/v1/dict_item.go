package v1

import (
	"wx_server_go/constants"
	"wx_server_go/models"
)

type SysController struct {
	BaseController
}

// @Title 获取
// @Description
// @Success 200 {object} models.DictItem
// @router /dictItem [get]
func (this *SysController) GetAll() {
	req := new(models.SeaDictitem)
	req.Dictcode = this.GetString("dictcode")

	if rs, err := req.GetAll(); err == nil {
		this.Data["json"] = ResData(constants.Success, rs)
	} else {
		this.Data["json"] = ResData(constants.DataNull, rs)
	}
	this.ServeJSON()
}
