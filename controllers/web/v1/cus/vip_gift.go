package cus

import (
	"encoding/json"
	"wx_server_go/constants"
	. "wx_server_go/controllers/web/v1"
	. "wx_server_go/models/cus"
)

type VipController struct {
	BaseController
}

// @router /giftOne [get]
func (this *VipController) GetOne() {
	giftCode := this.GetString("giftCode")

	if giftCode == "" {
		this.Data["json"] = ResCode(constants.InvalidParams)
		this.ServeJSON()
		return
	}
	if res, err := GetVipGift(giftCode); err == nil {
		this.Data["json"] = ResData(constants.Success, res)
	} else {
		this.Data["json"] = ResCode(constants.DBError)
	}
	this.ServeJSON()
}

// @Title 获取礼品信息
// @Description 分页获取
// @Success 200 {object} models.VipGift
// @router /gift [get]
func (this *VipController) GetVipGift() {
	var query = make(map[string]string)
	var page int = 1
	var size int = 10

	if v, err := this.GetInt("page"); err == nil {
		page = v
	}
	if v, err := this.GetInt("size"); err == nil {
		size = v
	}
	if v := this.GetString("giftName"); v != "" {
		query["giftName"] = v
	}
	if v := this.GetString("giftType"); v != "" {
		query["giftType"] = v
	}
	if v := this.GetString("begin"); v != "" {
		query["begin"] = v
	}
	if v := this.GetString("end"); v != "" {
		query["end"] = v
	}

	if total, rs, err := GetPageVipGift(query, page, size); err == nil {
		this.Data["json"] = ResData(constants.Success, PageData{Data: rs, Total: total})
	} else {
		this.Data["json"] = ResData(constants.DataNull, PageData{Data: rs, Total: total})
	}
	this.ServeJSON()
}

// @Title 新增礼品
// @Description 新增礼品
// @router /gift [post]
func (this *VipController) Post() {
	var v VipGift
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &v)
	//	this.ParseForm(&v)
	if err = CreateVipGift(&v); err == nil {
		this.Data["json"] = ResCode(constants.Success)
	} else {
		this.Data["json"] = ResCode(constants.DBError)
	}
	this.ServeJSON()
}

// @router /gift [put]
func (this *VipController) Put() {
	var v VipGift
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &v)
	//	this.ParseForm(&v)
	if err = UpdateVipGift(&v); err == nil {
		this.Data["json"] = ResCode(constants.Success)
	} else {
		this.Data["json"] = ResCode(constants.DBError)
	}
	this.ServeJSON()
}

// @router /gift/:giftCode [delete]
func (this *VipController) Delete() {
	if giftCode := this.GetString(":giftCode"); giftCode == "" {
		this.Data["json"] = ResCode(constants.InvalidParams)
	} else {
		if err := DelVipGift(giftCode); err == nil {
			this.Data["json"] = ResCode(constants.Success)
		} else {
			this.Data["json"] = ResCode(constants.DBError)
		}
	}
	this.ServeJSON()
}

//-------------------------------------------------------------------
// @Title 获取礼品兑换信息
// @Description 分页获取
// @Success 200 {object} models.VipGiftExch
// @router /giftexch [get]
func (this *VipController) GetVipGiftExch() {
	var query = make(map[string]string)
	var page int = 1
	var size int = 10

	if v, err := this.GetInt("page"); err == nil {
		page = v
	}
	if v, err := this.GetInt("size"); err == nil {
		size = v
	}
	if v := this.GetString("mbr"); v != "" {
		query["mbr"] = v
	}
	if v := this.GetString("giftName"); v != "" {
		query["giftName"] = v
	}
	if v := this.GetString("begin"); v != "" {
		query["begin"] = v
	}
	if v := this.GetString("end"); v != "" {
		query["end"] = v
	}

	if total, rs, err := GetPageVipGiftExch(query, page, size); err == nil {
		this.Data["json"] = ResData(constants.Success, PageData{Data: rs, Total: total})
	} else {
		this.Data["json"] = ResData(constants.DataNull, PageData{Data: rs, Total: total})
	}
	this.ServeJSON()
}
