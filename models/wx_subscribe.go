package models

import (
	"errors"
	"time"

	"github.com/astaxie/beego"
	"github.com/go-xorm/xorm"
	"github.com/jdongdong/go-lib/slog"
)

type Wxsubscribe struct {
	Uid               string    `xorm:"not null pk VARCHAR(48)"`
	Wxopenid          string    `xorm:"not null unique(Index_wxAttTmp) VARCHAR(40)"`
	Wxunionid         string    `xorm:"VARCHAR(48)"`
	Comwxid           string    `xorm:"not null unique(Index_wxAttTmp) VARCHAR(100)"`
	Comid             string    `xorm:"VARCHAR(60)"`
	Wxnickname        string    `xorm:"VARCHAR(100)"`
	Wxsex             string    `xorm:"CHAR(1)"`
	Subscribed        string    `xorm:"CHAR(1)"`
	Wxsubscribetime   time.Time `xorm:"not null DATETIME"`
	Wxunsubscribetime time.Time `xorm:"DATETIME"`
	Wxcountry         string    `xorm:"VARCHAR(40)"`
	Wxprovince        string    `xorm:"VARCHAR(40)"`
	Wxcity            string    `xorm:"VARCHAR(40)"`
	Wxheadimgurl      string    `xorm:"VARCHAR(255)"`
	Wxsubscribecount  int       `xorm:"INT(11)"`
	Wxbrief           string    `xorm:"VARCHAR(600)"`
	Binddate          time.Time `xorm:"DATETIME"`
	Bindway           string    `xorm:"VARCHAR(30)"`
	Mbrid             string    `xorm:"VARCHAR(32)"`
	Aduitdate         time.Time `xorm:"DATETIME"`
	Aduitperson       string    `xorm:"VARCHAR(30)"`
	Status            string    `xorm:"not null CHAR(2)"`
	Mbrname           string    `xorm:"VARCHAR(64)"`
	Mbrtype           string    `xorm:"VARCHAR(10)"`
	Mobile            string    `xorm:"VARCHAR(30)"`
	Idno              string    `xorm:"VARCHAR(20)"`
	Birthdate         string    `xorm:"CHAR(8)"`
	Addr              string    `xorm:"VARCHAR(255)"`
	Createdate        time.Time `xorm:"not null DATETIME"`
	Changedate        time.Time `xorm:"DATETIME"`
	Applybrief        string    `xorm:"VARCHAR(255)"`
	Vipclsid          string    `xorm:"VARCHAR(10)"`
}

type ReqWxsubscribe struct {
	Uid               string    `json:"uid"`
	Wxopenid          string    `json:"wxOpenId"`
	Wxunionid         string    `json:"wxUnionId"`
	Comwxid           string    `json:"comWxID"`
	Comid             string    `json:"comID"`
	Wxnickname        string    `json:"wxNickName"`
	Wxsex             string    `json:"wxSex"`
	Subscribed        string    `json:"subscribed"`
	Wxsubscribetime   time.Time `json:"wxSubscribeTime"`
	Wxunsubscribetime time.Time `json:"wxUnsubscribeTime"`
	Wxcountry         string    `json:"wxCountry"`
	Wxprovince        string    `json:"wxProvince"`
	Wxcity            string    `json:"wxCity"`
	Wxheadimgurl      string    `json:"wxHeadImgUrl"`
	Wxsubscribecount  int       `json:"wxSubscribeCount"`
	Wxbrief           string    `json:"wxBrief"`
	Binddate          time.Time `json:"BindDate"`
	Bindway           string    `json:"BindWay"`
	Mbrid             string    `json:"mbrId"`
	Aduitdate         time.Time `json:"aduitDate"`
	Aduitperson       string    `json:"aduitPerson"`
	Status            string    `json:"status"`
	Mbrname           string    `json:"mbrName"`
	Mbrtype           string    `json:"mbrType"`
	Mobile            string    `json:"mobile"`
	Idno              string    `json:"idno"`
	Birthdate         string    `json:"birthDate"`
	Addr              string    `json:"addr"`
	Createdate        time.Time `json:"createDate"`
	Changedate        time.Time `json:"changeDate"`
	Applybrief        string    `json:"apply_brief"`
	Vipclsid          string    `json:"vipClsId"`
}

type SeaWxsubscribe struct {
	SeaModel
	Uid        string `json:"uid"`
	Wxnickname string `json:"wxNickName"`
	Begin      string `json:"begin"`
	End        string `json:"end"`
	Status     string `json:"status"`
	Key        string `json:"key"`
	SeaMbrid   string `json:"seaMbrId"`
	BindBegin  string
	BindEnd    string
}

type WxsubscribeModel struct {
	Uid               string    `json:"uid"`
	Wxopenid          string    `json:"wxOpenId"`
	Wxunionid         string    `json:"wxUnionId"`
	Comwxid           string    `json:"comWxID"`
	Comid             string    `json:"comID"`
	Wxnickname        string    `json:"wxNickName"`
	Wxsex             string    `json:"wxSex"`
	Subscribed        string    `json:"subscribed"`
	Wxsubscribetime   time.Time `json:"wxSubscribeTime"`
	Wxunsubscribetime time.Time `json:"wxUnsubscribeTime"`
	Wxcountry         string    `json:"wxCountry"`
	Wxprovince        string    `json:"wxProvince"`
	Wxcity            string    `json:"wxCity"`
	Wxheadimgurl      string    `json:"wxHeadImgUrl"`
	Wxsubscribecount  int       `json:"wxSubscribeCount"`
	Wxbrief           string    `json:"wxBrief"`
	Binddate          time.Time `json:"BindDate"`
	Bindway           string    `json:"BindWay"`
	Mbrid             string    `json:"mbrId"`
	Aduitdate         time.Time `json:"aduitDate"`
	Aduitperson       string    `json:"aduitPerson"`
	Status            string    `json:"status"`
	Mbrname           string    `json:"mbrName"`
	Mbrtype           string    `json:"mbrType"`
	Mobile            string    `json:"mobile"`
	Idno              string    `json:"idno"`
	Birthdate         string    `json:"birthDate"`
	Addr              string    `json:"addr"`
	Createdate        time.Time `json:"createDate"`
	Changedate        time.Time `json:"changeDate"`
	Applybrief        string    `json:"apply_brief"`
	Vipclsid          string    `json:"vipClsId"`
}

func (this *SeaWxsubscribe) where() *xorm.Session {
	session := x.NewSession().Table("wxsubscribe").Alias("a")
	if this.Wxnickname != "" {
		session.And("a.WxNickName like ?", toLike(this.Wxnickname))
	}
	if this.Begin != "" {
		session.And("a.WxSubscribeTime >= ?", this.Begin)
	}
	if this.End != "" {
		session.And("a.WxSubscribeTime <= ?", this.End)
	}
	if this.BindBegin != "" {
		session.And("a.BindDate >= ?", this.BindBegin)
	}
	if this.BindEnd != "" {
		session.And("a.BindDate <= ?", this.BindEnd)
	}
	if this.Status != "" {
		session.And("a.status = ?", this.Status)
	}
	if this.SeaMbrid != "" {
		session.And("a.mbrId like ?", toLike(this.SeaMbrid))
	}
	if this.Key != "" {
		session.And("a.mbrName like ? or a.mobile like ?", toLike(this.Key), toLike(this.Key))
	}
	if this.Uid != "" {
		session.And("a.uid = ?", this.Uid)
	}
	return session.Desc("a.createDate")
}

func (this *SeaWxsubscribe) GetPaging() ([]WxsubscribeModel, int64, error) {
	items := make([]WxsubscribeModel, 0, this.PageSize)
	if total, err := this.getPaging(this.where, new(WxsubscribeModel), &items); err != nil {
		return nil, 0, err
	} else {
		return items, total, nil
	}
}

func (this *ReqWxsubscribe) UpdateById() error {
	item := Wxsubscribe(*this)
	item.Changedate = time.Now()
	_, err := x.Omit("wxSubscribeTime", "wxUnsubscribeTime", "aduitDate", "createDate", "BindDate").ID(item.Uid).Update(item)
	slog.Error(err)
	return err
}

type UpdateOpenIDModel struct {
	MbrID  string `json:"mbrID"`
	OpenID string `json:"openID"`
}

func (this *SeaWxsubscribe) BindCardByUID() (string, error) {
	item := new(Wxsubscribe)
	if err := this.getOne(this.where, item); err != nil {
		slog.Error(err)
		return "", err
	}

	//验证会员卡号是否有效
	serverAddress := beego.AppConfig.String("serveraddr")
	url := serverAddress + "wxopenapi/CheckMbrID?mbrID=" + item.Mbrid
	body, err := doGet(url)
	if err != nil {
		slog.Error(err)
		return "", err
	}
	if string(body) != "success" {
		err = errors.New("会员卡号无效或已被绑定")
		slog.Error(err)
		return "会员卡号无效或已被绑定", nil
	}
	//更新商场库会员openid
	url = serverAddress + "wxopenapi/UpdateOpenID"
	uptOpenID := new(UpdateOpenIDModel)
	uptOpenID.MbrID = item.Mbrid
	uptOpenID.OpenID = item.Wxopenid
	_, err = doPost(url, uptOpenID)
	if err != nil {
		slog.Error(err)
		return "", err
	}
	if string(body) != "success" {
		err = errors.New("同步商场数据失败")
		slog.Error(err)
		return "同步商场数据失败", nil
	}

	uptItem := new(Wxsubscribe)
	uptItem.Changedate = time.Now()
	uptItem.Status = "aa"
	uptItem.Bindway = "后台绑定"
	uptItem.Binddate = time.Now()
	_, err = x.Omit("wxSubscribeTime", "wxUnsubscribeTime", "aduitDate", "createDate").ID(this.Uid).Update(uptItem)
	if err != nil {
		slog.Error(err)
		return "", err
	}

	send := NewSendMsgModel(item.Mbrid, item.Comid)
	err = send.SendTplMsg_OPENTM405904851()
	if err != nil {
		return "", err
	}

	return "success", nil
}

func (this *SeaWxsubscribe) UnBindCardByUID() error {
	item := new(Wxsubscribe)
	item.Uid = this.Uid
	item.Changedate = time.Now()
	item.Status = "na"
	_, err := x.Cols("BindDate", "BindWay", "mbrId", "status", "mbrType", "changeDate", "vipclsid").ID(item.Uid).Update(item)
	slog.Error(err)
	return err
}
