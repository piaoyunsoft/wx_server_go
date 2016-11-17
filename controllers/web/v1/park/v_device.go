package park

import (
	"wx_server_go/constants"
	. "wx_server_go/controllers/web/v1"
	. "wx_server_go/models/park"
)

type DeviceController struct {
	BaseController
}

// @Title 获取设备信息
// @Description
// @Success 200 {object} models.DeviceView
// @router / [get]
func (this *DeviceController) GetAll() {
	var query = make(map[string]string)

	if v := this.GetString("regionID"); v != "" {
		query["regionID"] = v
	}
	if rs, err := GetDeviceViewList(query); err == nil {
		this.Data["json"] = ResData(constants.Success, rs)
	} else {
		this.Data["json"] = ResData(constants.DataNull, rs)
	}
	this.ServeJSON()
}
