package v1

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"wx_server_go/constants"
	"wx_server_go/controllers/api"

	"fmt"

	"wx_server_go/models"

	"github.com/astaxie/beego"
	"github.com/jdongdong/go-lib/slog"
)

type WxSubscribeController struct {
	BaseController
}

// @Title Get pageing WxSubscribe by query
// @Description get WxSubscribe
// @Success 200 {object} models.WxSubscribe
// @router / [get]
func (this *WxSubscribeController) GetAll() {
	req := new(models.SeaWxsubscribe)
	req.PageSize = this.ToIntEx("size", 10)
	req.PageIndex = this.ToIntEx("page", 1)
	req.Wxnickname = this.GetString("wxNickName")
	req.SeaMbrid = this.GetString("seaMbrId")
	req.Begin = this.GetString("begin")
	req.End = this.GetString("end")

	if rs, total, err := req.GetPaging(); err == nil {
		this.Data["json"] = ResData(constants.Success, PageData{Data: rs, Total: total})
	} else {
		this.Data["json"] = ResData(constants.DataNull, PageData{Data: rs, Total: total})
	}
	this.ServeJSON()
}

// @router /apply [get]
func (this *WxSubscribeController) GetApplyCards() {
	req := new(models.SeaWxsubscribe)
	req.PageSize = this.ToIntEx("size", 10)
	req.PageIndex = this.ToIntEx("page", 1)
	req.Key = this.GetString("key")
	req.Status = "np"

	if rs, total, err := req.GetPaging(); err == nil {
		this.Data["json"] = ResData(constants.Success, PageData{Data: rs, Total: total})
	} else {
		this.Data["json"] = ResData(constants.DataNull, PageData{Data: rs, Total: total})
	}
	this.ServeJSON()
}

// @router / [put]
func (this *WxSubscribeController) Put() {
	req := new(models.ReqWxsubscribe)
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &req)

	if err = req.UpdateById(); err == nil {
		this.Data["json"] = v1.ResCode(constants.Success)
	} else {
		this.Data["json"] = v1.ResCode(constants.DBError)
	}
	this.ServeJSON()
}

// @router /bind [put]
func (this *WxSubscribeController) BindCardByUID() {
	req := new(models.SeaWxsubscribe)
	json.Unmarshal(this.Ctx.Input.RequestBody, &req)

	if req.Uid == "" {
		this.Data["json"] = ResCode(constants.InvalidParams)
		this.ServeJSON()
		return
	}

	if rs, err := req.BindCardByUID(); err == nil {
		this.Data["json"] = v1.ResData(constants.Success, rs)
	} else {
		this.Data["json"] = v1.ResCode(constants.DBError)
	}
	this.ServeJSON()
}

// @router /checkMbrid [get]
func (this *WxSubscribeController) CheckMbrId() {
	mbrID := this.GetString("mbrId")
	serverAddress := beego.AppConfig.String("serveraddr")
	url := serverAddress + "wxopenapi/CheckMbrID?mbrID=" + mbrID

	resp, err := http.Get(url)
	if err != nil {
		slog.Error(err)
		this.Data["json"] = v1.ResData(constants.Success, "fail")
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		slog.Error(err)
		this.Data["json"] = v1.ResData(constants.Success, "fail")
		return
	}
	slog.Info(fmt.Sprintf("url:%s rs:%s", url, string(body)))
	if string(body) != "success" {
		this.Data["json"] = v1.ResData(constants.Success, "fail")
	} else {
		this.Data["json"] = v1.ResData(constants.Success, "success")
	}
	this.ServeJSON()
}

// @router /unbind [put]
func (this *WxSubscribeController) UnBindCardByUID() {
	req := new(models.SeaWxsubscribe)
	json.Unmarshal(this.Ctx.Input.RequestBody, &req)

	if req.Uid == "" {
		this.Data["json"] = ResCode(constants.InvalidParams)
		this.ServeJSON()
		return
	}

	if err := req.UnBindCardByUID(); err == nil {
		this.Data["json"] = v1.ResCode(constants.Success)
	} else {
		this.Data["json"] = v1.ResCode(constants.DBError)
	}
	this.ServeJSON()
}
