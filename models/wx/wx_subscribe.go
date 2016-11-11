package wx

import (
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type Wxsubscribe struct {
	Uid               string    `orm:"column(uid);size(48);pk;" json:"uid"`
	WxOpenId          string    `orm:"column(wxOpenId);size(40);" json:"wxOpenId"`
	WxUnionId         string    `orm:"column(wxUnionId);size(48);null;" json:"wxUnionId"`
	ComWxID           string    `orm:"column(comWxID);size(100);" json:"comWxID"`
	ComID             string    `orm:"column(comID);size(60);null;" json:"comID"`
	WxNickName        string    `orm:"column(wxNickName);size(100);null;" json:"wxNickName"`
	WxSex             string    `orm:"column(wxSex);size(1);null;" json:"wxSex"`
	Subscribed        string    `orm:"column(subscribed);size(1);null;" json:"subscribed"`
	WxSubscribeTime   time.Time `orm:"column(wxSubscribeTime);" json:"wxSubscribeTime"`
	WxUnsubscribeTime time.Time `orm:"column(wxUnsubscribeTime);null;" json:"wxUnsubscribeTime"`
	WxCountry         string    `orm:"column(wxCountry);size(40);null;" json:"wxCountry"`
	WxProvince        string    `orm:"column(wxProvince);size(40);null;" json:"wxProvince"`
	WxCity            string    `orm:"column(wxCity);size(40);null;" json:"wxCity"`
	WxHeadImgUrl      string    `orm:"column(wxHeadImgUrl);size(255);null;" json:"wxHeadImgUrl"`
	WxSubscribeCount  int       `orm:"column(wxSubscribeCount);size(11);null;" json:"wxSubscribeCount"`
	WxBrief           string    `orm:"column(wxBrief);size(600);null;" json:"wxBrief"`
	BindDate          time.Time `orm:"column(BindDate);null;" json:"BindDate"`
	BindWay           string    `orm:"column(BindWay);size(30);null;" json:"BindWay"`
	MbrId             string    `orm:"column(mbrId);size(32);null;" json:"mbrId"`
	AduitDate         time.Time `orm:"column(aduitDate);null;" json:"aduitDate"`
	AduitPerson       string    `orm:"column(aduitPerson);size(30);null;" json:"aduitPerson"`
	Status            string    `orm:"column(status);size(2);" json:"status"`
	MbrName           string    `orm:"column(mbrName);size(64);null;" json:"mbrName"`
	MbrType           string    `orm:"column(mbrType);size(10);null;" json:"mbrType"`
	Mobile            string    `orm:"column(mobile);size(30);null;" json:"mobile"`
	Idno              string    `orm:"column(idno);size(20);null;" json:"idno"`
	BirthDate         string    `orm:"column(birthDate);size(8);null;" json:"birthDate"`
	Addr              string    `orm:"column(addr);size(255);null;" json:"addr"`
	CreateDate        time.Time `orm:"column(createDate);" json:"createDate"`
	ChangeDate        time.Time `orm:"column(changeDate);null;" json:"changeDate"`
}

func init() {
	orm.RegisterModel(new(Wxsubscribe))
}

func ReadWxSubscribeList(query map[string]string, page int64, limit int64) (total int64, res []Wxsubscribe, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Wxsubscribe))
	cond := orm.NewCondition()
	// query k=v
	for k, v := range query {
		k = strings.Replace(k, ".", "__", -1)
		//		qs = qs.Filter(k, v)
		if k == "WxNickName" {
			cond = cond.And(k+"__icontains", v)
		} else if k == "begin" {
			cond = cond.And("WxSubscribeTime__gte", v)
		} else if k == "end" {
			cond = cond.And("WxSubscribeTime__lte", v)
		}
	}
	qs = qs.SetCond(cond)
	qs = qs.OrderBy("-createDate")
	if total, err = qs.Count(); err == nil {
		offset := (page - 1) * limit
		if _, err := qs.Limit(limit, offset).All(&res); err == nil {
			return total, res, nil
		}
	}

	return 0, nil, err
}
