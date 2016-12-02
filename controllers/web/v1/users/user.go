package users

import (
	"encoding/json"
	"fmt"
	"strconv"
	"wx_server_go/constants"
	. "wx_server_go/controllers/web/v1"
	"wx_server_go/utils"

	"github.com/astaxie/beego/logs"
)

type UserController struct {
	BaseController
}

type User struct {
	Id       int64
	UserName string `json:"username"`
	Password string `json:"password"`
}

// @Title 登陆
// @Description create user
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 constants.Success
// @Failure 200 constants.DBError
// @router /login [post]
func (this *UserController) Post() {
	//	username := this.GetString("username")
	//	password := this.GetString("password")

	var user User
	json.Unmarshal(this.Ctx.Input.RequestBody, &user)
	fmt.Println(user)
	if user.UserName == "" || user.Password == "" {
		this.Data["json"] = ResCode(constants.InvalidParams)
		this.ServeJSON()
		return
	}

	if user.UserName == "admin" && user.Password == "000000" {
		user.Id = 111

		if token, err := utils.CreateToken(strconv.FormatInt(user.Id, 10), ""); err == nil {
			this.Data["json"] = ResData(constants.Success, token)
		} else {
			logs.Error(err)
			this.Data["json"] = ResCode(constants.LoginFail)
		}
	} else {
		this.Data["json"] = ResCode(constants.LoginFail)
	}
	this.ServeJSON()
}
