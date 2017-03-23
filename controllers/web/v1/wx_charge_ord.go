package v1

import (
	"wx_server_go/constants"
	"wx_server_go/models"
)

type ChargeController struct {
	BaseController
}

// @Title 获取微信订单信息
// @Description 分页获取
// @Success 200 {object} models.WxChargeOrd
// @router /getWxOrd [get]
func (this *ChargeController) GetWxChargeOrd() {
	req := new(models.SeaWxchargeodr)
	req.PageSize = this.ToIntEx("size", 10)
	req.PageIndex = this.ToIntEx("page", 1)
	req.Keyword = this.GetString("keyword")
	req.Begin = this.GetString("begin")
	req.End = this.GetString("end")
	req.Status = this.GetString("status")
	req.PayPtf = this.GetString("payPtf")

	if rs, total, err := req.GetPaging(); err == nil {
		this.Data["json"] = ResData(constants.Success, PageData{Data: rs, Total: total})
	} else {
		this.Data["json"] = ResData(constants.DataNull, PageData{Data: rs, Total: total})
	}
	this.ServeJSON()
}
