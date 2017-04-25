package models

import (
	"time"

	"github.com/go-xorm/xorm"
)

type Wxchargeodr struct {
	Odrid       string    `xorm:"not null pk VARCHAR(64)" json:"odrID"`
	Payptf      string    `xorm:"VARCHAR(12)" json:"payPtf"`
	Subuid      string    `xorm:"VARCHAR(48)" json:"Subuid"`
	Comid       string    `xorm:"VARCHAR(60)" json:"comId"`
	Wxopenid    string    `xorm:"VARCHAR(40)" json:"wxOpenId"`
	Mbrid       string    `xorm:"VARCHAR(64)" json:"mbrId"`
	Amt         string    `xorm:"DECIMAL(10,2)" json:"amt"`
	Payptfodrid string    `xorm:"VARCHAR(64)" json:"payPtfOdrID"`
	Paytime     time.Time `xorm:"DATETIME" json:"payTime"`
	Status      string    `xorm:"CHAR(2)" json:"status"`
	Errmsg      string    `xorm:"VARCHAR(128)" json:"errMsg"`
	Createdate  time.Time `xorm:"DATETIME" json:"createDate"`
	Changedate  time.Time `xorm:"DATETIME" json:"changeDate"`
	Getrealamt  string    `xorm:"DECIMAL(10,2)" json:"getRealAmt"`
	Getgiftamt  string    `xorm:"DECIMAL(10,2)" json:"getGiftAmt"`
}

type SeaWxchargeodr struct {
	SeaModel
	Keyword string `json:"keyword"`
	Begin   string `json:"begin"`
	End     string `json:"end"`
	Status  string `json:"status"`
	PayPtf  string `json:"payPtf"`
}

type WxchargeodrModel struct {
	Odrid       string    `json:"odrID"`
	Payptf      string    `json:"payPtf"`
	Subuid      string    `json:"Subuid"`
	Comid       string    `json:"comId"`
	Wxopenid    string    `json:"wxOpenId"`
	Mbrid       string    `json:"mbrId"`
	Amt         string    `json:"amt"`
	Payptfodrid string    `json:"payPtfOdrID"`
	Paytime     time.Time `json:"payTime"`
	Status      string    `json:"status"`
	Errmsg      string    `json:"errMsg"`
	Createdate  time.Time `json:"createDate"`
	Changedate  time.Time `json:"changeDate"`
	Getrealamt  string    `json:"getRealAmt"`
	Getgiftamt  string    `json:"getGiftAmt"`
	NickName    string    `json:"nickName"`
}

func (this *SeaWxchargeodr) where() *xorm.Session {
	session := x.NewSession().Table("wxchargeodr").Alias("a")
	if this.Keyword != "" {
		session.And("wxsubscribe.wxNickName like ?", toLike(this.Keyword))
	}
	if this.Begin != "" {
		session.And("a.createDate >= ?", this.Begin)
	}
	if this.End != "" {
		session.And("a.createDate <= ?", this.End)
	}
	if this.PayPtf != "" {
		session.And("a.payPtf = ?", this.PayPtf)
	}
	if this.Status != "" {
		session.And("a.status = ?", this.Status)
	}
	return session.Join("LEFT", "wxsubscribe", "wxsubscribe.wxOpenId = a.wxOpenID and wxsubscribe.comID=a.comId").Desc("a.createDate")
}

func (this *SeaWxchargeodr) GetPaging() ([]WxchargeodrModel, int64, error) {
	items := make([]WxchargeodrModel, 0, this.PageSize)
	if total, err := this.getPagingSel(this.where, "a.*, wxsubscribe.wxNickName as nick_name", new(WxchargeodrModel), &items); err == nil {
		return items, total, nil
	} else {
		return nil, 0, err
	}
}
