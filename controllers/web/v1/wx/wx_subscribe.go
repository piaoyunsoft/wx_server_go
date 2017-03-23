package wx

//import (
//	"encoding/json"
//	"io/ioutil"
//	"net/http"
//	"pt_server/utils"
//	"strings"
//	"wx_server_go/constants"
//	"wx_server_go/controllers/api"
//	. "wx_server_go/controllers/web/v1"
//	. "wx_server_go/models/wx"
//
//	"fmt"
//
//	"github.com/astaxie/beego"
//)
//
//type WxSubscribeController struct {
//	BaseController
//}
//
//// @Title Get pageing WxSubscribe by query
//// @Description get WxSubscribe
//// @Success 200 {object} models.WxSubscribe
//// @router / [get]
//func (this *WxSubscribeController) GetAll() {
//	var query = make(map[string]string)
//	var page int64 = 1
//	var size int64 = 10
//
//	if v, err := this.GetInt64("page"); err == nil {
//		page = v
//	}
//	if v, err := this.GetInt64("size"); err == nil {
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
//	if total, rs, err := ReadWxSubscribeList(query, page, size); err == nil {
//		this.Data["json"] = ResData(constants.Success, PageData{Data: rs, Total: total})
//	} else {
//		this.Data["json"] = ResData(constants.DataNull, PageData{Data: rs, Total: total})
//	}
//	this.ServeJSON()
//}
//
//// @router /apply [get]
//func (this *WxSubscribeController) GetApplyCards() {
//	var query = make(map[string]string)
//	var page int64 = 1
//	var size int64 = 10
//
//	if v, err := this.GetInt64("page"); err == nil {
//		page = v
//	}
//	if v, err := this.GetInt64("size"); err == nil {
//		size = v
//	}
//	if key := this.GetString("key"); key != "" {
//		query["key"] = this.GetString("key")
//	}
//	query["status"] = "np"
//
//	if total, rs, err := ReadWxSubscribeList(query, page, size); err == nil {
//		this.Data["json"] = ResData(constants.Success, PageData{Data: rs, Total: total})
//	} else {
//		this.Data["json"] = ResData(constants.DataNull, PageData{Data: rs, Total: total})
//	}
//	this.ServeJSON()
//}
//
//// @router / [put]
//func (this *WxSubscribeController) Put() {
//	var v Wxsubscribe
//	err := json.Unmarshal(this.Ctx.Input.RequestBody, &v)
//	//	this.ParseForm(&v)
//	if err = UpdateSubscribeByUID(&v); err == nil {
//		this.Data["json"] = v1.ResCode(constants.Success)
//	} else {
//		this.Data["json"] = v1.ResCode(constants.DBError)
//	}
//	this.ServeJSON()
//}
//
//// @router /bind [put]
//func (this *WxSubscribeController) BindCardByUID() {
//	var v Wxsubscribe
//	json.Unmarshal(this.Ctx.Input.RequestBody, &v)
//	//	this.ParseForm(&v)
//	if rs, err := BindCardByUID(&v); err == nil {
//		this.Data["json"] = v1.ResData(constants.Success, rs)
//	} else {
//		this.Data["json"] = v1.ResCode(constants.DBError)
//	}
//	this.ServeJSON()
//}
//
//// @router /checkMbrid [get]
//func (this *WxSubscribeController) CheckMbrId() {
//	mbrID := this.GetString("mbrId")
//	serverAddress := beego.AppConfig.String("serveraddr")
//	url := serverAddress + "wxopenapi/CheckMbrID?mbrID=" + mbrID
//
//	resp, err := http.Get(url)
//	if err != nil {
//		utils.Error(err)
//		this.Data["json"] = v1.ResData(constants.Success, "fail")
//		return
//	}
//
//	defer resp.Body.Close()
//	body, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		utils.Error(err)
//		this.Data["json"] = v1.ResData(constants.Success, "fail")
//		return
//	}
//	utils.Info(fmt.Sprintf("url:%s rs:%s", url, string(body)))
//	if string(body) != "success" {
//		this.Data["json"] = v1.ResData(constants.Success, "fail")
//	} else {
//		this.Data["json"] = v1.ResData(constants.Success, "success")
//	}
//	this.ServeJSON()
//}
