package models

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/astaxie/beego"
	"github.com/jdongdong/go-lib/slog"
	"github.com/pquerna/ffjson/ffjson"
)

type SeaCoupon struct {
	ComID string
	Title string `json:"title"`
}

type Coupon struct {
	ID              string
	ComID           string  `json:"comID"`
	CouponType      string  `json:"couponType"`
	GetWay          string  `json:"getWay"`
	LogoUrl         string  `json:"logo_url"`
	Title           string  `json:"title"`
	Content         string  `json:"content"`
	Notice          string  `json:"notice"`
	UseScope        string  `json:"useScope"`
	DefCouponValue  float32 `json:"defCouponValue"`
	Qty             int     `json:"qty"`
	TimeLimitWay    string  `json:"timeLimitWay"`
	TimeLimitValue  string  `json:"timeLimitValue"`
	Status          string  `json:"status"`
	AlreadyGetCount int     `json:"AlreadyGetCount"`
	AlreadyUseCount int     `json:"AlreadyUseCount"`
	AlreadyGetAmt   float64 `json:"AlreadyGetAmt"`
}

type CouponModal struct {
	Data  []Coupon `json:"data"`
	Total int64    `json:"total"`
}

type SeaCouponItem struct {
	CouponID string `json:"couponID"`
}

type CouponItem struct {
	ItemID     string  `json:"itemID"`
	CouponID   string  `json:"couponID"`
	GetValue   float32 `json:"getValue"`
	GetDate    string  `json:"getDate"`
	GetVipcdid string  `json:"getVipcdid"`
	FromTb     string  `json:"fromTb"`
	FromPK     string  `json:"fromPK"`
	Status     string  `json:"status"`
	UsedDate   string  `json:"usedDate"`
	UsedLinkTb string  `json:"usedLinkTb"`
	UsedLinkPK string  `json:"usedLinkPK"`
	MbrName    string  `json:"mbrName"`
}

func (this *SeaCoupon) GetCoupon() ([]Coupon, error) {
	serverAddress := beego.AppConfig.String("serveraddr")
	url := serverAddress + "wxopenapi/GetCoupons?title=" + this.Title + "&comID=" + this.ComID
	resp, err := http.Get(url)
	if err != nil {
		slog.Error(err)
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		slog.Error(err)
		return nil, err
	}

	rsModel := make([]Coupon, 0)
	err = ffjson.Unmarshal(body, &rsModel)
	if err != nil {
		slog.Error(err)
		return nil, err
	}
	slog.Info(fmt.Sprintf("url:%s rs:%+v", url, rsModel))
	return rsModel, nil
}

func (this *SeaCouponItem) GetCouponItem() ([]CouponItem, error) {
	serverAddress := beego.AppConfig.String("serveraddr")
	url := serverAddress + "wxopenapi/GetCouponItems?couponID=" + this.CouponID

	resp, err := http.Get(url)
	if err != nil {
		slog.Error(err)
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		slog.Error(err)
		return nil, err
	}

	rsModel := make([]CouponItem, 0)
	err = ffjson.Unmarshal(body, &rsModel)
	if err != nil {
		slog.Error(err)
		return nil, err
	}
	slog.Info(fmt.Sprintf("url:%s rs:%+v", url, rsModel))
	return rsModel, nil
}
