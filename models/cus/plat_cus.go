package cus

import (
	"wx_server_go/models"
	"wx_server_go/utils/sqltool"

	"github.com/astaxie/beego/orm"
)

type PlatCus struct {
	CusID   string `orm:"column(cusID);pk" json:"cusID"`
	CusName string `orm:"column(cusName)" json:"cusName"`
}

const tablename = "platcus"

func (this *PlatCus) TableName() string {
	return models.TableName(tablename)
}

func init() {
	orm.RegisterModel(new(PlatCus))
}

func GetPlatCus(query map[string]string) (res []PlatCus, err error) {
	filterFunc := func(qs orm.QuerySeter) orm.QuerySeter {
		return qs
	}

	if err := sqltool.Query_QS(new(PlatCus), filterFunc, &res); err == nil {
		return res, nil
	} else {
		return nil, err
	}
}
