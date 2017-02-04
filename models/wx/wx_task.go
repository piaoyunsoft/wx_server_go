package wx

import (
	"strings"
	"time"
	"wx_server_go/utils"
	"wx_server_go/utils/sqltool"

	"github.com/astaxie/beego/orm"
)

type Platcuswxtask struct {
	Id           int64     `orm:"column(id);size(20);pk;" json:"id"`
	FromCusID    string    `orm:"column(FromCusID);" json:"FromCusID"`
	ToWxOpenId   string    `orm:"column(ToWxOpenId);" json:"ToWxOpenId"`
	ToCusMbrId   string    `orm:"column(ToCusMbrId);" json:"ToCusMbrId"`
	ToCusMbrName string    `orm:"column(ToCusMbrName);" json:"ToCusMbrName"`
	ToAccountId  string    `orm:"column(ToAccountId);" json:"ToAccountId"`
	TplCode      string    `orm:"column(tplCode);" json:"tplCode"`
	TplId        string    `orm:"column(tplId);" json:"tplId"`
	Type         string    `orm:"column(type);" json:"type"`
	MsgBody      string    `orm:"column(msgBody);" json:"msgBody"`
	CreateDtm    time.Time `orm:"column(createDtm);" json:"createDtm"`
	Status       string    `orm:"column(status);" json:"status"`
	SendDtm      string    `orm:"column(sendDtm);" json:"sendDtm"`
	WxMsgId      string    `orm:"column(wxMsgId);" json:"wxMsgId"`
	ReceiveDtm   string    `orm:"column(receiveDtm);" json:"receiveDtm"`
	Remark       string    `orm:"column(remark);" json:"remark"`
	Url          string    `orm:"column(url);" json:"url"`
}

type WxTaskApi struct {
	Id        int64  `orm:"column(id)" json:"id"`
	NickName  string `orm:"column(wxNickName);" json:"nickName"`
	MsgBody   string `orm:"column(msgBody)" json:"msgBody"`
	MbrName   string `orm:"column(ToCusMbrName)" json:"mbrName"`
	Status    string `orm:"column(status)" json:"status"`
	Remark    string `orm:"column(remark)" json:"remark"`
	SendDtm   string `orm:"column(sendDtm)" json:"sendDtm"`
	CreateDtm string `orm:"column(createDtm)" json:"createDtm"`
	Type      string `orm:"column(type)" json:"type"`
	TplName   string `orm:"column(title)" json:"tplName"`
	CusName   string `orm:"column(cusName)" json:"cusName"`
}

func init() {
	orm.RegisterModel(new(Platcuswxtask))
}

func ReadDtl_WxTask(query map[string]string, page int, limit int) (total int64, res []WxTaskApi, err error) {
	var tasks []WxTaskApi
	qb := sqltool.GetQueryBuilder()

	_w := "1=1"
	for k, v := range query {
		k = strings.Replace(k, ".", "__", -1)
		if k == "nickName" {
			_w += " and (b.wxNickName like '%" + sqltool.SqlFormat(v) + "%' or a.ToCusMbrName like '%" + sqltool.SqlFormat(v) + "%' )"
		} else if k == "begin" {
			_w += " and sendDtm >= '" + sqltool.SqlFormat(v) + "' "
		} else if k == "end" {
			_w += " and sendDtm <= '" + sqltool.SqlFormat(v) + "' "
		} else if k == "status" {
			_w += " and a.status = '" + sqltool.SqlFormat(v) + "' "
		}
	}

	qb.Select("a.*", "b.wxNickName", "c.title", "d.cusName").From("platcuswxtask as a").
		LeftJoin("wxsubscribe as b").On("b.wxOpenId = a.ToWxOpenId").
		LeftJoin("wxtplmsgtpllib as c").On("c.tplCode = a.tplCode").
		LeftJoin("platcus as d").On("d.cusID = a.FromCusID").
		Where(_w).
		OrderBy("a.createDtm").
		Desc()

	if total, err = sqltool.PageQuery_QB(qb, &tasks, page, limit); err == nil {
		return total, tasks, nil
	} else {
		return 0, nil, err
	}
}

func ReadWxTaskList(query map[string]string, page int64, limit int64) (total int64, res []Platcuswxtask, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Platcuswxtask))
	cond := orm.NewCondition()
	// query k=v
	for k, v := range query {
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
		//		if k == "WxNickName" {
		//			cond = cond.And(k+"__icontains", v)
		//		} else if k == "begin" {
		//			cond = cond.And("WxSubscribeTime__gte", v)
		//		} else if k == "end" {
		//			cond = cond.And("WxSubscribeTime__lte", v)
		//		} else {
		//			cond = cond.And(k, v)
		//		}
	}
	qs = qs.SetCond(cond)
	qs = qs.OrderBy("-createDtm")

	if total, err = qs.Count(); err == nil {
		offset := (page - 1) * limit
		if _, err := qs.Limit(limit, offset).All(&res); err == nil {
			return total, res, nil
		}
	}
	utils.Error(err)
	return 0, nil, err
}
