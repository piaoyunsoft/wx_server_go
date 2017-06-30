package models

import (
	"time"

	"strconv"

	"wx_server_go/utils/sqltool"

	"github.com/astaxie/beego/orm"
	"github.com/jdongdong/go-lib/tool"
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
	Date      string  `json:"date"`
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

	total_subscribe, today_subscribe, total_bind, today_bind, total_charge, today_charge, total_gift, today_gift, err := total_statistics(o)
	if err != nil {
		return nil, err
	} else {
		item.TotalSubscribe = total_subscribe
		item.TodaySubscribe = today_subscribe
		item.TotalBind = total_bind
		item.TodayBind = today_bind
		item.TotalCharge = total_charge
		item.TodayCharge = today_charge
		item.TotalGift = total_gift
		item.TodayGift = today_gift
	}

	//total_subscribe, err := total_subscribe_statistics(o)
	//if err != nil {
	//	return nil, err
	//} else {
	//	item.TotalSubscribe = total_subscribe
	//}
	//
	//today_subscribe, err := today_subscribe_statisitcs(o)
	//if err != nil {
	//	return nil, err
	//} else {
	//	item.TodaySubscribe = today_subscribe
	//}
	//
	//total_bind, err := total_bind_statisitcs(o)
	//if err != nil {
	//	return nil, err
	//} else {
	//	item.TotalBind = total_bind
	//}
	//
	//today_bind, err := today_bind_statisitcs(o)
	//if err != nil {
	//	return nil, err
	//} else {
	//	item.TodayBind = today_bind
	//}
	//
	//total_charge, err := total_charge_statisitcs(o)
	//if err != nil {
	//	return nil, err
	//} else {
	//	item.TotalCharge = total_charge
	//}
	//
	//today_charge, err := today_charge_statisitcs(o)
	//if err != nil {
	//	return nil, err
	//} else {
	//	item.TodayCharge = today_charge
	//}
	//
	//total_gift, err := total_gift_statistics(o)
	//if err != nil {
	//	return nil, err
	//} else {
	//	item.TotalGift = total_gift
	//}
	//
	//today_gift, err := today_gift_statistics(o)
	//if err != nil {
	//	return nil, err
	//} else {
	//	item.TodayGift = today_gift
	//}

	user_chart_rs, err := user_chart_statistics(o)
	if err != nil {
		return nil, err
	} else {
		item.UserStatistics = user_chart_rs
	}

	charge_chart_rs, err := charge_chart_statistics(o)
	if err != nil {
		return nil, err
	} else {
		item.ChargeStatistics = charge_chart_rs
	}

	return item, nil
}

//合计统计
func total_statistics(o orm.Ormer) (int64, int64, int64, int64, float64, float64, int64, int64, error) {
	timeBegin := sqltool.TimeToStr(time.Now(), "yyyy-MM-dd") + " 00:00:00"
	timeEnd := sqltool.TimeToStr(time.Now(), "yyyy-MM-dd") + " 23:59:59"

	var maps []orm.Params
	//sql:=`select * from (select count(1) as total_subscribe from wxsubscribe where (status='na' or status='aa') and subscribed='y') as aa,
	//(select count(1) as today_subscribe from wxsubscribe where (status='na' or status='aa') and subscribed='y' and wxSubscribeTime >=? and wxSubscribeTime <=?) as bb,
	//(select count(1) as total_bind from wxsubscribe where status='aa' and subscribed='y') as cc,
	//(select count(1) as today_bind from wxsubscribe where status='aa' and subscribed='y' and BindDate >=? and BindDate <=?) as dd,
	//(select sum(amt) as total_amt from wxchargeodr where status = 'aa') as ee,
	//(select sum(amt) as today_amt from wxchargeodr where status = 'aa' and payTime >=? and payTime <=?) as ff,
	//(select count(1) as total_gift from vipgiftexch where status='aa') as gg,
	//(select count(1) as today_gift from vipgiftexch where status='aa' and createDate >=? and createDate <=?) as hh`

	_, err := o.Raw("select * from (select count(1) as total_subscribe from wxsubscribe where (status='na' or status='aa') and subscribed='y') as aa,(select count(1) as today_subscribe from wxsubscribe where (status='na' or status='aa') and subscribed='y' and wxSubscribeTime >=? and wxSubscribeTime <=?) as bb,(select count(1) as total_bind from wxsubscribe where status='aa' and subscribed='y') as cc,(select count(1) as today_bind from wxsubscribe where status='aa' and subscribed='y' and BindDate >=? and BindDate <=?) as dd,(select sum(amt) as total_amt from wxchargeodr where status = 'aa') as ee,(select sum(amt) as today_amt from wxchargeodr where status = 'aa' and payTime >=? and payTime <=?) as ff,(select count(1) as total_gift from vipgiftexch where status='aa') as gg,(select count(1) as today_gift from vipgiftexch where status='aa' and createDate >=? and createDate <=?) as hh", timeBegin, timeEnd, timeBegin, timeEnd, timeBegin, timeEnd, timeBegin, timeEnd).Values(&maps)
	if err != nil {
		return 0, 0, 0, 0, 0, 0, 0, 0, err
	}
	return strToInt64(maps[0]["total_subscribe"]), strToInt64(maps[0]["today_subscribe"]), strToInt64(maps[0]["total_bind"]), strToInt64(maps[0]["today_bind"]), strToFloat(maps[0]["total_amt"]), strToFloat(maps[0]["today_amt"]), strToInt64(maps[0]["total_gift"]), strToInt64(maps[0]["today_gift"]), nil
}

//用户7日趋势图
func user_chart_statistics(o orm.Ormer) ([]UserCharts, error) {
	timeBegin := sqltool.TimeToStr(time.Now().Add(time.Hour*24*6*-1), "yyyy-MM-dd") + " 00:00:00"
	timeEnd := sqltool.TimeToStr(time.Now(), "yyyy-MM-dd") + " 23:59:59"

	var newSubscribeMaps, newBindMaps, cancelSubscribeMaps []orm.Params
	_, err := o.Raw("SELECT DATE(wxSubscribeTime) AS time, count(1) AS sum FROM wxsubscribe WHERE (STATUS = 'na' OR STATUS = 'aa') AND subscribed = 'y' and wxSubscribeTime>=? and wxSubscribeTime <=? GROUP BY DATE(wxSubscribeTime) ORDER BY time DESC LIMIT 0, 7;", timeBegin, timeEnd).Values(&newSubscribeMaps)
	if err != nil {
		return nil, err
	}

	_, err = o.Raw("SELECT DATE(wxUnsubscribeTime) AS time, count(1) AS sum FROM wxsubscribe WHERE (STATUS = 'na' OR STATUS = 'aa') AND subscribed = 'n' and wxUnsubscribeTime>=? and wxUnsubscribeTime <=? GROUP BY DATE(wxUnsubscribeTime) ORDER BY time DESC LIMIT 0, 7;", timeBegin, timeEnd).Values(&cancelSubscribeMaps)
	if err != nil {
		return nil, err
	}

	_, err = o.Raw("SELECT DATE(BindDate) AS time, count(1) AS sum FROM wxsubscribe WHERE STATUS = 'aa' and subscribed = 'y' and BindDate>=? and BindDate <=?  GROUP BY DATE(BindDate) ORDER BY time DESC LIMIT 0, 7;", timeBegin, timeEnd).Values(&newBindMaps)
	if err != nil {
		return nil, err
	}

	items := make([]UserCharts, 7)
	for i := 6; i >= 0; i-- {
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
		items[6-i] = *item
	}
	return items, nil
}

//7日充值趋势图
func charge_chart_statistics(o orm.Ormer) ([]ChargeCharts, error) {
	timeBegin := sqltool.TimeToStr(time.Now().Add(time.Hour*24*6*-1), "yyyy-MM-dd") + " 00:00:00"
	timeEnd := sqltool.TimeToStr(time.Now(), "yyyy-MM-dd") + " 23:59:59"

	var maps []orm.Params
	_, err := o.Raw("SELECT DATE(payTime) as time, sum(amt) as sum, (SELECT sum(amt) FROM wxchargeodr a WHERE STATUS = 'aa' AND payPtf = 'Alipay' AND DATE(a.payTime) = DATE(wxchargeodr.payTime)) as alipayamt,	(SELECT sum(amt) FROM wxchargeodr a WHERE STATUS = 'aa' AND payPtf = 'WechatPay' AND DATE(a.payTime) = DATE(wxchargeodr.payTime)) as wechatpayamt	FROM wxchargeodr WHERE STATUS = 'aa' and NOT ISNULL(payTime) and payTime>=? and payTime <=?  GROUP BY DATE(payTime) ORDER BY time DESC LIMIT 0, 7;", timeBegin, timeEnd).Values(&maps)
	if err != nil {
		return nil, err
	}
	items := make([]ChargeCharts, 7)
	for i := 6; i >= 0; i-- {
		item := new(ChargeCharts)
		item.Date = sqltool.TimeToStr(time.Now().Add(time.Hour*24*time.Duration(i)*-1), "yyyy-MM-dd")
		for _, v := range maps {
			if tool.ToString(v["time"]) == item.Date {
				item.TotalAmt = strToFloat(v["sum"])
				item.AlipayAmt = strToFloat(v["alipayamt"])
				item.WechatAmt = strToFloat(v["wechatpayamt"])
			}
		}
		items[6-i] = *item
	}
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
