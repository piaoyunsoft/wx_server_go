package models

import "github.com/go-xorm/xorm"

type Vipcls struct {
	Comid     string `xorm:"not null pk VARCHAR(60)"`
	Vipclsid  string `xorm:"not null pk VARCHAR(10)"`
	Vipclsdes string `xorm:"not null VARCHAR(60)"`
	Status    string `xorm:"not null CHAR(2)"`
}

type ReqVipcls struct {
	Comid     string `json:"comid"`
	Vipclsid  string `json:"vipclsid"`
	Vipclsdes string `json:"vipclsdes"`
	Status    string `json:"status"`
}

type SeaVipcls struct {
	SeaModel
	Comid  string `json:"comid"`
	Status string `json:"status"`
}

type VipclsModel struct {
	Comid     string `json:"comid"`
	Vipclsid  string `json:"vipclsid"`
	Vipclsdes string `json:"vipclsdes"`
	Status    string `json:"status"`
}

func (this *SeaVipcls) where() *xorm.Session {
	session := x.NewSession().Table("vipgiftlist").Alias("a")
	if this.Comid != "" {
		session.And("a.comid = ?", this.Comid)
	}
	if this.Status != "" {
		session.And("a.status = ?", this.Status)
	}
	return session.Desc("a.createDate")
}

func (this *SeaVipcls) GetAll() ([]VipclsModel, error) {
	items := make([]VipclsModel, 0)
	if err := this.getAll(this.where, &items); err != nil {
		return nil, err
	} else {
		return items, nil
	}
}

//import (
//	"wx_server_go/utils/sqltool"
//
//	"strings"
//
//	"github.com/astaxie/beego/orm"
//)
//
//type Vipcls struct {
//	Comid     string `orm:"column(comid);size(60)" json:"comid"`
//	Vipclsid  string `orm:"column(vipclsid);pk;size(10)" json:"vipclsid"`
//	Vipclsdes string `orm:"column(vipclsdes);size(60)" json:"vipclsdes"`
//	Status    string `orm:"column(status);size(2)" json:"status"`
//}
//
//func init() {
//	orm.RegisterModel(new(Vipcls))
//}
//
//func (u *Vipcls) TableUnique() [][]string {
//	return [][]string{
//		[]string{"comid", "vipclsid"},
//	}
//}
//
//func GetVipCls(query map[string]string) (res []Vipcls, err error) {
//	filterFunc := func(qs orm.QuerySeter) orm.QuerySeter {
//		for k, v := range query {
//			k = strings.Replace(k, ".", "__", -1)
//			if k == "status" {
//				qs = qs.Filter(k, v)
//			} else if k == "comid" {
//				qs = qs.Filter(k, v)
//			}
//		}
//		return qs
//	}
//
//	if err := sqltool.Query_QS(new(Vipcls), filterFunc, &res); err == nil {
//		return res, nil
//	} else {
//		return nil, err
//	}
//}
