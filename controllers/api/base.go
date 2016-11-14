package v1

import (
	"wx_server_go/constants"

	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

var userid int64

type Response struct {
	ErrCode constants.ErrCode `json:"errcode"`
	ErrMsg  string            `json:"errmsg"`
	Data    interface{}       `json:"data"`
}

type PageData struct {
	Data  interface{} `json:"data"`
	Total int64       `json:"total"`
}

func (this *BaseController) Prepare() {
	//	runmode := beego.AppConfig.DefaultString("runmode", "pro")
	//	if runmode == "pro" && !strings.Contains(this.Ctx.Request.RequestURI, "/web/v1/user/login") {
	//		token := this.Ctx.Request.Header.Get("token")
	//		if token == "" {
	//			this.Data["json"] = ResCode(constants.InvalidToken)
	//			this.ServeJSON()
	//		}
	//		valid, id := utils.ParseToken(token)
	//		if !valid {
	//			this.Data["json"] = ResCode(constants.InvalidToken)
	//			this.ServeJSON()
	//		} else {
	//			userid = id
	//		}
	//	}
}

func ResBase(errCode constants.ErrCode, data interface{}, msg string) Response {
	if data == nil {
		data = ""
	} else {
		//		if reflect.ValueOf(data).IsNil() {
		//			data = ""
		//		}
		//		if reflect.TypeOf(data).String() == "[]interface {}" {
		//			if reflect.ValueOf(data).Len() == 0 {
		//				data = ""
		//			}
		//		}
		//		if reflect.TypeOf(data).String() == "[]orm.Params" {
		//			if reflect.ValueOf(data).Len() == 0 {
		//				data = ""
		//			}
		//		}
	}

	res := Response{ErrCode: errCode, Data: data, ErrMsg: msg}
	return res
}

func ResData(errCode constants.ErrCode, data interface{}) Response {
	var errmsg string
	switch errCode {
	case constants.Success:
		errmsg = ""
	case constants.InvalidParams:
		errmsg = "参数无效"
	case constants.DataNull:
		errmsg = "数据为空"
	case constants.IdExist:
		errmsg = "主键冲突"
	case constants.DBError:
		errmsg = "数据库操作异常"
	case constants.InvalidToken:
		errmsg = "登录失效"
	case constants.LoginFail:
		errmsg = "用户名或密码错误"
	}
	res := ResBase(errCode, data, errmsg)
	return res
}

func ResCode(errCode constants.ErrCode) Response {
	res := ResData(errCode, nil)
	return res
}
