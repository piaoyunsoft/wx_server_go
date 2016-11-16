package park

import (
	"strings"
	"wx_server_go/models"
	"wx_server_go/utils/sqltool"

	"github.com/astaxie/beego/orm"
)

type ParkView struct {
	PARK_ID   string  `orm:"column(PARK_ID);pk"`
	PARK_NAME string  `orm:"column(PARK_NAME)"`
	LONGITUDE float32 `orm:"column(LONGITUDE)"`
	LATITUDE  float32 `orm:"column(LATITUDE)"`
}

func (this *ParkView) TableName() string {
	return models.TableName("v_park")
}

func init() {
	orm.RegisterModel(new(ParkView))
}

func GetParkViewList(query map[string]string, page int, limit int) (total int64, res []ParkView, err error) {
	var parks []ParkView
	o := orm.NewOrm()
	qs := o.QueryTable(new(ParkView))
	cond := orm.NewCondition()
	// query k=v
	for k, v := range query {
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
		if k == "keyword" {
			cond = cond.AndCond(cond.And("mbrName__icontains", v).Or("mobile__icontains", v))
		}
	}
	qs = qs.SetCond(cond)
	qs = qs.OrderBy("PARK_ID")

	if total, err = sqltool.PageQuery_QS(qs, &parks, page, limit); err == nil {
		return total, parks, nil
	} else {
		return 0, nil, err
	}
}
