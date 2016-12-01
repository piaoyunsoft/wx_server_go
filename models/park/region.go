package park

import (
	"strings"
	"wx_server_go/models"
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
		data := recuRegion(rs, "0")
		return data, nil
	} else {
		return nil, err
	}
}

func recuRegion(res []Region, regionID string) []CascaderData {
	count := 0
	for _, item := range res {
		if item.PARENT_ID == regionID {
			count++
		}
	}
	if count == 0 {
		return nil
	}
	key := 0
	data := make([]CascaderData, count)
	for _, item := range res {
		if item.PARENT_ID == regionID {
			temp := new(CascaderData)
			temp.Value = item.ID
			temp.Label = item.REGION_NAME
			temp.Children = recuRegion(res, item.ID)
			data[key] = *temp
			key++
		}
	}
	return data
}

func GetRegionList(query map[string]string, page int, limit int) (total int64, res []Region, err error) {
	filterFunc := func(qs orm.QuerySeter) orm.QuerySeter {
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
		return qs
	}

	if total, err = sqltool.PageQuery_QS(new(Region), filterFunc, &res, page, limit); err == nil {
		return total, res, nil
	} else {
		return 0, nil, err
	}
}
