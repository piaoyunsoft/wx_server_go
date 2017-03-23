package sys

//import (
//	"encoding/json"
//	"wx_server_go/constants"
//	"wx_server_go/controllers/web/v1"
//	"wx_server_go/models/sys"
//)
//
//type ModuleController struct {
//	v1.BaseController
//}
//
//// @Title 获取模块信息
//// @Description
//// @Success 200 {object} models.PowerModule
//// @router / [get]
//func (this *ModuleController) GetAll() {
//	query := make(map[string]string)
//
//	//	if v := this.GetString("dictcode"); v != "" {
//	//		query["dictcode"] = v
//	//	}
//	if rs, err := sys.GetPowerModules(query); err == nil {
//		this.Data["json"] = v1.ResData(constants.Success, rs)
//	} else {
//		this.Data["json"] = v1.ResData(constants.DataNull, rs)
//	}
//	this.ServeJSON()
//}
//
//// @router /valid [get]
//func (this *ModuleController) GetValidModule() {
//	query := make(map[string]string)
//
//	//	if v := this.GetString("dictcode"); v != "" {
//	//		query["dictcode"] = v
//	//	}
//	query["status"] = "aa"
//	if rs, err := sys.GetPowerModules(query); err == nil {
//		this.Data["json"] = v1.ResData(constants.Success, rs)
//	} else {
//		this.Data["json"] = v1.ResData(constants.DataNull, rs)
//	}
//	this.ServeJSON()
//}
//
//// @router / [post]
//func (this *ModuleController) Post() {
//	var v sys.PowerModule
//	err := json.Unmarshal(this.Ctx.Input.RequestBody, &v)
//	//	this.ParseForm(&v)
//	if err = sys.CreatePowerModule(&v); err == nil {
//		this.Data["json"] = v1.ResCode(constants.Success)
//	} else {
//		this.Data["json"] = v1.ResCode(constants.DBError)
//	}
//	this.ServeJSON()
//}
//
//// @router / [put]
//func (this *ModuleController) Put() {
//	var v sys.PowerModule
//	err := json.Unmarshal(this.Ctx.Input.RequestBody, &v)
//	//	this.ParseForm(&v)
//	if err = sys.UpdatePowerModule(&v); err == nil {
//		this.Data["json"] = v1.ResCode(constants.Success)
//	} else {
//		this.Data["json"] = v1.ResCode(constants.DBError)
//	}
//	this.ServeJSON()
//}
//
//// @router /:id [delete]
//func (this *ModuleController) Delete() {
//	if id, err := this.GetInt(":id"); err != nil {
//		this.Data["json"] = v1.ResCode(constants.InvalidParams)
//	} else {
//		if err := sys.DelPowerModule(id); err == nil {
//			this.Data["json"] = v1.ResCode(constants.Success)
//		} else {
//			this.Data["json"] = v1.ResCode(constants.DBError)
//		}
//	}
//	this.ServeJSON()
//}
