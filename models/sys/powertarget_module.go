package sys

import (
	"wx_server_go/models"
	"wx_server_go/utils/sqltool"

	"github.com/astaxie/beego/orm"
)

type PowerModule struct {
	ID       int    `orm:"column(ID);pk" json:"id"`
	PID      int    `orm:"column(PID)" json:"pid"`
	NAME     string `orm:"column(NAME)" json:"name"`
	SORT     int    `orm:"column(SORT)" json:"sort"`
	TYPE     string `orm:"column(TYPE)" json:"type"`
	TARGET   string `orm:"column(TARGET)" json:"url"`
	REMARK   string `orm:"column(REMARK)" json:"remark"`
	STATUS   string `orm:"column(STATUS)" json:"status"`
	ACTIONID int    `orm:"column(ACTIONID)" json:"-"`
	ICON     string `orm:"column(ICON)" json:"icon"`
}

func (this *PowerModule) TableName() string {
	return models.TableName("powertarget_module")
}

func init() {
	orm.RegisterModel(new(PowerModule))
}

func GetPowerModules() (res []PowerModule, err error) {
	filterFunc := func(qs *orm.QuerySeter) {

	}

	if err := sqltool.Query_QS(new(PowerModule), filterFunc, &res); err == nil {
		return res, nil
	} else {
		return nil, err
	}
}
