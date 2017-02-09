package charge

import (
	"strings"
	"time"
	"wx_server_go/models"
	"wx_server_go/utils/sqltool"

	"github.com/astaxie/beego/orm"
)

type WxChargeOrd struct {
	OrdID       string    `orm:"column(odrID);pk" json:"odrID"`
	PayPtf      string    `orm:"column(payPtf)" json:"payPtf"`
	Subuid      string    `orm:"column(Subuid)" json:"Subuid"`
	ComId       string    `orm:"column(comId)" json:"comId"`
	WxOpenId    string    `orm:"column(wxOpenId)" json:"wxOpenId"`
	MbrId       string    `orm:"column(mbrId)" json:"mbrId"`
	Amt         float32   `orm:"column(amt)" json:"amt"`
	PayPtfOdrID int64     `orm:"column(payPtfOdrID)" json:"payPtfOdrID"`
	PayTime     time.Time `orm:"column(payTime)" json:"payTime"`
	Status      string    `orm:"column(status)" json:"status"`
	ErrMsg      string    `orm:"column(errMsg)" json:"errMsg"`
	CreateDate  time.Time `orm:"column(createDate)" json:"createDate"`
	ChangeDate  time.Time `orm:"column(changeDate)" json:"changeDate"`
	NickName    string    `orm:"column(nickName)" json:"nickName"`
}

func (this *WxChargeOrd) TableName() string {
	return models.TableName("wxchargeodr")
}

func init() {
	orm.RegisterModel(new(WxChargeOrd))
}

func GetPageChargeOrds(query map[string]string, page int, limit int) (total int64, res []WxChargeOrd, err error) {
	qb := sqltool.GetQueryBuilder()

	_w := "1=1"
	for k, v := range query {
		k = strings.Replace(k, ".", "__", -1)
		if k == "keyword" {
			_w += " and (b.wxNickName like '%" + sqltool.SqlFormat(v) + "%' )"
		} else if k == "begin" {
			_w += " and payTime >= '" + sqltool.SqlFormat(v) + "' "
		} else if k == "end" {
			_w += " and payTime <= '" + sqltool.SqlFormat(v) + "' "
		}
	}

	qb.Select("a.*", "b.wxNickName as nickName").From("wxchargeodr as a").
		LeftJoin("wxsubscribe as b").On("(b.wxOpenId = a.wxOpenID and b.comID = a.comId)").
		Where(_w).
		OrderBy("a.odrID").
		Desc()

	if total, err = sqltool.PageQuery_QB(qb, &res, page, limit); err == nil {
		return total, res, nil
	} else {
		return 0, nil, err
	}
	//	o := orm.NewOrm()
	//	qs := o.QueryTable(new(WxChargeOrd))
	//	cond := orm.NewCondition()
	//	// query k=v
	//	for k, v := range query {
	//		k = strings.Replace(k, ".", "__", -1)
	//		qs = qs.Filter(k, v)
	//		if k == "keyword" {
	//			cond = cond.AndCond(cond.And("mbrName__icontains", v).Or("mobile__icontains", v))
	//		}
	//	}
	//	qs = qs.SetCond(cond)
	//	qs = qs.OrderBy("odrID")

	//	if total, err = sqltool.PageQuery_QS(qs, &res, page, limit); err == nil {
	//		return total, res, nil
	//	} else {
	//		return 0, nil, err
	//	}
}
