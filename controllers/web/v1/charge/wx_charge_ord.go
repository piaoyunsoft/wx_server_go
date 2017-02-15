package charge

import (
	"wx_server_go/constants"
	. "wx_server_go/controllers/web/v1"
	. "wx_server_go/models/charge"
)

type ChargeController struct {
	BaseController
}

// @Title 获取微信订单信息
// @Description 分页获取
// @Success 200 {object} models.WxChargeOrd
// @router /getWxOrd [get]
func (this *ChargeController) GetWxChargeOrd() {
	var query = make(map[string]string)
	var page int = 1
	var size int = 10

	if v, err := this.GetInt("page"); err == nil {
		page = v
	}
	if v, err := this.GetInt("size"); err == nil {
		size = v
	}
	if v := this.GetString("keyword"); v != "" {
		query["keyword"] = v
	}
	if v := this.GetString("begin"); v != "" {
		query["begin"] = v
	}
	if v := this.GetString("end"); v != "" {
		query["end"] = v
	}
	if v := this.GetString("status"); v != "" {
		query["status"] = v
	}
	if v := this.GetString("payPtf"); v != "" {
		query["payPtf"] = v
	}
	if total, rs, err := GetPageChargeOrds(query, page, size); err == nil {
		this.Data["json"] = ResData(constants.Success, PageData{Data: rs, Total: total})
	} else {
		this.Data["json"] = ResData(constants.DataNull, PageData{Data: rs, Total: total})
	}
	this.ServeJSON()
}
