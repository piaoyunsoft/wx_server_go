package wx

import (
	"strings"
	"wx_server_go/constants"
	. "wx_server_go/controllers/web/v1"
	. "wx_server_go/models/wx"
)

type WxSubscribeController struct {
	BaseController
}

// @Title Get pageing WxSubscribe by query
// @Description get WxSubscribe
// @Success 200 {object} models.WxSubscribe
// @router / [get]
func (this *WxSubscribeController) GetAll() {
	var query = make(map[string]string)
	var page int64 = 1
	var size int64 = 10

	if v, err := this.GetInt64("page"); err == nil {
		page = v
	}
	if v, err := this.GetInt64("size"); err == nil {
		size = v
	}
	if v := this.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.Split(cond, "=")
			if len(kv) != 2 {
				this.Data["json"] = ResCode(constants.InvalidParams)
				this.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			if kv[1] == "" {
				continue
			}
			query[k] = v
		}
	}
	if total, rs, err := ReadWxSubscribeList(query, page, size); err == nil {
		this.Data["json"] = ResData(constants.Success, PageData{Data: rs, Total: total})
	} else {
		this.Data["json"] = ResData(constants.DataNull, PageData{Data: rs, Total: total})
	}
	this.ServeJSON()
}
