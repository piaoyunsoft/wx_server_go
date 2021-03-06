package v1

import (
	"encoding/json"
	"wx_server_go/constants"
	"wx_server_go/models"
)

type WxChargeListController struct {
	BaseController
}

// @router /one [get]
func (this *WxChargeListController) GetOne() {
	req := new(models.SeaWxchargelist)
	req.Id = this.GetString("id")

	if req.Id == "" {
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

// @router /page [get]
func (this *WxChargeListController) GetPaging() {
	req := new(models.SeaWxchargelist)
	req.PageSize = this.ToIntEx("size", 10)
	req.PageIndex = this.ToIntEx("page", 1)
	req.Name = this.GetString("name")
	req.Comid = this.CusId

	if rs, total, err := req.GetPaging(); err == nil {
		this.Data["json"] = ResData(constants.Success, PageData{Data: rs, Total: total})
	} else {
		this.Data["json"] = ResData(constants.DataNull, PageData{Data: rs, Total: total})
	}
	this.ServeJSON()
}

// @router / [post]
func (this *WxChargeListController) Post() {
	req := new(models.ReqWxchargelist)
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &req)
	req.Comid = this.CusId
	if err = req.Insert(); err == nil {
		this.Data["json"] = ResCode(constants.Success)
	} else {
		this.Data["json"] = ResCode(constants.DBError)
	}
	this.ServeJSON()
}

// @router / [put]
func (this *WxChargeListController) Put() {
	req := new(models.ReqWxchargelist)
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &req)

	if err = req.UpdateById(); err == nil {
		this.Data["json"] = ResCode(constants.Success)
	} else {
		this.Data["json"] = ResCode(constants.DBError)
	}
	this.ServeJSON()
}

// @router / [delete]
func (this *WxChargeListController) Delete() {
	req := new(models.ReqWxchargelist)
	req.Id = this.GetString("id")

	if req.Id == "" {
		this.Data["json"] = ResCode(constants.InvalidParams)
	} else {
		if err := req.DeleteById(); err == nil {
			this.Data["json"] = ResCode(constants.Success)
		} else {
			this.Data["json"] = ResCode(constants.DBError)
		}
	}
	this.ServeJSON()
}

// @router /checkChargeName [get]
func (this *WxChargeListController) CheckChargeName() {
	req := new(models.SeaWxchargelist)
	req.Id = this.GetString("id")
	req.Name = this.GetString("name")
	req.Comid = this.CusId

	if err := req.CheckChargeName(); err == nil {
		this.Data["json"] = ResData(constants.Success, "success")
	} else {
		this.Data["json"] = ResData(constants.Success, "fail")
	}
	this.ServeJSON()
}

// @router /checkChargeAmt [get]
func (this *WxChargeListController) CheckChargeAmt() {
	req := new(models.SeaWxchargelist)
	req.Id = this.GetString("id")
	req.Payamt = this.GetString("amt")
	req.Vipclsid = this.GetString("vipclsid")
	req.Comid = this.CusId

	if err := req.CheckChargeAmt(); err == nil {
		this.Data["json"] = ResData(constants.Success, "success")
	} else {
		this.Data["json"] = ResData(constants.Success, "fail")
	}
	this.ServeJSON()
}

//import (
//	"encoding/json"
//	"wx_server_go/constants"
//	"wx_server_go/models"
//)
//
//type WxChargeListController struct {
//	BaseController
//}
//
//// @router /one [get]
//func (this *WxChargeListController) GetOne() {
//	id := this.GetString("id")
//
//	if id == "" {
//		this.Data["json"] = ResCode(constants.InvalidParams)
//		this.ServeJSON()
//		return
//	}
//	if res, err := models.GetCharge(id); err == nil {
//		this.Data["json"] = ResData(constants.Success, res)
//	} else {
//		this.Data["json"] = ResCode(constants.DBError)
//	}
//	this.ServeJSON()
//}
//
//// @router /page [get]
//func (this *WxChargeListController) GetPaging() {
//	var query = make(map[string]string)
//	var page int = 1
//	var size int = 10
//
//	if v, err := this.GetInt("page"); err == nil {
//		page = v
//	}
//	if v, err := this.GetInt("size"); err == nil {
//		size = v
//	}
//	if v := this.GetString("name"); v != "" {
//		query["name"] = v
//	}
//	query["comID"] = CusId
//
//	if total, rs, err := models.GetPageCharge(query, page, size); err == nil {
//		this.Data["json"] = ResData(constants.Success, PageData{Data: rs, Total: total})
//	} else {
//		this.Data["json"] = ResData(constants.DataNull, PageData{Data: rs, Total: total})
//	}
//	this.ServeJSON()
//}
//
//// @router / [post]
//func (this *WxChargeListController) Post() {
//	var v models.Wxchargelist
//	err := json.Unmarshal(this.Ctx.Input.RequestBody, &v)
//	//	this.ParseForm(&v)
//	v.ComID = CusId
//	if err = models.CreateWxchargelist(&v); err == nil {
//		this.Data["json"] = ResCode(constants.Success)
//	} else {
//		this.Data["json"] = ResCode(constants.DBError)
//	}
//	this.ServeJSON()
//}
//
//// @router / [put]
//func (this *WxChargeListController) Put() {
//	var v models.Wxchargelist
//	err := json.Unmarshal(this.Ctx.Input.RequestBody, &v)
//	//	this.ParseForm(&v)
//	if err = models.UpdateWxchargelist(&v); err == nil {
//		this.Data["json"] = ResCode(constants.Success)
//	} else {
//		this.Data["json"] = ResCode(constants.DBError)
//	}
//	this.ServeJSON()
//}
//
//// @router / [delete]
//func (this *WxChargeListController) Delete() {
//	if id := this.GetString("id"); id == "" {
//		this.Data["json"] = ResCode(constants.InvalidParams)
//	} else {
//		if err := models.DelWxchargelist(id); err == nil {
//			this.Data["json"] = ResCode(constants.Success)
//		} else {
//			this.Data["json"] = ResCode(constants.DBError)
//		}
//	}
//	this.ServeJSON()
//}
//
//// @router /checkChargeName [get]
//func (this *WxChargeListController) CheckChargeName() {
//	id := this.GetString("id")
//	name := this.GetString("name")
//
//	if flag := models.CheckChargeName(name, CusId, id); flag {
//		this.Data["json"] = ResData(constants.Success, "success")
//	} else {
//		this.Data["json"] = ResData(constants.Success, "fail")
//	}
//	this.ServeJSON()
//}
//
//// @router /checkChargeAmt [get]
//func (this *WxChargeListController) CheckChargeAmt() {
//	id := this.GetString("id")
//	amt, _ := this.GetFloat("amt", 0)
//
//	if flag := models.CheckChargeAmt(amt, CusId, id); flag {
//		this.Data["json"] = ResData(constants.Success, "success")
//	} else {
//		this.Data["json"] = ResData(constants.Success, "fail")
//	}
//	this.ServeJSON()
//}
