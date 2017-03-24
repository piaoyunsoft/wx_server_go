package models

import (
	"time"

	"github.com/go-xorm/xorm"
)

type Vipgiftexch struct {
	Exchid        string    `xorm:"not null pk VARCHAR(32)"`
	Comid         string    `xorm:"VARCHAR(10)"`
	Mbrid         string    `xorm:"VARCHAR(32)"`
	Wxopenid      string    `xorm:"VARCHAR(40)"`
	Giftcode      string    `xorm:"not null VARCHAR(32)"`
	Giftname      string    `xorm:"VARCHAR(255)"`
	Getway        string    `xorm:"VARCHAR(10)"`
	Mailaddr      string    `xorm:"VARCHAR(255)"`
	Mailpsnname   string    `xorm:"VARCHAR(255)"`
	Mailpsnmobile string    `xorm:"VARCHAR(24)"`
	Exchqty       string    `xorm:"DECIMAL(10,2)"`
	Usescore      int       `xorm:"INT(11)"`
	Status        string    `xorm:"CHAR(2)"`
	Mailstatus    string    `xorm:"CHAR(2)"`
	Createdate    time.Time `xorm:"DATETIME"`
	Changedate    time.Time `xorm:"DATETIME"`
}

type ReqVipgiftexch struct {
	Exchid        string    `json:"exchID"`
	Comid         string    `json:"comId"`
	Mbrid         string    `json:"mbrID"`
	Wxopenid      string    `json:"wxOpenID"`
	Giftcode      string    `json:"giftCode"`
	Giftname      string    `json:"giftName"`
	Getway        string    `json:"getWay"`
	Mailaddr      string    `json:"mailAddr"`
	Mailpsnname   string    `json:"mailPsnName"`
	Mailpsnmobile string    `json:"mailPsnMobile"`
	Exchqty       string    `json:"exchQty"`
	Usescore      int       `json:"useScore"`
	Status        string    `json:"status"`
	Mailstatus    string    `json:"mailStatus"`
	Createdate    time.Time `json:"createDate"`
	Changedate    time.Time `json:"changeDate"`
}

type SeaVipgiftexch struct {
	SeaModel
	Giftname string `json:"giftName"`
	Begin    string `json:"begin"`
	End      string `json:"end"`
	Mbr      string `json:"mbr"`
}

type VipgiftexchModel struct {
	Exchid         string    `json:"exchID"`
	Comid          string    `json:"comId"`
	Mbrid          string    `json:"mbrID"`
	Wxopenid       string    `json:"wxOpenID"`
	Giftcode       string    `json:"giftCode"`
	Giftname       string    `json:"giftName"`
	Getway         string    `json:"getWay"`
	Mailaddr       string    `json:"mailAddr"`
	Mailpsnname    string    `json:"mailPsnName"`
	Mailpsnmobile  string    `json:"mailPsnMobile"`
	Exchqty        string    `json:"exchQty"`
	Usescore       int       `json:"useScore"`
	Status         string    `json:"status"`
	Mailstatus     string    `json:"mailStatus"`
	Createdate     time.Time `json:"createDate"`
	Changedate     time.Time `json:"changeDate"`
	NickName       string    `json:"nickName"`
	GetWayName     string    `json:"getWayName"`
	MailStatusName string    `json:"mailStatusName"`
}

func (this *SeaVipgiftexch) where() *xorm.Session {
	session := x.NewSession().Table("vipgiftexch").Alias("a")
	if this.Giftname != "" {
		session.And("a.giftName like ?", toLike(this.Giftname))
	}
	if this.Begin != "" {
		session.And("a.createDate >= ?", this.Begin)
	}
	if this.End != "" {
		session.And("a.createDate <= ?", this.End)
	}
	if this.Mbr != "" {
		session.And("b.wxNickName like ?", toLike(this.Mbr))
	}
	return session.
		Join("LEFT", []string{"wxsubscribe", "b"}, "b.wxOpenId = a.wxOpenID and b.comID = a.comId").
		Join("LEFT", []string{"dictitem", "c"}, "c.itemcode = a.getWay and c.dictcode = '002'").
		Join("LEFT", []string{"dictitem", "d"}, "d.itemcode = a.mailStatus and d.dictcode = '003'").
		Desc("a.createDate")
}

func (this *SeaVipgiftexch) GetPaging() ([]VipgiftexchModel, int64, error) {
	items := make([]VipgiftexchModel, 0, this.PageSize)
	if total, err := this.getPagingSel(this.where, "a.*, b.wxNickName as nick_name, c.itemname as get_way_name, d.itemname as mail_status_name", new(VipgiftexchModel), &items); err != nil {
		return nil, 0, err
	} else {
		return items, total, nil
	}
}
