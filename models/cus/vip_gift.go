package cus

import (
	"strings"
	"time"
	"wx_server_go/models"
	"wx_server_go/utils/sqltool"

	"github.com/astaxie/beego/orm"
)

type VipGift struct {
	GiftCode    string    `orm:"column(giftCode);pk;" json:"giftCode"`
	CusID       string    `orm:"column(cusID)" json:"cusID"`
	GiftType    string    `orm:"column(giftType)" json:"giftType"`
	GiftAmt     float32   `orm:"column(giftAmt)" json:"giftAmt"`
	GiftName    string    `orm:"column(giftName)" json:"giftName"`
	GetWay      string    `orm:"column(getWay)" json:"getWay"`
	GetBrief    string    `orm:"column(getBrief)" json:"getBrief"`
	VldDays     string    `orm:"column(vldDays)" json:"vldDays"`
	GiftPic     string    `orm:"column(giftPic)" json:"giftPic"`
	ScoreNeed   int       `orm:"column(scoreNeed)" json:"scoreNeed"`
	StkQty      float32   `orm:"column(stkQty)" json:"stkQty"`
	DeptID      string    `orm:"column(DeptID)" json:"DeptID"`
	BegDate     time.Time `orm:"column(begDate)" json:"begDate"`
	EndDate     time.Time `orm:"column(endDate)" json:"endDate"`
	Status      string    `orm:"column(status)" json:"status"`
	MakePerson  string    `orm:"column(makePerson)" json:"makePerson"`
	MakeDate    time.Time `orm:"column(makeDate)" json:"makeDate"`
	AuditPerson string    `orm:"column(auditPerson)" json:"auditPerson"`
	AuditDate   time.Time `orm:"column(auditDate)" json:"auditDate"`
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
