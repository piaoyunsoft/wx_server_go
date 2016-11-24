package sys

import (
	"strings"
	"wx_server_go/models"

	"github.com/astaxie/beego/orm"
)

type DictItem struct {
	DictCode  string `orm:"column(dictcode)" json:"dictcode"`
	ItemCode  string `orm:"column(itemcode);pk" json:"itemcode"`
	ItemName  string `orm:"column(itemname)" json:"itemname"`
	Status    string `orm:"column(status)" json:"status"`
	IsSysDict string `orm:"column(issysdict)" json:"issysdict"`
}

func (this *DictItem) TableName() string {
	return models.TableName("dictitem")
}

func init() {
	orm.RegisterModel(new(DictItem))
}

func (u *DictItem) TableUnique() [][]string {
	return [][]string{
		[]string{"dictcode", "itemcode"},
	}
}

func GetDictItemsByDictCode(query map[string]string) (res []DictItem, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(DictItem))
	for k, v := range query {
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
	}
	if _, err = qs.All(&res); err == nil {
		return res, nil
	} else {
		return nil, err
	}
}
