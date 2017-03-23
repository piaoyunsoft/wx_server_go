package cus

//import (
//	"strings"
//	"time"
//	"wx_server_go/models"
//	"wx_server_go/utils/sqltool"
//
//	"github.com/astaxie/beego/orm"
//)
//
//type VipGiftExch struct {
//	ExchID        string    `orm:"column(exchID);pk" json:"exchID"`
//	ComId         string    `orm:"column(comId)" json:"comId"`
//	MbrID         string    `orm:"column(mbrID)" json:"mbrID"`
//	WxOpenID      string    `orm:"column(wxOpenID)" json:"wxOpenID"`
//	GiftCode      string    `orm:"column(giftCode)" json:"giftCode"`
//	GiftName      string    `orm:"column(giftName)" json:"giftName"`
//	GetWay        string    `orm:"column(getWay)" json:"getWay"`
//	MailAddr      string    `orm:"column(mailAddr)" json:"mailAddr"`
//	MailPsnName   string    `orm:"column(mailPsnName)" json:"mailPsnName"`
//	MailPsnMobile string    `orm:"column(mailPsnMobile)" json:"mailPsnMobile"`
//	ExchQty       string    `orm:"column(exchQty)" json:"exchQty"`
//	UseScore      int       `orm:"column(useScore)" json:"useScore"`
//	Status        string    `orm:"column(status)" json:"status"`
//	MailStatus    string    `orm:"column(mailStatus)" json:"mailStatus"`
//	CreateDate    time.Time `orm:"column(createDate)" json:"createDate"`
//	ChangeDate    time.Time `orm:"column(changeDate)" json:"changeDate"`
//}
//
//type VipGiftExchDtl struct {
//	ExchID        string    `orm:"column(exchID);pk" json:"exchID"`
//	ComId         string    `orm:"column(comId)" json:"comId"`
//	MbrID         string    `orm:"column(mbrID)" json:"mbrID"`
//	WxOpenID      string    `orm:"column(wxOpenID)" json:"wxOpenID"`
//	GiftCode      string    `orm:"column(giftCode)" json:"giftCode"`
//	GiftName      string    `orm:"column(giftName)" json:"giftName"`
//	GetWay        string    `orm:"column(getWay)" json:"getWay"`
//	MailAddr      string    `orm:"column(mailAddr)" json:"mailAddr"`
//	MailPsnName   string    `orm:"column(mailPsnName)" json:"mailPsnName"`
//	MailPsnMobile string    `orm:"column(mailPsnMobile)" json:"mailPsnMobile"`
//	ExchQty       string    `orm:"column(exchQty)" json:"exchQty"`
//	UseScore      int       `orm:"column(useScore)" json:"useScore"`
//	Status        string    `orm:"column(status)" json:"status"`
//	MailStatus    string    `orm:"column(mailStatus)" json:"mailStatus"`
//	CreateDate    time.Time `orm:"column(createDate)" json:"createDate"`
//	ChangeDate    time.Time `orm:"column(changeDate)" json:"changeDate"`
//	//扩展
//	NickName       string `orm:"column(nickName)" json:"nickName"`
//	GetWayName     string `orm:"column(getWayName)" json:"getWayName"`
//	MailStatusName string `orm:"column(mailStatusName)" json:"mailStatusName"`
//}
//
//func (this *VipGiftExch) TableName() string {
//	return models.TableName("vipgiftexch")
//}
//
//func init() {
//	orm.RegisterModel(new(VipGiftExch))
//}
//
//func GetPageVipGiftExch(query map[string]string, page int, limit int) (total int64, res []VipGiftExchDtl, err error) {
//	qb := sqltool.GetQueryBuilder()
//	_w := "1=1"
//	for k, v := range query {
//		k = strings.Replace(k, ".", "__", -1)
//		if k == "giftName" {
//			_w += " and a.giftName like '%" + sqltool.SqlFormat(v) + "%'"
//		} else if k == "begin" {
//			_w += " and a.createDate >= '" + sqltool.SqlFormat(v) + "' "
//		} else if k == "end" {
//			_w += " and a.createDate <= '" + sqltool.SqlFormat(v) + "' "
//		} else if k == "mbr" {
//			_w += " and b.wxNickName like '%" + sqltool.SqlFormat(v) + "%'"
//		}
//	}
//	qb.Select("a.*", "b.wxNickName as nickName", "c.itemname as getWayName", "d.itemname as mailStatusName").From("vipgiftexch as a").
//		LeftJoin("wxsubscribe as b").On("(b.wxOpenId = a.wxOpenID and b.comID = a.comId)").
//		LeftJoin("dictitem as c").On("(c.itemcode = a.getWay and c.dictcode = '002')").
//		LeftJoin("dictitem as d").On("(d.itemcode = a.mailStatus and d.dictcode = '003')").
//		Where(_w).
//		OrderBy("a.createDate").
//		Desc()
//
//	if total, err = sqltool.PageQuery_QB(qb, &res, page, limit); err == nil {
//		return total, res, nil
//	} else {
//		return 0, nil, err
//	}
//}
