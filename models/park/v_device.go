package park

import (
	"strings"
	"wx_server_go/models"

	"github.com/astaxie/beego/orm"
)

type DeviceView struct {
	Id         int64   `orm:"column(id)" json:"-"`
	Park_ID    string  `orm:"column(Park_ID)" json:"parkID"`
	PARK_NAME  string  `orm:"column(PARK_NAME)" json:"parkName"`
	DevTypeDes string  `orm:"column(devTypeDes)" json:"devTypeDes"`
	REGION_ID  string  `orm:"column(REGION_ID)" json:"regionID"`
	LONGITUDE  float32 `orm:"column(LONGITUDE)" json:"lng"`
	LATITUDE   float32 `orm:"column(LATITUDE)" json:"lat"`
}

func (this *DeviceView) TableName() string {
	return models.TableName("v_device")
}

func init() {
	orm.RegisterModel(new(DeviceView))
}

func GetDeviceViewList(query map[string]string) (res []DeviceView, err error) {
	var devices []DeviceView
	o := orm.NewOrm()
	qs := o.QueryTable(new(DeviceView))
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
	qs = qs.OrderBy("id")

	if _, err = qs.All(&devices); err == nil {
		return devices, nil
	} else {
		return nil, err
	}
}
