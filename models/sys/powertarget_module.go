package sys

//import (
//	"strings"
//	"wx_server_go/models"
//	"wx_server_go/utils/sqltool"
//
//	"github.com/astaxie/beego/orm"
//)
//
//type PowerModule struct {
//	ID       int    `orm:"column(ID);pk" json:"id"`
//	PID      int    `orm:"column(PID)" json:"pid"`
//	NAME     string `orm:"column(NAME)" json:"name"`
//	SORT     int    `orm:"column(SORT)" json:"sort"`
//	TYPE     string `orm:"column(TYPE)" json:"type"`
//	TARGET   string `orm:"column(TARGET)" json:"url"`
//	REMARK   string `orm:"column(REMARK)" json:"remark"`
//	STATUS   string `orm:"column(STATUS)" json:"status"`
//	ACTIONID int    `orm:"column(ACTIONID)" json:"-"`
//	ICON     string `orm:"column(ICON)" json:"icon"`
//}
//
//const tablename = "powertarget_module"
//
//func (this *PowerModule) TableName() string {
//	return models.TableName(tablename)
//}
//
//func init() {
//	orm.RegisterModel(new(PowerModule))
//}
//
//func GetPowerModules(query map[string]string) (res []PowerModule, err error) {
//	filterFunc := func(qs orm.QuerySeter) orm.QuerySeter {
//		for k, v := range query {
//			k = strings.Replace(k, ".", "__", -1)
//			qs = qs.Filter(k, v)
//		}
//		return qs
//	}
//
//	if err := sqltool.Query_QS(new(PowerModule), filterFunc, &res); err == nil {
//		return res, nil
//	} else {
//		return nil, err
//	}
//}
//
//type IdStruct struct {
//	Newid int `orm:"column(newid)"`
//}
//
//func CreatePowerModule(item *PowerModule) error {
//	if id, err := sqltool.NewId(tablename); err == nil {
//		item.ID = id
//		return sqltool.Create(item)
//	} else {
//		return err
//	}
//}
//
//func UpdatePowerModule(item *PowerModule) error {
//	return sqltool.Update(item)
//}
//
//func DelPowerModule(id int) error {
//	filter := PowerModule{ID: id}
//	return sqltool.Delete(&filter)
//}
