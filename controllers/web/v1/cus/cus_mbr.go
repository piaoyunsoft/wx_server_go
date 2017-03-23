package cus

//import (
//	"strings"
//	"wx_server_go/constants"
//	. "wx_server_go/controllers/web/v1"
//	. "wx_server_go/models/cus"
//)
//
//type CusMbrController struct {
//	BaseController
//}
//
//// @Title 获取商户会员信息
//// @Description 分页获取
//// @Success 200 {object} models.CusMbrApi
//// @router / [get]
//func (this *CusMbrController) GetAll() {
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
//	if v := this.GetString("query"); v != "" {
//		for _, cond := range strings.Split(v, ",") {
//			kv := strings.Split(cond, "=")
//			if len(kv) != 2 {
//				this.Data["json"] = ResCode(constants.InvalidParams)
//				this.ServeJSON()
//				return
//			}
//			k, v := kv[0], kv[1]
//			if kv[1] == "" {
//				continue
//			}
//			query[k] = v
//		}
//	}
//	if total, rs, err := ReadDtl_CusMbr(query, page, size); err == nil {
//		this.Data["json"] = ResData(constants.Success, PageData{Data: rs, Total: total})
//	} else {
//		this.Data["json"] = ResData(constants.DataNull, PageData{Data: rs, Total: total})
//	}
//	this.ServeJSON()
//}
