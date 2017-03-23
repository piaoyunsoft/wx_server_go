package models

import (
	"time"

	"github.com/go-xorm/xorm"
)

type Platcuswxtask struct {
	Id           int64     `xorm:"BIGINT(20)"`
	Fromcusid    string    `xorm:"VARCHAR(12)"`
	Towxopenid   string    `xorm:"not null VARCHAR(40)"`
	Tocusmbrid   string    `xorm:"VARCHAR(48)"`
	Tocusmbrname string    `xorm:"VARCHAR(200)"`
	Toaccountid  string    `xorm:"VARCHAR(48)"`
	Tplcode      string    `xorm:"VARCHAR(16)"`
	Tplid        string    `xorm:"VARCHAR(100)"`
	Type         string    `xorm:"VARCHAR(20)"`
	Msgbody      string    `xorm:"TEXT"`
	Createdtm    time.Time `xorm:"DATETIME"`
	Status       string    `xorm:"CHAR(2)"`
	Senddtm      string    `xorm:"CHAR(14)"`
	Wxmsgid      string    `xorm:"VARCHAR(50)"`
	Receivedtm   string    `xorm:"CHAR(10)"`
	Remark       string    `xorm:"VARCHAR(200)"`
	Url          string    `xorm:"VARCHAR(200)"`
}

type ReqPlatcuswxtask struct {
	Id           int64     `json:"id"`
	Fromcusid    string    `json:"FromCusID"`
	Towxopenid   string    `json:"ToWxOpenId"`
	Tocusmbrid   string    `json:"ToCusMbrId"`
	Tocusmbrname string    `json:"ToCusMbrName"`
	Toaccountid  string    `json:"ToAccountId"`
	Tplcode      string    `json:"tplCode"`
	Tplid        string    `json:"tplId"`
	Type         string    `json:"type"`
	Msgbody      string    `json:"msgBody"`
	Createdtm    time.Time `json:"createDtm"`
	Status       string    `json:"status"`
	Senddtm      string    `json:"sendDtm"`
	Wxmsgid      string    `json:"wxMsgId"`
	Receivedtm   string    `json:"receiveDtm"`
	Remark       string    `json:"remark"`
	Url          string    `json:"url"`
}

type SeaPlatcuswxtask struct {
	SeaModel
	NickName string `json:"nickName"`
	Begin    string `json:"begin"`
	End      string `json:"end"`
	Status   string `json:"status"`
}

type PlatcuswxtaskModel struct {
	Id           int64     `json:"id"`
	Fromcusid    string    `json:"FromCusID"`
	Towxopenid   string    `json:"ToWxOpenId"`
	Tocusmbrid   string    `json:"ToCusMbrId"`
	Tocusmbrname string    `json:"mbrName"`
	Toaccountid  string    `json:"ToAccountId"`
	Tplcode      string    `json:"tplCode"`
	Tplid        string    `json:"tplId"`
	Type         string    `json:"type"`
	Msgbody      string    `json:"msgBody"`
	Createdtm    time.Time `json:"createDtm"`
	Status       string    `json:"status"`
	Senddtm      string    `json:"sendDtm"`
	Wxmsgid      string    `json:"wxMsgId"`
	Receivedtm   string    `json:"receiveDtm"`
	Remark       string    `json:"remark"`
	Url          string    `json:"url"`
	NickName     string    `json:"nickName"`
	CusName      string    `json:"cusName"`
	TplName      string    `json:"tplName"`
}

func (this *SeaPlatcuswxtask) where() *xorm.Session {
	session := x.NewSession().Table("platcuswxtask").Alias("a")
	if this.NickName != "" {
		session.And("b.wxNickName like ? or a.ToCusMbrName like ?", toLike(this.NickName), toLike(this.NickName))
	}
	if this.Begin != "" {
		session.And("a.sendDtm >= ?", this.Begin)
	}
	if this.End != "" {
		session.And("a.sendDtm <= ?", this.End)
	}
	if this.Status != "" {
		session.And("a.status = ?", this.Status)
	}
	return session.
		Join("LEFT", []string{"wxsubscribe", "b"}, "b.wxOpenId = a.ToWxOpenId").
		Join("LEFT", []string{"wxtplmsgtpllib", "c"}, "c.tplCode = a.tplCode").
		Join("LEFT", []string{"platcus", "d"}, "d.cusID = a.FromCusID").
		Desc("a.createDtm")
}

func (this *SeaPlatcuswxtask) GetPaging() ([]PlatcuswxtaskModel, int64, error) {
	items := make([]PlatcuswxtaskModel, 0, this.PageSize)
	if total, err := this.getPagingSel(this.where, "a.*, b.wxNickName as nick_name, c.title as tpl_name, d.cusName as cus_name", new(PlatcuswxtaskModel), &items); err != nil {
		return nil, 0, err
	} else {
		return items, total, nil
	}
}
