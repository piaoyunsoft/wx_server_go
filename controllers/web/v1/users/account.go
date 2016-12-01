package users

import (
	"encoding/json"
	"fmt"
	"wx_server_go/constants"
	"wx_server_go/controllers/api"
	"wx_server_go/models/users"
)

type AccountController struct {
	v1.BaseController
}

// @router / [get]
func (this *AccountController) GetAll() {
	query := make(map[string]string)
	page := 1
	size := 10

	if v, err := this.GetInt("page"); err == nil {
		page = v
	}
	if v, err := this.GetInt("size"); err == nil {
		size = v
	}
	if v := this.GetString("accountName"); v != "" {
		query["accountName"] = v
	}
	if v := this.GetString("mobile"); v != "" {
		query["mobile"] = v
	}
	if v := this.GetString("status"); v != "" {
		query["status"] = v
	}
	if total, rs, err := users.GetAccounts(query, page, size); err == nil {
		this.Data["json"] = v1.ResData(constants.Success, v1.PageData{Data: rs, Total: total})
	} else {
		this.Data["json"] = v1.ResData(constants.DataNull, rs)
	}
	this.ServeJSON()
}

// @router / [post]
func (this *AccountController) Post() {
	var v users.Account
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &v)
	fmt.Println(v)
	//	this.ParseForm(&v)
	if err = users.CreateAccount(&v); err == nil {
		this.Data["json"] = v1.ResCode(constants.Success)
	} else {
		this.Data["json"] = v1.ResCode(constants.DBError)
	}
	this.ServeJSON()
}

// @router / [put]
func (this *AccountController) Put() {
	var v users.Account
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &v)
	//	this.ParseForm(&v)
	if err = users.UpdateAccount(&v); err == nil {
		this.Data["json"] = v1.ResCode(constants.Success)
	} else {
		this.Data["json"] = v1.ResCode(constants.DBError)
	}
	this.ServeJSON()
}

// @router /:unicode [delete]
func (this *AccountController) Delete() {
	if unicode := this.GetString(":id"); unicode == "" {
		this.Data["json"] = v1.ResCode(constants.InvalidParams)
	} else {
		if err := users.DelAccount(unicode); err == nil {
			this.Data["json"] = v1.ResCode(constants.Success)
		} else {
			this.Data["json"] = v1.ResCode(constants.DBError)
		}
	}
	this.ServeJSON()
}
