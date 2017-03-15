package models

import (
	"fmt"
	"pt_server/utils"
	"strings"
	"wx_server_go/utils/sqltool"

	"github.com/astaxie/beego/orm"
	"github.com/zheng-ji/goSnowFlake"
)

type Wxchargelist struct {
	Id         string  `orm:"column(id);pk" json:"id"`
	Name       string  `orm:"column(name);size(60)" json:"name"`
	ComID      string  `orm:"column(comID);size(60)" json:"com_id"`
	Vipclsid   string  `orm:"column(vipclsid);size(10)" json:"vipclsid"`
	PayAmt     float64 `orm:"column(payAmt);digits(10);decimals(2)" json:"pay_amt"`
	GetRealAmt float64 `orm:"column(getRealAmt);null;digits(10);decimals(2)" json:"get_real_amt"`
	GetGiftAmt float64 `orm:"column(getGiftAmt);null;digits(10);decimals(2)" json:"get_gift_amt"`
	Status     string  `orm:"column(status);size(2);null" json:"status"`
}

type WxchargelistDtl struct {
	Id         string  `orm:"column(id);pk" json:"id"`
	Name       string  `orm:"column(name);size(60)" json:"name"`
	ComID      string  `orm:"column(comID);size(60)" json:"com_id"`
	Vipclsid   string  `orm:"column(vipclsid);size(10)" json:"vipclsid"`
	PayAmt     float64 `orm:"column(payAmt);digits(10);decimals(2)" json:"pay_amt"`
	GetRealAmt float64 `orm:"column(getRealAmt);null;digits(10);decimals(2)" json:"get_real_amt"`
	GetGiftAmt float64 `orm:"column(getGiftAmt);null;digits(10);decimals(2)" json:"get_gift_amt"`
	Status     string  `orm:"column(status);size(2);null" json:"status"`
	ComName    string  `orm:"column(comName)" json:"com_name"`
	VipClsName string  `orm:"column(vipClsName)" json:"vip_cls_name"`
}

func (t *Wxchargelist) TableName() string {
	return "wxchargelist"
}

func init() {
	orm.RegisterModel(new(Wxchargelist))
}

func CheckChargeName(name string, id string) bool {
	var query = make(map[string]string)
	if id != "" {
		query["id"] = id
	}
	query["name"] = name
	if count, result, err := GetPageCharge(query, 1, 2); err == nil {
		if count > 1 {
			return false
		} else if count > 0 {
			if result[0].Id == id {
				return true
			} else {
				return false
			}
		} else {
			return true
		}
	} else {
		return false
	}
}

func GetCharge(id string) (interface{}, error) {
	var query = make(map[string]string)
	query["id"] = id
	if count, results, err := GetPageCharge(query, 1, 1); count > 0 && err == nil {
		return results[0], err
	} else {
		return nil, err
	}
}

func GetPageCharge(query map[string]string, page int, limit int) (total int64, res []WxchargelistDtl, err error) {
	qb := sqltool.GetQueryBuilder()
	_w := "1=1"
	for k, v := range query {
		k = strings.Replace(k, ".", "__", -1)
		if k == "name" {
			_w += " and name like '%" + sqltool.SqlFormat(v) + "%'"
		}
	}
	qb.Select("a.*", "b.cusName as comName", "c.itemname as vipClsName").From("wxchargelist as a").
		LeftJoin("platcus as b").On("(b.cusID = a.comID)").
		LeftJoin("vipcls as c").On("(c.comid = a.comID and c.vipclsid = a.vipclsid)").
		Where(_w).
		OrderBy("a.id").
		Desc()

	if total, err = sqltool.PageQuery_QB(qb, &res, page, limit); err == nil {
		return total, res, nil
	} else {
		return 0, nil, err
	}
}

func CreateWxchargelist(m *Wxchargelist) error {
	iw, err := goSnowFlake.NewIdWorker(1)
	if err != nil {
		utils.Error(err)
		return err
	}
	if id, err := iw.NextId(); err != nil {
		utils.Error(err)
		return err
	} else {
		o := orm.NewOrm()
		m.Id = fmt.Sprintf("%d", id)
		if _, err = o.Insert(m); err == nil {
			return nil
		} else {
			utils.Error(err)
			return err
		}
	}
}

func UpdateWxchargelist(item *Wxchargelist) error {
	o := orm.NewOrm()
	_, err := o.Update(item)
	if err != nil {
		utils.Error(err)
	}
	return err
}

func DelWxchargelist(id string) error {
	o := orm.NewOrm()
	filter := Wxchargelist{Id: id}
	_, err := o.Delete(&filter)
	return err
}
