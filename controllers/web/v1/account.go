package v1

import (
	"encoding/json"
	"wx_server_go/constants"
	"wx_server_go/utils"

	"wx_server_go/models"
)

type AccountController struct {
	BaseController
}

// @router / [get]
func (this *AccountController) GetAll() {
	req := new(models.SeaAccount)
	req.PageSize = this.ToIntEx("size", 10)
	req.PageIndex = this.ToIntEx("page", 1)
	req.Accountname = this.GetString("accountName")
	req.Mobile = this.GetString("mobile")
	req.Status = this.GetString("status")

	if rs, total, err := req.GetPaging(); err == nil {
		this.Data["json"] = ResData(constants.Success, PageData{Data: rs, Total: total})
	} else {
		this.Data["json"] = ResData(constants.DataNull, rs)
	}
	this.ServeJSON()
}

// @router /:unicode [get]
func (this *AccountController) GetOne() {
	req := new(models.SeaAccount)
	req.Unicode = this.GetString(":unicode")

	if req.Unicode == "" {
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

// @router /check [get]
func (this *AccountController) CheckAccount() {
	req := new(models.SeaAccount)
	req.Unicode = this.GetString("unicode")
	req.Mobile_Full = this.GetString("mobile")
	req.Accountname_Full = this.GetString("accountName")

	if err := req.CheckAccount(); err == nil {
		this.Data["json"] = ResData(constants.Success, "success")
	} else {
		this.Data["json"] = ResData(constants.DataNull, "fail")
	}
	this.ServeJSON()
}

// @router /login [post]
func (this *AccountController) Login() {
	req := new(models.SeaAccount)
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &req)

	if err != nil || req.Key == "" || req.Password == "" {
		this.Data["json"] = ResCode(constants.InvalidParams)
		this.ServeJSON()
		return
	}
	if res, err := req.Login(); err == nil {
		if token, err := utils.CreateToken(res.Unicode, res.Fromdeptid); err == nil {
			loginRs := make(map[string]interface{})
			loginRs["token"] = token
			this.Data["json"] = ResData(constants.Success, loginRs)
		} else {
			utils.Error(err)
			this.Data["json"] = ResCode(constants.LoginFail)
		}
	} else {
		this.Data["json"] = ResCode(constants.DBError)
	}
	this.ServeJSON()
}

// @router / [post]
func (this *AccountController) Post() {
	req := new(models.ReqAccount)
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &req)
	//	this.ParseForm(&v)
	if err = req.Insert(); err == nil {
		this.Data["json"] = ResCode(constants.Success)
	} else {
		this.Data["json"] = ResCode(constants.DBError)
	}
	this.ServeJSON()
}

// @router / [put]
func (this *AccountController) Put() {
	req := new(models.ReqAccount)
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &req)
	//	this.ParseForm(&v)
	if err = req.UpdateById(); err == nil {
		this.Data["json"] = ResCode(constants.Success)
	} else {
		this.Data["json"] = ResCode(constants.DBError)
	}
	this.ServeJSON()
}

// @router /:unicode [delete]
func (this *AccountController) Delete() {
	req := new(models.ReqAccount)
	req.Unicode = this.GetString(":unicode")

	if req.Unicode == "" {
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

// @router /resetPwd [post]
func (this *AccountController) ResetPwd() {
	req := new(models.ReqAccount)
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &req)
	//	this.ParseForm(&v)
	if err = req.ResetPwd(); err == nil {
		this.Data["json"] = ResCode(constants.Success)
	} else {
		this.Data["json"] = ResCode(constants.DBError)
	}
	this.ServeJSON()
}
