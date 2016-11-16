package park

import (
	"OPMS/models"
	"strings"
	"wx_server_go/utils/sqltool"

	"github.com/astaxie/beego/orm"
)

type Region struct {
	ID          string `orm:"column(ID);pk"`
	REGION_NAME string `orm:"column(REGION_NAME)"`
	PARENT_ID   string `orm:"column(PARENT_ID)"`
}

type CascaderData struct {
	Value    string         `json:"value"`
	Label    string         `json:"label"`
	Children []CascaderData `json:"children"`
}

func (this *Region) TableName() string {
	return models.TableName("parking_region")
}

func init() {
	orm.RegisterModel(new(Region))
}

func GetRegionCasData() (res []CascaderData, err error) {
	var query = make(map[string]string)
	var page int = 1
	var size int = 1000

	if _, rs, err := GetRegionList(query, page, size); err == nil {
		var data []CascaderData
		count := 0
		for _, item := range rs {
			if item.PARENT_ID == "" {
				count++
			}
		}
		data = make([]CascaderData, count)
		for k, item := range rs {
			if item.PARENT_ID == "" {
				temp := CascaderData{Value: item.ID, Label: item.REGION_NAME}
				data[k] = temp
			}
		}
		return data, nil
	} else {
		return nil, err
	}
}

func GetRegionList(query map[string]string, page int, limit int) (total int64, res []Region, err error) {
	var regions []Region
	o := orm.NewOrm()
	qs := o.QueryTable(new(Region))
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
	qs = qs.OrderBy("ID")

	if total, err = sqltool.PageQuery_QS(qs, &regions, page, limit); err == nil {
		return total, regions, nil
	} else {
		return 0, nil, err
	}
}
