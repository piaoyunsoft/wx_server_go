package wx

import (
	"strings"
	"wx_server_go/constants"
	. "wx_server_go/controllers/api/v1"
	. "wx_server_go/models/wx"
)

type WxTaskController struct {
	BaseController
}

// @Title 新增待发消息
// @Description
// @Success 200 {object}
// @router /pushmsg [post]
func (this *WxTaskController) PushMsg() {

}
