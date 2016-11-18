package park

import (
	"wx_server_go/constants"
	. "wx_server_go/controllers/web/v1"
	. "wx_server_go/models/park"
)

type TrafficController struct {
	BaseController
}

// @Title 获取通行记录信息
// @Description
// @Success 200 {object} models.TrafficView
// @router / [get]
func (this *TrafficController) GetAll() {
	var query = make(map[string]string)

	if v := this.GetString("carNum"); v != "" {
		query["carNum"] = v
	}
	if v := this.GetString("phone"); v != "" {
		query["phone"] = v
	}
	if v := this.GetString("userName"); v != "" {
		query["userName"] = v
	}
	if v := this.GetString("begin"); v != "" {
		query["begin"] = v
	}
	if v := this.GetString("end"); v != "" {
		query["end"] = v
	}
	if rs, err := GetTrafficViewList(query); err == nil {
		this.Data["json"] = ResData(constants.Success, rs)
	} else {
		this.Data["json"] = ResData(constants.DataNull, rs)
	}
	this.ServeJSON()
}
