package models

import (
	"encoding/json"
	"time"

	"github.com/astaxie/beego"
	"github.com/ddliao/go-lib/slog"
	"github.com/ddliao/go-lib/tool"
)

type SendMsgModel struct {
	ClientId     string `json:"clientId"`
	Timestamp    int64  `json:"timestamp"`
	Sign         string `json:"sign"`
	MbrId        string `json:"mbrId"`
	MbrOwner     string `json:"mbrOwner"`
	TplCode      string `json:"tplCode"`
	FormattedMsg string `json:"formattedMsg"`
}

//开通成功通知
type TplMsg_OPENTM405904851 struct {
	First    string `json:"first"`
	Keyword1 string `json:"keyword1"`
	Keyword2 string `json:"keyword2"`
	Remark   string `json:"remark"`
}

//订单发货通知
type TplMsg_OPENTM405840986 struct {
	First    string `json:"first"`
	Keyword1 string `json:"keyword1"`
	Keyword2 string `json:"keyword2"`
	Keyword3 string `json:"keyword3"`
	Remark   string `json:"remark"`
}

func NewSendMsgModel(mbrID, comID string) *SendMsgModel {
	model := new(SendMsgModel)
	model.MbrOwner = comID
	model.MbrId = mbrID
	model.ClientId = "web"
	model.Sign = ""
	model.Timestamp = time.Now().Unix()
	return model
}

//发送<订单发货通知>消息
func (this *SendMsgModel) SendTplMsg_OPENTM405840986(company, bllno, detail string) error {
	tpl := new(TplMsg_OPENTM405840986)
	tpl.First = "您好，天赐名店积分换礼-礼品已发货了"
	tpl.Keyword1 = company
	tpl.Keyword2 = bllno
	tpl.Keyword3 = detail
	tpl.Remark = "有任何问题请致电0817-2255339联系我们，谢谢"
	str, _ := json.Marshal(tpl)

	this.FormattedMsg = string(str)
	this.TplCode = "OPENTM405840986"

	serverAddress := beego.AppConfig.String("serveraddr")
	url := serverAddress + "wxopenapi/SendTplMsg"
	_, err := doPost(url, this)
	if err != nil {
		slog.Error(err)
		return err
	}
	return nil
}

//发送<开通成功通知>消息
func (this *SendMsgModel) SendTplMsg_OPENTM405904851() error {
	tpl := new(TplMsg_OPENTM405904851)
	tpl.First = "您好，会员卡绑定成功"
	tpl.Keyword1 = this.MbrId
	tpl.Keyword2 = tool.TimeToStr(time.Now(), "MM月dd日 HH时mm分")
	tpl.Remark = "恭喜您成为天赐时代绿卡会员，自此您将获得我们为您提供的购物积分、积分换礼、特价通知、停车折价、免费杂志、周年庆受邀等服务"
	str, _ := json.Marshal(tpl)

	this.FormattedMsg = string(str)
	this.TplCode = "OPENTM405904851"

	serverAddress := beego.AppConfig.String("serveraddr")
	url := serverAddress + "wxopenapi/SendTplMsg"
	_, err := doPost(url, this)
	if err != nil {
		slog.Error(err)
		return err
	}

	return nil
}
