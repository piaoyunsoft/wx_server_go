package v1

import (
	"wx_server_go/constants"
	"wx_server_go/models"
)

type CouponController struct {
	BaseController
}

// @router /coupon [get]
func (this *CouponController) GetCoupon() {
	req := new(models.SeaCoupon)
	req.PageSize = this.ToIntEx("size", 10)
	req.PageIndex = this.ToIntEx("page", 1)
	req.Title = this.GetString("title")
	req.ComID = this.CusId

	if total, rs, err := req.GetCoupon(); err == nil {
		this.Data["json"] = ResData(constants.Success, PageData{Data: rs, Total: total})
	} else {
		this.Data["json"] = ResData(constants.DataNull, nil)
	}
	this.ServeJSON()
}

// @router /couponItem [get]
func (this *CouponController) GetCouponItem() {
	req := new(models.SeaCouponItem)
	req.CouponID = this.GetString("couponID")

	if rs, err := req.GetCouponItem(); err == nil {
		this.Data["json"] = ResData(constants.Success, rs)
	} else {
		this.Data["json"] = ResData(constants.DataNull, nil)
	}
	this.ServeJSON()
}
