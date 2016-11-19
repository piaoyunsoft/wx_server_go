package park

import (
	"strings"
	"time"
	"wx_server_go/models"

	"github.com/astaxie/beego/orm"
)

type TrafficView struct {
	Id          int64     `orm:"column(id);pk" json:"id"`
	PARK_NAME   string    `orm:"column(PARK_NAME)" json:"parkName"`
	LONGITUDE   float64   `orm:"column(LONGITUDE)" json:"lng"`
	LATITUDE    float64   `orm:"column(LATITUDE)" json:"lat"`
	CarNum      string    `orm:"column(carNum)" json:"carNum"`
	CAR_OWNER   string    `orm:"column(CAR_OWNER)" json:"userName"`
	OWNER_PHONE string    `orm:"column(OWNER_PHONE)" json:"phone"`
	ThroughTime time.Time `orm:"column(throughTime)" json:"throughTime"`
	PassType    string    `orm:"column(passType)" json:"passType"`
}

func (this *TrafficView) TableName() string {
	return models.TableName("v_trafficrecords")
}

func init() {
	orm.RegisterModel(new(TrafficView))
}

func GetTrafficViewList(query map[string]string) (res []TrafficView, err error) {
	var devices []TrafficView
	o := orm.NewOrm()
	qs := o.QueryTable(new(TrafficView))
	cond := orm.NewCondition()
	// query k=v
	for k, v := range query {
		k = strings.Replace(k, ".", "__", -1)
		if k == "carNum" {
			cond = cond.And("carNum__icontains", v)
		} else if k == "phone" {
			cond = cond.And("OWNER_PHONE__icontains", v)
		} else if k == "userName" {
			cond = cond.And("CAR_OWNER__icontains", v)
		} else if k == "begin" {
			cond = cond.And("throughTime__gte", v)
		} else if k == "end" {
			cond = cond.And("throughTime__lte", v)
		}
	}
	qs = qs.SetCond(cond)
	qs = qs.OrderBy("throughTime")

	if _, err = qs.All(&devices); err == nil {
		return devices, nil
	} else {
		return nil, err
	}
}
