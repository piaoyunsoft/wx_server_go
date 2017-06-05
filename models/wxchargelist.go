package models

import (
	"errors"

	"fmt"

	"github.com/go-xorm/xorm"
	"github.com/jdongdong/go-lib/slog"
	"github.com/jdongdong/go-lib/tool"
)

type Wxchargelist struct {
	Id         string `xorm:"not null pk VARCHAR(32)"`
	Name       string `xorm:"not null VARCHAR(60)"`
	Comid      string `xorm:"not null VARCHAR(60)"`
	Vipclsid   string `xorm:"not null VARCHAR(10)"`
	Payamt     string `xorm:"not null DECIMAL(10,2)"`
	Getrealamt string `xorm:"DECIMAL(10,2)"`
	Getgiftamt string `xorm:"DECIMAL(10,2)"`
	Status     string `xorm:"CHAR(2)"`
}

type ReqWxchargelist struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Comid      string `json:"com_id"`
	Vipclsid   string `json:"vipclsid"`
	Payamt     string `json:"pay_amt"`
	Getrealamt string `json:"get_real_amt"`
	Getgiftamt string `json:"get_gift_amt"`
	Status     string `json:"status"`
}

type SeaWxchargelist struct {
	SeaModel
	Id       string `json:"id"`
	Name     string `json:"name"`
	Comid    string `json:"com_id"`
	Payamt   string `json:"pay_amt"`
	Vipclsid string `json:"vipclsid"`
}

type WxchargelistModel struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Comid      string `json:"com_id"`
	Vipclsid   string `json:"vipclsid"`
	Payamt     string `json:"pay_amt"`
	Getrealamt string `json:"get_real_amt"`
	Getgiftamt string `json:"get_gift_amt"`
	Status     string `json:"status"`
	ComName    string ` json:"com_name"`
	VipClsName string `json:"vip_cls_name"`
}

func (this *SeaWxchargelist) where() *xorm.Session {
	session := x.NewSession().Table("wxchargelist").Alias("a")
	if this.Id != "" {
		session.And("a.id = ?", this.Id)
	}
	if this.Name != "" {
		session.And("a.name like ?", toLike(this.Name))
	}
	if this.Comid != "" {
		session.And("a.comID = ?", this.Comid)
	}
	if this.Payamt != "" {
		session.And("a.payAmt = ?", this.Payamt)
	}
	return session.
		Join("LEFT", []string{"platcus", "b"}, "b.cusID = a.comID").
		Join("LEFT", []string{"vipcls", "c"}, "c.comid = a.comID and c.vipclsid = a.vipclsid").
		Desc("a.id")
}

func (this *SeaWxchargelist) GetOne() (*WxchargelistModel, error) {
	item := new(WxchargelistModel)
	if err := this.getOneSel(this.where, "a.*, b.cusName as com_name, c.vipclsdes as vip_cls_name", item); err != nil {
		return nil, err
	} else {
		return item, nil
	}
}

func (this *SeaWxchargelist) GetPaging() ([]WxchargelistModel, int64, error) {
	items := make([]WxchargelistModel, 0, this.PageSize)
	if total, err := this.getPagingSel(this.where, "a.*, b.cusName as com_name, c.vipclsdes as vip_cls_name", new(WxchargelistModel), &items); err != nil {
		return nil, 0, err
	} else {
		return items, total, nil
	}
}

func (this *SeaWxchargelist) CheckChargeName() error {
	var count int64 = 0
	var err error = nil
	item := new(Wxchargelist)
	if this.Id != "" {
		count, err = x.Where("name=? and id<>? and comID=?", this.Name, this.Id, this.Comid).Count(item)
	} else {
		count, err = x.Where("name=? and comID=?", this.Name, this.Comid).Count(item)
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

func (this *SeaWxchargelist) CheckChargeAmt() error {
	var count int64 = 0
	var err error = nil
	item := new(Wxchargelist)
	if this.Id != "" {
		count, err = x.Where("payAmt=? and id<>? and comID=? and vipclsid=?", this.Payamt, this.Id, this.Comid, this.Vipclsid).Count(item)
	} else {
		count, err = x.Where("payAmt=? and comID=? and vipclsid=?", this.Payamt, this.Comid, this.Vipclsid).Count(item)
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

func (this *ReqWxchargelist) Insert() error {
	item := Wxchargelist(*this)
	item.Id = tool.NewStrID()
	item.Getrealamt = item.Payamt
	_, err := x.Insert(item)
	slog.Error(err)
	return err
}

func (this *ReqWxchargelist) UpdateById() error {
	item := Wxchargelist(*this)
	fmt.Println(fmt.Sprintf("%+v", item))
	item.Getrealamt = item.Payamt
	_, err := x.ID(item.Id).Update(item)
	slog.Error(err)
	return err
}

func (this *ReqWxchargelist) DeleteById() error {
	item := Wxchargelist(*this)
	_, err := x.Id(item.Id).Delete(new(Wxchargelist))
	slog.Error(err)
	return err
}

//import (
//	"fmt"
//	"pt_server/utils"
//	"strings"
//	"wx_server_go/utils/sqltool"
//
//	"github.com/astaxie/beego/orm"
//	"github.com/zheng-ji/goSnowFlake"
//)
//
//type Wxchargelist struct {
//	Id         string  `orm:"column(id);pk" json:"id"`
//	Name       string  `orm:"column(name);size(60)" json:"name"`
//	ComID      string  `orm:"column(comID);size(60)" json:"com_id"`
//	Vipclsid   string  `orm:"column(vipclsid);size(10)" json:"vipclsid"`
//	PayAmt     float64 `orm:"column(payAmt);digits(10);decimals(2)" json:"pay_amt"`
//	GetRealAmt float64 `orm:"column(getRealAmt);null;digits(10);decimals(2)" json:"get_real_amt"`
//	GetGiftAmt float64 `orm:"column(getGiftAmt);null;digits(10);decimals(2)" json:"get_gift_amt"`
//	Status     string  `orm:"column(status);size(2);null" json:"status"`
//}
//
//type WxchargelistDtl struct {
//	Id         string  `orm:"column(id);pk" json:"id"`
//	Name       string  `orm:"column(name);size(60)" json:"name"`
//	ComID      string  `orm:"column(comID);size(60)" json:"com_id"`
//	Vipclsid   string  `orm:"column(vipclsid);size(10)" json:"vipclsid"`
//	PayAmt     float64 `orm:"column(payAmt);digits(10);decimals(2)" json:"pay_amt"`
//	GetRealAmt float64 `orm:"column(getRealAmt);null;digits(10);decimals(2)" json:"get_real_amt"`
//	GetGiftAmt float64 `orm:"column(getGiftAmt);null;digits(10);decimals(2)" json:"get_gift_amt"`
//	Status     string  `orm:"column(status);size(2);null" json:"status"`
//	ComName    string  `orm:"column(comName)" json:"com_name"`
//	VipClsName string  `orm:"column(vipClsName)" json:"vip_cls_name"`
//}
//
//func (t *Wxchargelist) TableName() string {
//	return "wxchargelist"
//}
//
//func init() {
//	orm.RegisterModel(new(Wxchargelist))
//}
//
//func CheckChargeName(name string, CusId string, id string) bool {
//	var query = make(map[string]string)
//	if id != "" {
//		query["id"] = id
//	}
//	query["checkname"] = name
//	query["comID"] = CusId
//	if count, result, err := GetPageCharge(query, 1, 2); err == nil {
//		if count > 1 {
//			return false
//		} else if count > 0 {
//			if result[0].Id == id {
//				return true
//			} else {
//				return false
//			}
//		} else {
//			return true
//		}
//	} else {
//		return false
//	}
//}
//
//func CheckChargeAmt(amt float64, CusId string, id string) bool {
//	var query = make(map[string]string)
//	if id != "" {
//		query["id"] = id
//	}
//	query["payAmt"] = tool.ToString(amt)
//	query["comID"] = CusId
//	if count, result, err := GetPageCharge(query, 1, 2); err == nil {
//		if count > 1 {
//			return false
//		} else if count > 0 {
//			if result[0].Id == id {
//				return true
//			} else {
//				return false
//			}
//		} else {
//			return true
//		}
//	} else {
//		return false
//	}
//}
//
//func GetCharge(id string) (interface{}, error) {
//	var query = make(map[string]string)
//	query["id"] = id
//	if count, results, err := GetPageCharge(query, 1, 1); count > 0 && err == nil {
//		return results[0], err
//	} else {
//		return nil, err
//	}
//}
//
//func GetPageCharge(query map[string]string, page int, limit int) (total int64, res []WxchargelistDtl, err error) {
//	qb := sqltool.GetQueryBuilder()
//	_w := "1=1"
//	for k, v := range query {
//		k = strings.Replace(k, ".", "__", -1)
//		if k == "name" {
//			_w += " and a.name like '%" + sqltool.SqlFormat(v) + "%'"
//		} else if k == "comID" {
//			_w += " and a.comID = '" + sqltool.SqlFormat(v) + "'"
//		} else if k == "id" {
//			_w += " and a.id = '" + sqltool.SqlFormat(v) + "'"
//		} else if k == "checkname" {
//			_w += " and a.name = '" + sqltool.SqlFormat(v) + "'"
//		} else if k == "payAmt" {
//			_w += " and a.payAmt = " + sqltool.SqlFormat(v)
//		}
//	}
//	qb.Select("a.*", "b.cusName as comName", "c.vipclsdes as vipClsName").From("wxchargelist as a").
//		LeftJoin("platcus as b").On("(b.cusID = a.comID)").
//		LeftJoin("vipcls as c").On("(c.comid = a.comID and c.vipclsid = a.vipclsid)").
//		Where(_w).
//		OrderBy("a.id").
//		Desc()
//
//	if total, err = sqltool.PageQuery_QB(qb, &res, page, limit); err == nil {
//		return total, res, nil
//	} else {
//		return 0, nil, err
//	}
//}
//
//func CreateWxchargelist(m *Wxchargelist) error {
//	fmt.Println(m, 1111111111)
//	iw, err := goSnowFlake.NewIdWorker(1)
//	if err != nil {
//		utils.Error(err)
//		return err
//	}
//	if id, err := iw.NextId(); err != nil {
//		utils.Error(err)
//		return err
//	} else {
//		o := orm.NewOrm()
//		m.Id = fmt.Sprintf("%d", id)
//		if _, err = o.Insert(m); err == nil {
//			return nil
//		} else {
//			utils.Error(err)
//			return err
//		}
//	}
//}
//
//func UpdateWxchargelist(item *Wxchargelist) error {
//	o := orm.NewOrm()
//	_, err := o.Update(item)
//	if err != nil {
//		utils.Error(err)
//	}
//	return err
//}
//
//func DelWxchargelist(id string) error {
//	o := orm.NewOrm()
//	filter := Wxchargelist{Id: id}
//	_, err := o.Delete(&filter)
//	return err
//}
