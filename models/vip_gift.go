package models

import (
	"time"

	"errors"

	"wx_server_go/utils/sqltool"

	"fmt"

	"github.com/astaxie/beego"
	"github.com/go-xorm/xorm"
	"github.com/jdongdong/go-lib/slog"
)

type Vipgiftlist struct {
	Giftcode    string    `xorm:"not null pk VARCHAR(32)"`
	Cusid       string    `xorm:"VARCHAR(10)"`
	Gifttype    string    `xorm:"VARCHAR(10)"`
	Giftamt     string    `xorm:"DECIMAL(10,2)"`
	Giftname    string    `xorm:"VARCHAR(255)"`
	Getway      string    `xorm:"VARCHAR(10)"`
	Getbrief    string    `xorm:"VARCHAR(255)"`
	Vlddays     time.Time `xorm:"DATE"`
	Giftpic     string    `xorm:"VARCHAR(255)"`
	Scoreneed   int       `xorm:"INT(11)"`
	Stkqty      string    `xorm:"DECIMAL(10,2)"`
	Deptid      string    `xorm:"VARCHAR(12)"`
	Begdate     time.Time `xorm:"DATETIME"`
	Enddate     time.Time `xorm:"DATETIME"`
	Status      string    `xorm:"CHAR(2)"`
	Makeperson  string    `xorm:"VARCHAR(20)"`
	Makedate    time.Time `xorm:"DATETIME"`
	Auditperson string    `xorm:"VARCHAR(20)"`
	Auditdate   time.Time `xorm:"DATETIME"`
}

type ReqVipgiftlist struct {
	Giftcode    string    `json:"giftCode"`
	Cusid       string    `json:"-"`
	Gifttype    string    `json:"giftType"`
	Giftamt     string    `json:"giftAmt"`
	Giftname    string    `json:"giftName"`
	Getway      string    `json:"getWay"`
	Getbrief    string    `json:"getBrief"`
	Vlddays     time.Time `json:"vldDays"`
	Giftpic     string    `json:"giftPic"`
	Scoreneed   int       `json:"scoreNeed"`
	Stkqty      string    `json:"stkQty"`
	Deptid      string    `json:"-"`
	Begdate     time.Time `json:"begDate"`
	Enddate     time.Time `json:"endDate"`
	Status      string    `json:"status"`
	Makeperson  string    `json:"-"`
	Makedate    time.Time `json:"-"`
	Auditperson string    `json:"-"`
	Auditdate   time.Time `json:"-"`
}

type SeaVipgiftlist struct {
	SeaModel
	Giftcode string `json:"giftCode"`
	Gifttype string `json:"giftType"`
	Giftname string `json:"giftName"`
	Begdate  string `json:"begDate"`
	Enddate  string `json:"endDate"`
}

type VipgiftlistModel struct {
	Giftcode     string    `json:"giftCode"`
	Cusid        string    `json:"-"`
	Gifttype     string    `json:"giftType"`
	Giftamt      string    `json:"giftAmt"`
	Giftname     string    `json:"giftName"`
	Getway       string    `json:"getWay"`
	Getbrief     string    `json:"getBrief"`
	Vlddays      time.Time `json:"vldDays"`
	Giftpic      string    `json:"giftPic"`
	Scoreneed    int       `json:"scoreNeed"`
	Stkqty       string    `json:"stkQty"`
	Deptid       string    `json:"-"`
	Begdate      time.Time `json:"begDate"`
	Enddate      time.Time `json:"endDate"`
	Status       string    `json:"status"`
	Makeperson   string    `json:"-"`
	Makedate     time.Time `json:"-"`
	Auditperson  string    `json:"-"`
	Auditdate    time.Time `json:"-"`
	GiftTypeName string    `json:"giftTypeName"`
	GetWayName   string    `json:"getWayName"`
	FullPicPath  string    `xorm:"-" json:"fullPicPath"`
}

func (this *SeaVipgiftlist) where() *xorm.Session {
	session := x.NewSession().Table("vipgiftlist").Alias("a")
	if this.Giftcode != "" {
		session.And("a.giftCode = ?", this.Giftcode)
	}
	if this.Begdate != "" {
		session.And("a.begDate >= ?", this.Begdate)
	}
	if this.Enddate != "" {
		session.And("a.begDate <= ?", this.Enddate)
	}
	if this.Gifttype != "" {
		session.And("a.giftType = ?", this.Gifttype)
	}
	if this.Giftname != "" {
		session.And("a.giftName like ?", toLike(this.Giftname))
	}
	return session.
		Join("LEFT", []string{"dictitem", "b"}, "b.itemcode = a.giftType and b.dictcode = '001'").
		Join("LEFT", []string{"dictitem", "c"}, "c.itemcode = a.getWay and c.dictcode = '002'").
		Desc("a.makeDate")
}

func (this *SeaVipgiftlist) GetPaging() ([]VipgiftlistModel, int64, error) {
	items := make([]VipgiftlistModel, 0, this.PageSize)
	if total, err := this.getPagingSel(this.where, "a.*, b.itemname as gift_type_name, c.itemname as get_way_name", new(VipgiftlistModel), &items); err != nil {
		return nil, 0, err
	} else {
		for k, v := range items {
			v.FullPicPath = "http://" + beego.AppConfig.String("fileserver") + v.Giftpic
			items[k] = v
		}
		return items, total, nil
	}
}

func (this *SeaVipgiftlist) GetOne() (*VipgiftlistModel, error) {
	item := new(VipgiftlistModel)
	if err := this.getOneSel(this.where, "a.*, b.itemname as gift_type_name, c.itemname as get_way_name", item); err != nil {
		return nil, err
	} else {
		item.FullPicPath = "http://" + beego.AppConfig.String("fileserver") + item.Giftpic
		return item, nil
	}
}

func (this *SeaVipgiftlist) CheckGiftName() error {
	var count int64 = 0
	var err error = nil
	item := new(Vipgiftlist)
	if this.Giftcode != "" {
		count, err = x.Where("giftName=? and giftCode<>?", this.Giftname, this.Giftcode).Count(item)
	} else {
		count, err = x.Where("giftName=?", this.Giftname).Count(item)
	}
	if err != nil {
		slog.Error(err)
		return err
	}
	if count > 0 {
		err = errors.New("exist data")
		slog.Error(err)
		return err
	}
	return nil
}

func (this *ReqVipgiftlist) Insert() error {
	item := Vipgiftlist(*this)
	id, err := sqltool.NewId("vipgiftlist")
	if err != nil {
		slog.Error(err)
		return err
	} else {
		item.Giftcode = fmt.Sprintf("g%07d", id)
		item.Makedate = time.Now()
		_, err := x.Insert(item)
		slog.Error(err)
		return err
	}
}

func (this *ReqVipgiftlist) UpdateById() error {
	item := Vipgiftlist(*this)
	item.Auditdate = time.Now()
	_, err := x.Omit("makeDate").ID(item.Giftcode).Update(item)
	slog.Error(err)
	return err
}

func (this *ReqVipgiftlist) DeleteById() error {
	item := Vipgiftlist(*this)
	_, err := x.Id(item.Giftcode).Delete(new(Vipgiftlist))
	slog.Error(err)
	return err
}
