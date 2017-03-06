package models

import (
	"time"

	"strconv"

	"wx_server_go/utils/sqltool"

	"fmt"

	"github.com/astaxie/beego/orm"
	"github.com/ddliao/go-lib/tool"
)

type Statistics struct {
	TotalSubscribe   int64          `json:"total_subscribe"`
	TodaySubscribe   int64          `json:"today_subscribe"`
	TotalBind        int64          `json:"total_bind"`
	TodayBind        int64          `json:"today_bind"`
	TotalCharge      float64        `json:"total_charge"`
	TodayCharge      float64        `json:"today_charge"`
	TotalGift        int64          `json:"total_gift"`
	TodayGift        int64          `json:"today_gift"`
	UserStatistics   []UserCharts   `json:"user_statistics"`
	ChargeStatistics []ChargeCharts `json:"charge_statistics"`
}

type UserCharts struct {
	Date            string `json:"date"`
	NewSubscribe    int64  `json:"new_subscribe"`
	NewBind         int64  `json:"new_bind"`
	CancelSubscribe int64  `json:"cancel_subscribe"`
}

type ChargeCharts struct {
	TotalAmt  float64 `json:"total_amt"`
	AlipayAmt float64 `json:"alipay_amt"`
	WechatAmt float64 `json:"wechat_amt"`
}

type UserChartOption struct {
	Time time.Time
	Sum  int64
}

func GetStatistics() (*Statistics, error) {
	item := new(Statistics)
	o := orm.NewOrm()

	total_subscribe, err := total_subscribe_statistics(o)
	if err != nil {
		return nil, err
	} else {
		item.TotalSubscribe = total_subscribe
	}

	today_subscribe, err := today_subscribe_statisitcs(o)
	if err != nil {
		return nil, err
	} else {
		item.TodaySubscribe = today_subscribe
	}

	total_bind, err := total_bind_statisitcs(o)
	if err != nil {
		return nil, err
	} else {
		item.TotalBind = total_bind
	}

	today_bind, err := today_bind_statisitcs(o)
	if err != nil {
		return nil, err
	} else {
		item.TodayBind = today_bind
	}

	total_charge, err := total_charge_statisitcs(o)
	if err != nil {
		return nil, err
	} else {
		item.TotalCharge = total_charge
	}

	today_charge, err := today_charge_statisitcs(o)
	if err != nil {
		return nil, err
	} else {
		item.TodayCharge = today_charge
	}

	total_gift, err := total_gift_statistics(o)
	if err != nil {
		return nil, err
	} else {
		item.TotalGift = total_gift
	}

	today_gift, err := today_gift_statistics(o)
	if err != nil {
		return nil, err
	} else {
		item.TodayGift = today_gift
	}

	user_chart_rs, err := user_chart_statistics(o)
	if err != nil {
		return nil, err
	} else {
		item.UserStatistics = user_chart_rs
	}

	return item, nil
}

//总关注人数
func total_subscribe_statistics(o orm.Ormer) (int64, error) {
	cond := orm.NewCondition()
	qs := o.QueryTable("wxsubscribe")
	cond = cond.AndCond(cond.And("status", "na").Or("status", "aa")).And("subscribed", "y")
	qs = qs.SetCond(cond)
	count, err := qs.Count()
	if err != nil {
		return 0, err
	} else {
		return count, nil
	}
}

//今日关注人数
func today_subscribe_statisitcs(o orm.Ormer) (int64, error) {
	cond := orm.NewCondition()
	qs := o.QueryTable("wxsubscribe")
	cond = cond.AndCond(cond.And("status", "na").Or("status", "aa")).And("subscribed", "y").And("wxSubscribeTime__gte", time.Now().Format("2006-01-02")).And("wxSubscribeTime__lte", time.Now().Format("2006-01-02")+" 23:59:59")
	qs = qs.SetCond(cond)
	count, err := qs.Count()
	if err != nil {
		return 0, err
	} else {
		return count, err
	}
}

//总绑定人数
func total_bind_statisitcs(o orm.Ormer) (int64, error) {
	cond := orm.NewCondition()
	qs := o.QueryTable("wxsubscribe")
	cond = cond.And("status", "aa").And("subscribed", "y")
	qs = qs.SetCond(cond)
	count, err := qs.Count()
	if err != nil {
		return 0, err
	} else {
		return count, err
	}
}

//今日绑定人数
func today_bind_statisitcs(o orm.Ormer) (int64, error) {
	cond := orm.NewCondition()
	qs := o.QueryTable("wxsubscribe")
	cond = cond.And("status", "aa").And("subscribed", "y").And("wxSubscribeTime__gte", time.Now().Format("2006-01-02")).And("wxSubscribeTime__lte", time.Now().Format("2006-01-02")+" 23:59:59")
	qs = qs.SetCond(cond)
	count, err := qs.Count()
	if err != nil {
		return 0, err
	} else {
		return count, err
	}
}

//总充值金额
func total_charge_statisitcs(o orm.Ormer) (float64, error) {
	var maps []orm.Params
	_, err := o.Raw("select sum(amt) as amt from wxchargeodr where status = 'yy'").Values(&maps)
	if err != nil {
		return 0, err
	} else {
		return strToFloat(maps[0]["amt"]), nil
	}
}

//今日充值金额
func today_charge_statisitcs(o orm.Ormer) (float64, error) {
	var maps []orm.Params
	_, err := o.Raw("select sum(amt) as amt from wxchargeodr where status = 'yy' and createDate >=? and createDate <= ?", time.Now().Format("2006-01-02"), time.Now().Format("2006-01-02")+" 23:59:59").Values(&maps)
	if err != nil {
		return 0, err
	} else {
		return strToFloat(maps[0]["amt"]), nil
	}
}

//总兑换笔数
func total_gift_statistics(o orm.Ormer) (int64, error) {
	cond := orm.NewCondition()
	qs := o.QueryTable("vipgiftexch")
	cond = cond.And("status", "aa")
	qs = qs.SetCond(cond)
	count, err := qs.Count()
	if err != nil {
		return 0, err
	} else {
		return count, nil
	}
}

//今日兑换笔数
func today_gift_statistics(o orm.Ormer) (int64, error) {
	cond := orm.NewCondition()
	qs := o.QueryTable("vipgiftexch")
	cond = cond.And("status", "aa").And("createDate__gte", time.Now().Format("2006-01-02")).And("createDate__lte", time.Now().Format("2006-01-02")+" 23:59:59")
	qs = qs.SetCond(cond)
	count, err := qs.Count()
	if err != nil {
		return 0, err
	} else {
		return count, err
	}
}

//用户7日趋势图
func user_chart_statistics(o orm.Ormer) ([]UserCharts, error) {
	timeBegin := sqltool.TimeToStr(time.Now().Add(time.Hour*24*6*-1), "yyyy-MM-dd") + "00:00:00"
	timeEnd := sqltool.TimeToStr(time.Now(), "yyyy-MM-dd") + "23:59:59"

	var newSubscribeMaps, newBindMaps, cancelSubscribeMaps []orm.Params
	_, err := o.Raw("SELECT DATE(wxSubscribeTime) AS time, count(1) AS sum FROM wxsubscribe WHERE (STATUS = 'na' OR STATUS = 'aa') AND subscribed = 'y' and wxSubscribeTime>=? and wxSubscribeTime <=? GROUP BY DATE(wxSubscribeTime) ORDER BY time DESC LIMIT 0, 7;", timeBegin, timeEnd).Values(&newSubscribeMaps)
	if err != nil {
		return nil, err
	}

	_, err = o.Raw("SELECT DATE(wxUnsubscribeTime) AS time, count(1) AS sum FROM wxsubscribe WHERE (STATUS = 'na' OR STATUS = 'aa') AND subscribed = 'n' and wxUnsubscribeTime>=? and wxUnsubscribeTime <=? GROUP BY DATE(wxUnsubscribeTime) ORDER BY time DESC LIMIT 0, 7;", timeBegin, timeEnd).Values(&newBindMaps)
	if err != nil {
		return nil, err
	}

	_, err = o.Raw("SELECT DATE(BindDate) AS time, count(1) AS sum FROM wxsubscribe WHERE STATUS = 'aa' and subscribed = 'y' and BindDate>=? and BindDate <=?  GROUP BY DATE(BindDate) ORDER BY time DESC LIMIT 0, 7;", timeBegin, timeEnd).Values(&cancelSubscribeMaps)
	if err != nil {
		return nil, err
	}

	items := make([]UserCharts, 7)
	for i := 0; i < 7; i++ {
		item := new(UserCharts)
		item.Date = sqltool.TimeToStr(time.Now().Add(time.Hour*24*time.Duration(i)*-1), "yyyy-MM-dd")
		for _, v := range newSubscribeMaps {
			if tool.ToString(v["time"]) == item.Date {
				item.NewSubscribe = strToInt64(v["sum"])
			}
		}
		for _, v := range newBindMaps {
			if tool.ToString(v["time"]) == item.Date {
				item.NewBind = strToInt64(v["sum"])
			}
		}
		for _, v := range cancelSubscribeMaps {
			if tool.ToString(v["time"]) == item.Date {
				item.CancelSubscribe = strToInt64(v["sum"])
			}
		}
		items[i] = *item
	}
	fmt.Println(fmt.Sprintf("%+v", items))
	return items, nil
}

func strToFloat(s interface{}) float64 {
	switch v := s.(type) {
	case string:
		str := v
		f, err := strconv.ParseFloat(str, 64)
		if err == nil {
			return f
		}
		break
	}
	return 0
}

func strToInt64(s interface{}) int64 {
	switch v := s.(type) {
	case string:
		str := v
		f, err := strconv.ParseInt(str, 10, 64)
		if err == nil {
			return f
		}
		break
	}
	return 0
}

//select DATE(payTime),payPtf, sum(amt)  from wxchargeodr where status='aa' GROUP BY DATE(payTime), payPtf
