package v1

import (
	"wx_server_go/constants"
	"wx_server_go/models"
)

type WxTaskController struct {
	BaseController
}

// @Title 获取微信消息任务
// @Description 分页获取
// @Success 200 {object} models.Platcuswxtask
// @router / [get]
func (this *WxTaskController) GetAll() {
	req := new(models.SeaPlatcuswxtask)
	req.PageSize = this.ToIntEx("size", 10)
	req.PageIndex = this.ToIntEx("page", 1)
	req.Begin = this.GetString("begin")
	req.End = this.GetString("end")
	req.Status = this.GetString("status")
	req.NickName = this.GetString("nickName")

	if rs, total, err := req.GetPaging(); err == nil {
		this.Data["json"] = ResData(constants.Success, PageData{Data: rs, Total: total})
	} else {
		this.Data["json"] = ResData(constants.DataNull, PageData{Data: rs, Total: total})
	}
	this.ServeJSON()
}
