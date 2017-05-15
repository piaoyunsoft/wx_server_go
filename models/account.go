package models

import (
	"time"

	"errors"

	"strconv"
	"wx_server_go/utils"
	"wx_server_go/utils/sqltool"

	"fmt"

	"github.com/ddliao/go-lib/slog"
	"github.com/go-xorm/xorm"
)

type Account struct {
	Unicode     string    `xorm:"not null pk CHAR(12)"`
	Password    string    `xorm:"VARCHAR(255)"`
	Status      string    `xorm:"not null CHAR(2)"`
	Mobile      string    `xorm:"not null VARCHAR(20)"`
	Userid      string    `xorm:"VARCHAR(20)"`
	Accountname string    `xorm:"VARCHAR(30)"`
	Fromsys     string    `xorm:"VARCHAR(6)"`
	Fromdeptid  string    `xorm:"VARCHAR(50)"`
	Remark      string    `xorm:"VARCHAR(100)"`
	Vlddtm      time.Time `xorm:"DATETIME"`
}

type ReqAccount struct {
	Unicode     string    `json:"unicode"`
	Password    string    `json:"password"`
	Status      string    `json:"status"`
	Mobile      string    `json:"mobile"`
	Userid      string    `json:"userId"`
	Accountname string    `json:"accountName"`
	Fromsys     string    `json:"-"`
	Fromdeptid  string    `json:"cusId"`
	Remark      string    `json:"remark"`
	Vlddtm      time.Time `json:"vldDtm"`
}

type SeaAccount struct {
	SeaModel
	Unicode          string `json:"unicode"`
	Password         string `json:"password"`
	Status           string `json:"status"`
	Mobile           string `json:"mobile"`
	Accountname      string `json:"accountName"`
	Mobile_Full      string `json:"mobile_full"`
	Accountname_Full string `json:"accountName_full"`
	Key              string `json:"username"`
	NewPwd           string `json:"new_pwd"`
	Old_Pwd          string `json:"old_pwd"`
}

type AccountModel struct {
	Unicode     string    `json:"unicode"`
	Password    string    `json:"password"`
	Status      string    `json:"status"`
	Mobile      string    `json:"mobile"`
	Userid      string    `json:"userId"`
	Accountname string    `json:"accountName"`
	Fromsys     string    `json:"-"`
	Fromdeptid  string    `json:"cusId"`
	Remark      string    `json:"remark"`
	Vlddtm      time.Time `json:"vldDtm"`
	CusName     string    `json:"cusName"`
}

func (this *SeaAccount) where() *xorm.Session {
	session := x.NewSession().Table("account").Alias("a")
	if this.Accountname != "" {
		session.And("a.AccountName like ?", toLike(this.Accountname))
	}
	if this.Mobile != "" {
		session.And("a.Mobile like ?", toLike(this.Mobile))
	}
	if this.Accountname_Full != "" {
		session.And("a.AccountName = ?", this.Accountname_Full)
	}
	if this.Mobile_Full != "" {
		session.And("a.Mobile = ?", this.Mobile_Full)
	}
	if this.Status != "" {
		session.And("a.Status = ?", this.Status)
	}
	if this.Unicode != "" {
		session.And("a.Unicode = ?", this.Unicode)
	}
	if this.Password != "" {
		session.And("a.Password = ?", this.Password)
	}
	if this.Key != "" {
		session.And("a.AccountName = ? or a.Mobile = ?", this.Key, this.Key)
	}
	return session.
		Join("LEFT", []string{"platcus", "b"}, "b.cusID = a.FromDeptId").
		Desc("a.Unicode")
}

func (this *SeaAccount) GetPaging() ([]AccountModel, int64, error) {
	items := make([]AccountModel, 0, this.PageSize)
	if total, err := this.getPagingSel(this.where, "a.*, b.cusName as cus_name", new(AccountModel), &items); err != nil {
		return nil, 0, err
	} else {
		return items, total, nil
	}
}

func (this *SeaAccount) GetOne() (*AccountModel, error) {
	item := new(AccountModel)
	if err := this.getOneSel(this.where, "a.*, b.cusName as cus_name", item); err != nil {
		return nil, err
	} else {
		return item, nil
	}
}

func (this *SeaAccount) Login() (*AccountModel, error) {
	this.Status = "aa"
	item, err := this.GetOne()
	if err != nil {
		return nil, err
	} else {
		return item, nil
	}
}

func (this *SeaAccount) CheckAccount() error {
	var count int64 = 0
	var err error = nil
	item := new(Account)
	if this.Unicode != "" {
		if this.Mobile_Full != "" {
			count, err = x.Where("Mobile=? and Unicode<>?", this.Mobile_Full, this.Unicode).Count(item)
		} else if this.Accountname_Full != "" {
			count, err = x.Where("AccountName=? and Unicode<>?", this.Accountname_Full, this.Unicode).Count(item)
		} else {
			err = errors.New("invalid params")
			slog.Error(err)
			return err
		}
	} else {
		if this.Mobile_Full != "" {
			count, err = x.Where("Mobile=?", this.Mobile_Full).Count(item)
		} else if this.Accountname_Full != "" {
			count, err = x.Where("AccountName=?", this.Accountname_Full).Count(item)
		} else {
			err = errors.New("invalid params")
			slog.Error(err)
			return err
		}
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

func (this *ReqAccount) Insert() error {
	item := Account(*this)
	if id, err := sqltool.NewId("account"); err == nil {
		item.Unicode = utils.Leftpad(strconv.Itoa(id), 12, 48)
		item.Password = "000000"
		_, err := x.Insert(item)
		slog.Error(err)
		return err
	} else {
		return err
	}
}

func (this *ReqAccount) UpdateById() error {
	item := Account(*this)
	_, err := x.Omit("vldDtm").ID(item.Unicode).Update(item)
	slog.Error(err)
	return err
}

func (this *ReqAccount) DeleteById() error {
	item := Account(*this)
	_, err := x.Id(item.Unicode).Delete(new(Account))
	slog.Error(err)
	return err
}

func (this *ReqAccount) ResetPwd() error {
	item := Account(*this)
	item.Password = "000000"
	_, err := x.Omit("vldDtm").ID(item.Unicode).Update(item)
	slog.Error(err)
	return err
}

func (this *SeaAccount) ChangePwd() (string, error) {
	fmt.Println(fmt.Sprintf("%+v", this))
	item := new(Account)
	count, err := x.Where("Unicode=? and Password=?", this.Unicode, this.Old_Pwd).Count(item)
	if err != nil {
		slog.Error(err)
		return "", err
	}
	if count <= 0 {
		err = errors.New("invalid user")
		slog.Error(err)
		return "原密码错误", nil
	}
	req := new(ReqAccount)
	req.Unicode = this.Unicode
	req.Password = this.NewPwd
	return "", req.UpdateById()
}
