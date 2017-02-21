package cus

import (
	"fmt"
	"strings"
	"time"
	"wx_server_go/models"
	"wx_server_go/utils"
	"wx_server_go/utils/sqltool"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/zheng-ji/goSnowFlake"
)

//type Time time.Time

//const (
//	timeFormart = "2006-01-02 15:04:05"
//)

//func (t *time.Time) UnmarshalJSON(data []byte) (err error) {
//	now, err := time.ParseInLocation(`"`+timeFormart+`"`, string(data), time.Local)
//	*t = Time(now)
//	return
//}

//func (t time.Time) MarshalJSON() ([]byte, error) {
//	b := make([]byte, 0, len(timeFormart)+2)
//	b = append(b, '"')
//	b = time.Time(t).AppendFormat(b, timeFormart)
//	b = append(b, '"')
//	return b, nil
//}

//func (t time.Time) String() string {
//	return time.Time(t).Format(timeFormart)
//}

type VipGift struct {
	GiftCode string  `orm:"column(giftCode);pk;" json:"giftCode"`
	CusID    string  `orm:"column(cusID)" json:"-"`
	GiftType string  `orm:"column(giftType)" json:"giftType"`
	GiftAmt  float64 `orm:"column(giftAmt)" json:"giftAmt"`
	GiftName string  `orm:"column(giftName)" json:"giftName"`
	GetWay   string  `orm:"column(getWay)" json:"getWay"`
	GetBrief string  `orm:"column(getBrief)" json:"getBrief"`
	//	VldDays     *time.Time `orm:"column(vldDays);type(date)" json:"vldDays,omitempty"`
	VldDays     time.Time `orm:"column(vldDays);type(date);null" json:"vldDays"`
	GiftPic     string    `orm:"column(giftPic)" json:"giftPic"`
	ScoreNeed   int       `orm:"column(scoreNeed)" json:"scoreNeed"`
	StkQty      float32   `orm:"column(stkQty)" json:"stkQty"`
	DeptID      string    `orm:"column(DeptID)" json:"-"`
	BegDate     time.Time `orm:"column(begDate)" json:"begDate"`
	EndDate     time.Time `orm:"column(endDate)" json:"endDate"`
	Status      string    `orm:"column(status)" json:"status"`
	MakePerson  string    `orm:"column(makePerson)" json:"-"`
	MakeDate    time.Time `orm:"column(makeDate);auto_now_add;type(datetime);null" json:"-"`
	AuditPerson string    `orm:"column(auditPerson)" json:"-"`
	AuditDate   time.Time `orm:"column(auditDate)" json:"-"`
}

type VipGiftDtl struct {
	GiftCode     string    `orm:"column(giftCode);pk;" json:"giftCode"`
	CusID        string    `orm:"column(cusID)" json:"-"`
	GiftType     string    `orm:"column(giftType)" json:"giftType"`
	GiftAmt      float64   `orm:"column(giftAmt)" json:"giftAmt"`
	GiftName     string    `orm:"column(giftName)" json:"giftName"`
	GetWay       string    `orm:"column(getWay)" json:"getWay"`
	GetBrief     string    `orm:"column(getBrief)" json:"getBrief"`
	VldDays      time.Time `orm:"column(vldDays);type(date)" json:"vldDays"`
	GiftPic      string    `orm:"column(giftPic)" json:"giftPic"`
	ScoreNeed    int       `orm:"column(scoreNeed)" json:"scoreNeed"`
	StkQty       float32   `orm:"column(stkQty)" json:"stkQty"`
	DeptID       string    `orm:"column(DeptID)" json:"-"`
	BegDate      time.Time `orm:"column(begDate)" json:"begDate"`
	EndDate      time.Time `orm:"column(endDate)" json:"endDate"`
	Status       string    `orm:"column(status)" json:"status"`
	MakePerson   string    `orm:"column(makePerson)" json:"-"`
	MakeDate     time.Time `orm:"column(makeDate);auto_now_add;type(datetime)" json:"-"`
	AuditPerson  string    `orm:"column(auditPerson)" json:"-"`
	AuditDate    time.Time `orm:"column(auditDate)" json:"-"`
	GiftTypeName string    `orm:"column(giftTypeName)" json:"giftTypeName"`
	GetWayName   string    `orm:"column(getWayName)" json:"getWayName"`
	FullPicPath  string    `json:"fullPicPath"`
}

func (this *VipGift) TableName() string {
	return models.TableName("vipgiftlist")
}

func init() {
	orm.RegisterModel(new(VipGift))
}

func CheckGiftName(giftCode string, giftName string) bool {
	var query = make(map[string]string)
	if giftCode != "" {
		query["giftCode"] = giftCode
	}
	query["giftNameValid"] = giftName
	if count, result, err := GetPageVipGift(query, 1, 2); err == nil {
		if count > 1 {
			return false
		} else if count > 0 {
			if result[0].GiftCode == giftCode {
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

	//	o := orm.NewOrm()
	//	filter := VipGift{GiftCode: giftCode}

	//	err = o.Read(&filter)
	//	filter.GiftPic = "http://" + beego.AppConfig.String("fileserver") + filter.GiftPic
	//	if err != nil {
	//		utils.Error(err)
	//	}
	//	return filter, err
}

func GetVipGift(giftCode string) (interface{}, error) {
	var query = make(map[string]string)
	query["giftCode"] = giftCode
	if count, results, err := GetPageVipGift(query, 1, 1); count > 0 && err == nil {
		return results[0], err
	} else {
		return nil, err
	}

	//	o := orm.NewOrm()
	//	filter := VipGift{GiftCode: giftCode}

	//	err = o.Read(&filter)
	//	filter.GiftPic = "http://" + beego.AppConfig.String("fileserver") + filter.GiftPic
	//	if err != nil {
	//		utils.Error(err)
	//	}
	//	return filter, err
}

func GetPageVipGift(query map[string]string, page int, limit int) (total int64, res []VipGiftDtl, err error) {
	qb := sqltool.GetQueryBuilder()
	_w := "1=1"
	for k, v := range query {
		k = strings.Replace(k, ".", "__", -1)
		if k == "giftName" {
			_w += " and giftName like '%" + sqltool.SqlFormat(v) + "%'"
		} else if k == "begin" {
			_w += " and begDate >= '" + sqltool.SqlFormat(v) + "' "
		} else if k == "end" {
			_w += " and endDate <= '" + sqltool.SqlFormat(v) + "' "
		} else if k == "giftType" {
			_w += " and giftType = '" + sqltool.SqlFormat(v) + "' "
		} else if k == "giftCode" {
			_w += " and giftCode = '" + sqltool.SqlFormat(v) + "' "
		} else if k == "giftNameValid" {
			_w += " and giftName = '" + sqltool.SqlFormat(v) + "' "
		}
	}
	qb.Select("a.*", "b.itemname as giftTypeName", "c.itemname as getWayName").From("vipgiftlist as a").
		LeftJoin("dictitem as b").On("(b.itemcode = a.giftType and b.dictcode = '001')").
		LeftJoin("dictitem as c").On("(c.itemcode = a.getWay and c.dictcode = '002')").
		Where(_w).
		OrderBy("a.makeDate").
		Desc()

	if total, err = sqltool.PageQuery_QB(qb, &res, page, limit); err == nil {
		for k, v := range res {
			v.FullPicPath = "http://" + beego.AppConfig.String("fileserver") + v.GiftPic
			res[k] = v
		}
		return total, res, nil
	} else {
		return 0, nil, err
	}
	//	o := orm.NewOrm()
	//	qs := o.QueryTable(new(VipGift))
	//	cond := orm.NewCondition()
	//	// query k=v
	//	for k, v := range query {
	//		k = strings.Replace(k, ".", "__", -1)
	//		qs = qs.Filter(k, v)
	//		if k == "giftName" {
	//			cond = cond.And("giftName__icontains", v)
	//			//			cond = cond.AndCond(cond.And("giftName__icontains", v).Or("mobile__icontains", v))
	//		} else if k == "giftType" {
	//			cond = cond.And("giftType", v)
	//		} else if k == "begin" {
	//			cond = cond.And("begDate__gte", v)
	//		} else if k == "end" {
	//			cond = cond.And("endDate__lte", v)
	//		}
	//	}

	//	qs = qs.SetCond(cond)
	//	qs = qs.OrderBy("makeDate")

	//	if total, err = sqltool.PageQuery_QS(qs, &res, page, limit); err == nil {
	//		for k, v := range res {
	//			v.GiftPic = "http://" + beego.AppConfig.String("fileserver") + v.GiftPic
	//			res[k] = v
	//		}
	//		return total, res, nil
	//	} else {
	//		return 0, nil, err
	//	}
}

func CreateVipGift(m *VipGift) error {
	fmt.Println(m)
	o := orm.NewOrm()
	iw, err := goSnowFlake.NewIdWorker(1)
	if err != nil {
		utils.Error(err)
		return err
	}
	if id, err := iw.NextId(); err != nil {
		utils.Error(err)
		return err
	} else {
		m.GiftCode = fmt.Sprintf("%d", id)
		m.MakeDate = time.Now()
		if _, err = o.Insert(m); err == nil {
			return nil
		} else {
			utils.Error(err)
			return err
		}
	}
}

func UpdateVipGift(item *VipGift) error {
	o := orm.NewOrm()
	item.MakeDate = time.Now()
	item.AuditDate = time.Now()
	_, err := o.Update(item)
	if err != nil {
		utils.Error(err)
	}
	return err
}

func DelVipGift(giftCode string) error {
	o := orm.NewOrm()
	filter := VipGift{GiftCode: giftCode}
	_, err := o.Delete(&filter)
	return err
}
