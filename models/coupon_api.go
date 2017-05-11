package models

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"pt_server/utils"
	"time"

	"strconv"

	"github.com/astaxie/beego"
	"github.com/pquerna/ffjson/ffjson"
)

type SeaCoupon struct {
	SeaModel
	ComID string
	Title string `json:"title"`
}

type Coupon struct {
	ID             string
	ComID          string  `json:"comID"`
	CouponType     string  `json:"couponType"`
	GetWay         string  `json:"getWay"`
	LogoUrl        string  `json:"logo_url"`
	Title          string  `json:"title"`
	Content        string  `json:"content"`
	Notice         string  `json:"notice"`
	UseScope       string  `json:"useScope"`
	DefCouponValue float32 `json:"defCouponValue"`
	Qty            int     `json:"qty"`
	TimeLimitWay   string  `json:"timeLimitWay"`
	TimeLimitValue string  `json:"timeLimitValue"`
	Status         string  `json:"status"`
}

type CouponModal struct {
	Data  []Coupon `json:"data"`
	Total int64    `json:"total"`
}

type SeaCouponItem struct {
	CouponID string `json:"couponID"`
}

type CouponItem struct {
	ItemID     string    `json:"itemID"`
	CouponID   string    `json:"couponID"`
	GetValue   float32   `json:"getValue"`
	GetDate    time.Time `json:"getDate"`
	GetVipcdid string    `json:"getVipcdid"`
	FromTb     string    `json:"fromTb"`
	FromPK     string    `json:"fromPK"`
	Status     string    `json:"status"`
}

func (this *SeaCoupon) GetCoupon() (int64, []Coupon, error) {
	serverAddress := beego.AppConfig.String("serveraddr")
	url := serverAddress + "wxopenapi/GetCoupons?title=" + this.Title + "&pageSize=" + strconv.Itoa(this.PageSize) + "&pageIndex=" + strconv.Itoa(this.PageIndex)

	resp, err := http.Get(url)
	if err != nil {
		utils.Error(err)
		return 0, nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		utils.Error(err)
		return 0, nil, err
	}

	rsModel := new(CouponModal)
	err = ffjson.Unmarshal(body, &rsModel)
	if err != nil {
		utils.Error(err)
		return 0, nil, err
	}
	utils.Info(fmt.Sprintf("url:%s rs:%+v", url, rsModel))
	return rsModel.Total, rsModel.Data, nil
}

func (this *SeaCouponItem) GetCouponItem() ([]CouponItem, error) {
	serverAddress := beego.AppConfig.String("serveraddr")
	url := serverAddress + "wxopenapi/GetCouponItems?couponID=" + this.CouponID

	resp, err := http.Get(url)
	if err != nil {
		utils.Error(err)
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		utils.Error(err)
		return nil, err
	}

	rsModel := make([]CouponItem, 0)
	err = ffjson.Unmarshal(body, &rsModel)
	if err != nil {
		utils.Error(err)
		return nil, err
	}
	utils.Info(fmt.Sprintf("url:%s rs:%+v", url, rsModel))
	return rsModel, nil
}
