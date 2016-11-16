package park

import (
	"wx_server_go/constants"
	. "wx_server_go/controllers/web/v1"
	. "wx_server_go/models/park"
)

type RegionController struct {
	BaseController
}

// @Title 获取行政区域信息
// @Description
// @Success 200 {object} models.ParkView
// @router / [get]
func (this *RegionController) GetAll() {
	if rs, err := GetRegionCasData(); err == nil {
		this.Data["json"] = ResData(constants.Success, rs)
	} else {
		this.Data["json"] = ResData(constants.DataNull, rs)
	}
	this.ServeJSON()
}
