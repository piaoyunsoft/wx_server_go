package users

//import (
//	"encoding/json"
//	"wx_server_go/constants"
//	"wx_server_go/controllers/web/v1"
//	"wx_server_go/models/users"
//	"wx_server_go/utils"
//
//	"github.com/astaxie/beego/logs"
//)
//
//type AccountController struct {
//	v1.BaseController
//}
//
//// @router / [get]
//func (this *AccountController) GetAll() {
//	query := make(map[string]string)
//	page := 1
//	size := 10
//
//	if v, err := this.GetInt("page"); err == nil {
//		page = v
//	}
//	if v, err := this.GetInt("size"); err == nil {
//		size = v
//	}
//	if v := this.GetString("accountName"); v != "" {
//		query["accountName"] = v
//	}
//	if v := this.GetString("mobile"); v != "" {
//		query["mobile"] = v
//	}
//	if v := this.GetString("status"); v != "" {
//		query["status"] = v
//	}
//	if total, rs, err := users.GetPageAccounts(query, page, size); err == nil {
//		this.Data["json"] = v1.ResData(constants.Success, v1.PageData{Data: rs, Total: total})
//	} else {
//		this.Data["json"] = v1.ResData(constants.DataNull, rs)
//	}
//	this.ServeJSON()
//}
//
//// @router /:unicode [get]
//func (this *AccountController) GetOne() {
//	unicode := this.GetString(":unicode")
//
//	if unicode == "" {
//		this.Data["json"] = v1.ResCode(constants.InvalidParams)
//		this.ServeJSON()
//		return
//	}
//	if res, err := users.GetAccountByUnicode(unicode); err == nil {
//		this.Data["json"] = v1.ResData(constants.Success, res)
//	} else {
//		this.Data["json"] = v1.ResCode(constants.DBError)
//	}
//	this.ServeJSON()
//}
//
//// @router /check [get]
//func (this *AccountController) CheckAccount() {
//	query := make(map[string]string)
//
//	if v := this.GetString("accountName"); v != "" {
//		query["accountName"] = v
//	}
//	if v := this.GetString("mobile"); v != "" {
//		query["mobile"] = v
//	}
//	if v := this.GetString("unicode"); v != "" {
//		query["unicode"] = v
//	}
//	if rs, err := users.CheckAccount(query); err == nil {
//		if rs {
//			this.Data["json"] = v1.ResData(constants.Success, "success")
//		} else {
//			this.Data["json"] = v1.ResData(constants.Success, "fail")
//		}
//	} else {
//		this.Data["json"] = v1.ResData(constants.DataNull, "fail")
//	}
//	this.ServeJSON()
//}
//
//// @router /login [post]
//func (this *AccountController) Login() {
//	var v users.Account
//	err := json.Unmarshal(this.Ctx.Input.RequestBody, &v)
//
//	if err != nil || v.AccountName == "" || v.Password == "" {
//		this.Data["json"] = v1.ResCode(constants.InvalidParams)
//		this.ServeJSON()
//		return
//	}
//	if res, err := users.Login(v.AccountName, v.Password); err == nil {
//		rs, ok := res.(users.Account)
//		if ok {
//			if token, err := utils.CreateToken(rs.Unicode, rs.FromDeptId); err == nil {
//				//				query := make(map[string]string)
//				//				query["cusID"] = v1.CusId
//				//				cus, _ := cus.GetPlatCus(query)
//				loginRs := make(map[string]interface{})
//				loginRs["token"] = token
//				//				loginRs["cus"] = cus
//				this.Data["json"] = v1.ResData(constants.Success, loginRs)
//			} else {
//				logs.Error(err)
//				this.Data["json"] = v1.ResCode(constants.LoginFail)
//			}
//		} else {
//			this.Data["json"] = v1.ResCode(constants.LoginFail)
//		}
//
//	} else {
//		this.Data["json"] = v1.ResCode(constants.DBError)
//	}
//	this.ServeJSON()
//}
//
//// @router / [post]
//func (this *AccountController) Post() {
//	var v users.Account
//	err := json.Unmarshal(this.Ctx.Input.RequestBody, &v)
//	//	this.ParseForm(&v)
//	if err = users.CreateAccount(&v); err == nil {
//		this.Data["json"] = v1.ResCode(constants.Success)
//	} else {
//		this.Data["json"] = v1.ResCode(constants.DBError)
//	}
//	this.ServeJSON()
//}
//
//// @router / [put]
//func (this *AccountController) Put() {
//	var v users.Account
//	err := json.Unmarshal(this.Ctx.Input.RequestBody, &v)
//	//	this.ParseForm(&v)
//	if err = users.UpdateAccount(&v); err == nil {
//		this.Data["json"] = v1.ResCode(constants.Success)
//	} else {
//		this.Data["json"] = v1.ResCode(constants.DBError)
//	}
//	this.ServeJSON()
//}
//
//// @router /:unicode [delete]
//func (this *AccountController) Delete() {
//	if unicode := this.GetString(":id"); unicode == "" {
//		this.Data["json"] = v1.ResCode(constants.InvalidParams)
//	} else {
//		if err := users.DelAccount(unicode); err == nil {
//			this.Data["json"] = v1.ResCode(constants.Success)
//		} else {
//			this.Data["json"] = v1.ResCode(constants.DBError)
//		}
//	}
//	this.ServeJSON()
//}
//
//// @router /resetPwd [post]
//func (this *AccountController) ResetPwd() {
//	var v users.Account
//	err := json.Unmarshal(this.Ctx.Input.RequestBody, &v)
//	//	this.ParseForm(&v)
//	if err = users.ResetPwd(&v); err == nil {
//		this.Data["json"] = v1.ResCode(constants.Success)
//	} else {
//		this.Data["json"] = v1.ResCode(constants.DBError)
//	}
//	this.ServeJSON()
//}
