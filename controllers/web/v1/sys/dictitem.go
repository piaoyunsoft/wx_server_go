package sys

import (
	"wx_server_go/constants"
	"wx_server_go/controllers/api"
	"wx_server_go/models/sys"
)

type SysController struct {
	v1.BaseController
}

// @Title 获取
// @Description
// @Success 200 {object} models.DictItem
// @router /dictItem [get]
func (this *SysController) GetAll() {
	var query = make(map[string]string)

	if v := this.GetString("dictcode"); v != "" {
		query["dictcode"] = v
	}
	if rs, err := sys.GetDictItemsByDictCode(query); err == nil {
		this.Data["json"] = v1.ResData(constants.Success, rs)
	} else {
		this.Data["json"] = v1.ResData(constants.DataNull, rs)
	}
	this.ServeJSON()
}
