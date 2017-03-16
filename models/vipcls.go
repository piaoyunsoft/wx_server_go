package models

import (
	"wx_server_go/utils/sqltool"

	"strings"

	"github.com/astaxie/beego/orm"
)

type Vipcls struct {
	Comid     string `orm:"column(comid);size(60)" json:"comid"`
	Vipclsid  string `orm:"column(vipclsid);pk;size(10)" json:"vipclsid"`
	Vipclsdes string `orm:"column(vipclsdes);size(60)" json:"vipclsdes"`
	Status    string `orm:"column(status);size(2)" json:"status"`
}

func init() {
	orm.RegisterModel(new(Vipcls))
}

func (u *Vipcls) TableUnique() [][]string {
	return [][]string{
		[]string{"comid", "vipclsid"},
	}
}

func GetVipCls(query map[string]string) (res []Vipcls, err error) {
	filterFunc := func(qs orm.QuerySeter) orm.QuerySeter {
		for k, v := range query {
			k = strings.Replace(k, ".", "__", -1)
			if k == "status" {
				qs = qs.Filter(k, v)
			} else if k == "comid" {
				qs = qs.Filter(k, v)
			}
		}
		return qs
	}

	if err := sqltool.Query_QS(new(Vipcls), filterFunc, &res); err == nil {
		return res, nil
	} else {
		return nil, err
	}
}
