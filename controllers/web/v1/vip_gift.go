package v1

import (
	"encoding/json"
	"fmt"
	"wx_server_go/constants"
	"wx_server_go/models"
)

type VipController struct {
	BaseController
}

// @router /checkGiftName [get]
func (this *VipController) CheckGiftName() {
	req := new(models.SeaVipgiftlist)
	req.Giftcode = this.GetString("giftCode")
	req.Giftname = this.GetString("giftName")

	if err := req.CheckGiftName(); err == nil {
		this.Data["json"] = ResData(constants.Success, "success")
	} else {
		this.Data["json"] = ResData(constants.Success, "fail")
	}
	this.ServeJSON()
}

// @router /giftOne [get]
func (this *VipController) GetOne() {
	req := new(models.SeaVipgiftlist)
	req.Giftcode = this.GetString("giftCode")

	if req.Giftcode == "" {
		this.Data["json"] = ResCode(constants.InvalidParams)
		this.ServeJSON()
		return
	}
	if res, err := req.GetOne(); err == nil {
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
func (this *VipController) GetPaging() {
	req := new(models.SeaVipgiftlist)
	req.PageSize = this.ToIntEx("size", 10)
	req.PageIndex = this.ToIntEx("page", 1)
	req.Giftname = this.GetString("giftName")
	req.Gifttype = this.GetString("giftType")
	req.Begdate = this.GetString("begin")
	req.Enddate = this.GetString("end")

	if rs, total, err := req.GetPaging(); err == nil {
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
	var v models.ReqVipgiftlist
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &v)
	//	this.ParseForm(&v)
	v.Cusid = CusId
	v.Makeperson = Userid
	fmt.Println(fmt.Sprintf("%+v", v))
	if err = v.Insert(); err == nil {
		this.Data["json"] = ResCode(constants.Success)
	} else {
		this.Data["json"] = ResCode(constants.DBError)
	}
	this.ServeJSON()
}

// @router /gift [put]
func (this *VipController) Put() {
	var v models.ReqVipgiftlist
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &v)
	//	this.ParseForm(&v)
	v.Auditperson = Userid
	if err = v.UpdateById(); err == nil {
		this.Data["json"] = ResCode(constants.Success)
	} else {
		this.Data["json"] = ResCode(constants.DBError)
	}
	this.ServeJSON()
}

// @router /gift/:giftCode [delete]
func (this *VipController) Delete() {
	item := new(models.ReqVipgiftlist)
	item.Giftcode = this.GetString(":giftCode")
	if item.Giftcode == "" {
		this.Data["json"] = ResCode(constants.InvalidParams)
	} else {
		if err := item.DeleteById(); err == nil {
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
	req := new(models.SeaVipgiftexch)
	req.PageSize = this.ToIntEx("pageSize", 10)
	req.PageIndex = this.ToIntEx("pageIndex", 1)
	req.Mbr = this.GetString("mbr")
	req.Begin = this.GetString("begin")
	req.End = this.GetString("end")
	req.Giftname = this.GetString("giftName")

	if rs, total, err := req.GetPaging(); err == nil {
		this.Data["json"] = ResData(constants.Success, PageData{Data: rs, Total: total})
	} else {
		this.Data["json"] = ResData(constants.DataNull, PageData{Data: rs, Total: total})
	}
	this.ServeJSON()
}
