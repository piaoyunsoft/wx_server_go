package cus

//import (
//	"OPMS/models"
//	"strings"
//	"wx_server_go/utils/sqltool"
//
//	"github.com/astaxie/beego/orm"
//)

//import (
//	"strings"
//	"time"
//	"wx_server_go/models"
//	"wx_server_go/utils/sqltool"
//
//	"github.com/astaxie/beego/orm"
//)
//
//type PlatCusMbr struct {
//	Uid         string    `orm:"column(uid);pk;" json:"uid"`
//	CusId       string    `orm:"column(cusId)" json:"cusId"`
//	MbrId       string    `orm:"column(mbrId)" json:"mbrId"`
//	MbrName     string    `orm:"column(mbrName)" json:"mbrName"`
//	MbrType     string    `orm:"column(mbrType)" json:"mbrType"`
//	Mobile      string    `orm:"column(mobile)" json:"mobile"`
//	Idno        string    `orm:"column(idno)" json:"-"`
//	Sex         string    `orm:"column(sex)" json:"sex"`
//	BirthDate   string    `orm:"column(birthDate)" json:"birthDate"`
//	BindKey1    string    `orm:"column(bindKey1)" json:"-"`
//	BindKey2    string    `orm:"column(bindKey2)" json:"-"`
//	BindVal1    string    `orm:"column(bindVal1)" json:"-"`
//	BindVal2    string    `orm:"column(bindVal2)" json:"-"`
//	WxOpenId    string    `orm:"column(wxOpenId)" json:"-"`
//	WxBindDate  time.Time `orm:"column(wxBindDate);null;" json:"wxBindDate"`
//	BriefFromWx string    `orm:"column(briefFromWx)" json:"-"`
//}
//
////type CusMbrApi struct {
////	CusId      string    `orm:"cusId" json:"cusId"`
////	MbrId      string    `orm:"mbrId" json:"mbrId"`
////	MbrName    string    `orm:"mbrName" json:"mbrName"`
////	MbrType    string    `orm:"mbrType" json:"mbrType"`
////	Mobile     string    `orm:"mobile" json:"mobile"`
//	Sex        string    `orm:"sex" json:"sex"`
//	BirthDate  string    `orm:"birthDate" json:"birthDate"`
//	WxBindDate time.Time `orm:"wxBindDate" json:"wxBindDate"`
//}

//func (this *PlatCusMbr) TableName() string {
//	return models.TableName("platcusmbr")
//}
//
//func init() {
//	orm.RegisterModel(new(PlatCusMbr))
//}
//
//func ReadDtl_CusMbr(query map[string]string, page int, limit int) (total int64, res []PlatCusMbr, err error) {
//	filterFunc := func(qs orm.QuerySeter) orm.QuerySeter {
//		cond := orm.NewCondition()
//		// query k=v
//		for k, v := range query {
//			k = strings.Replace(k, ".", "__", -1)
//			qs = qs.Filter(k, v)
//			if k == "keyword" {
//				cond = cond.AndCond(cond.And("mbrName__icontains", v).Or("mobile__icontains", v))
//			}
//			//		if k == "WxNickName" {
//			//			cond = cond.And(k+"__icontains", v)
//			//		} else if k == "begin" {
//			//			cond = cond.And("WxSubscribeTime__gte", v)
//			//		} else if k == "end" {
//			//			cond = cond.And("WxSubscribeTime__lte", v)
//			//		} else {
//			//			cond = cond.And(k, v)
//			//		}
//		}
//		qs = qs.SetCond(cond)
//		qs = qs.OrderBy("mbrName")
//		return qs
//	}
//
//	if total, err = sqltool.PageQuery_QS(new(PlatCusMbr), filterFunc, &res, page, limit); err == nil {
//		return total, res, nil
//	} else {
//		return 0, nil, err
//	}
//}
