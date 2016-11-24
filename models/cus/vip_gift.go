package cus

import (
	"fmt"
	"strings"
	"time"
	"wx_server_go/models"
	"wx_server_go/utils"
	"wx_server_go/utils/sqltool"

	"github.com/astaxie/beego/orm"
	"github.com/zheng-ji/goSnowFlake"
)

type VipGift struct {
	GiftCode    string    `orm:"column(giftCode);pk;" json:"giftCode"`
	CusID       string    `orm:"column(cusID)" json:"-"`
	GiftType    string    `orm:"column(giftType)" json:"giftType"`
	GiftAmt     float64   `orm:"column(giftAmt)" json:"giftAmt"`
	GiftName    string    `orm:"column(giftName)" json:"giftName"`
	GetWay      string    `orm:"column(getWay)" json:"getWay"`
	GetBrief    string    `orm:"column(getBrief)" json:"getBrief"`
	VldDays     time.Time `orm:"column(vldDays);type(date)" json:"vldDays"`
	GiftPic     string    `orm:"column(giftPic)" json:"giftPic"`
	ScoreNeed   int       `orm:"column(scoreNeed)" json:"scoreNeed"`
	StkQty      float32   `orm:"column(stkQty)" json:"stkQty"`
	DeptID      string    `orm:"column(DeptID)" json:"-"`
	BegDate     time.Time `orm:"column(begDate)" json:"begDate"`
	EndDate     time.Time `orm:"column(endDate)" json:"endDate"`
	Status      string    `orm:"column(status)" json:"status"`
	MakePerson  string    `orm:"column(makePerson)" json:"-"`
	MakeDate    time.Time `orm:"column(makeDate);auto_now_add;type(datetime)" json:"-"`
	AuditPerson string    `orm:"column(auditPerson)" json:"-"`
	AuditDate   time.Time `orm:"column(auditDate)" json:"-"`
}

func (this *VipGift) TableName() string {
	return models.TableName("vipgiftlist")
}

func init() {
	orm.RegisterModel(new(VipGift))
}

func GetPageVipGift(query map[string]string, page int, limit int) (total int64, res []VipGift, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(VipGift))
	cond := orm.NewCondition()
	// query k=v
	for k, v := range query {
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
		if k == "giftName" {
			cond = cond.And("giftName__icontains", v)
			//			cond = cond.AndCond(cond.And("giftName__icontains", v).Or("mobile__icontains", v))
		} else if k == "giftType" {
			cond = cond.And("giftType", v)
		} else if k == "begin" {
			cond = cond.And("begDate__gte", v)
		} else if k == "end" {
			cond = cond.And("endDate__lte", v)
		}
	}
	qs = qs.SetCond(cond)
	qs = qs.OrderBy("makeDate")

	if total, err = sqltool.PageQuery_QS(qs, &res, page, limit); err == nil {
		return total, res, nil
	} else {
		return 0, nil, err
	}
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
		if _, err = o.Insert(m); err == nil {
			return nil
		} else {
			utils.Error(err)
			return err
		}
	}
}
