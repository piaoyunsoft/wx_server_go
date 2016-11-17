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
	RcdType   string  `orm:"column(RcdType)"`
	LONGITUDE float32 `orm:"column(LONGITUDE)"`
	LATITUDE  float32 `orm:"column(LATITUDE)"`
	REGION_ID string  `orm:"column(REGION_ID)"`
	PARK_TYPE string  `orm:"column(PARK_TYPE)"`
	BERTH_MAX string  `orm:"column(BERTH_MAX)"`
	BERTH_RES string  `orm:"column(BERTH_RES)"`
	ORDER_NUM string  `orm:"column(ORDER_NUM)"`
	ADDRESS   string  `orm:"column(ADDRESS)"`
}

func (this *ParkView) TableName() string {
	return models.TableName("v_park_gis")
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
		//		qs = qs.Filter(k, v)
		if k == "REGION_ID" {
			cond = cond.And(k+"__icontains", v)
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
