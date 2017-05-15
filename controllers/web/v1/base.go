package v1

import (
	"strings"
	"wx_server_go/constants"

	"wx_server_go/utils"

	"encoding/json"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/ddliao/go-lib/slog"
)

type BaseController struct {
	beego.Controller
	CusId  string
	UserId string
}

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
	req := make(map[string]interface{})
	this.ToJson(&req)
	slog.Trace(fmt.Sprintf("url:%s body:%+v", this.Ctx.Request.RequestURI, req))

	//	runmode := beego.AppConfig.DefaultString("runmode", "pro")
	if !strings.Contains(this.Ctx.Request.RequestURI, "/web/v1/account/login") {
		token := this.Ctx.Request.Header.Get("token")
		if token == "" {
			this.Data["json"] = ResCode(constants.InvalidToken)
			this.ServeJSON()
		}
		valid, userid, cusid := utils.ParseToken(token)
		if !valid {
			this.Data["json"] = ResCode(constants.InvalidToken)
			this.ServeJSON()
		} else {
			this.UserId = userid
			this.CusId = cusid
		}
	}
}

func (this *BaseController) ToJson(i interface{}) error {
	return json.Unmarshal(this.Ctx.Input.RequestBody, &i)
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

	slog.Trace(fmt.Sprintf("errCode:%d data:%+v msg:%s", errCode, data, msg))
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

func (this *BaseController) ToIntEx(s string, defaultVlu ...int) int {
	if rs, err := this.GetInt(s); err == nil {
		return rs
	} else if len(defaultVlu) > 0 {
		return defaultVlu[0]
	} else {
		return 0
	}

}
