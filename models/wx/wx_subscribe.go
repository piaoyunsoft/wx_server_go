package wx

//import (
//	"strings"
//	"time"
//
//	"wx_server_go/utils"
//
//	"io/ioutil"
//	"net/http"
//
//	"encoding/json"
//
//	"errors"
//
//	"fmt"
//
//	"github.com/astaxie/beego"
//	"github.com/astaxie/beego/orm"
//	"github.com/ddliao/go-lib/tool"
//)
//
//type Wxsubscribe struct {
//	Uid               string    `orm:"column(uid);size(48);pk;" json:"uid"`
//	WxOpenId          string    `orm:"column(wxOpenId);size(40);" json:"wxOpenId"`
//	WxUnionId         string    `orm:"column(wxUnionId);size(48);null;" json:"wxUnionId"`
//	ComWxID           string    `orm:"column(comWxID);size(100);" json:"comWxID"`
//	ComID             string    `orm:"column(comID);size(60);null;" json:"comID"`
//	WxNickName        string    `orm:"column(wxNickName);size(100);null;" json:"wxNickName"`
//	WxSex             string    `orm:"column(wxSex);size(1);null;" json:"wxSex"`
//	Subscribed        string    `orm:"column(subscribed);size(1);null;" json:"subscribed"`
//	WxSubscribeTime   time.Time `orm:"column(wxSubscribeTime);" json:"wxSubscribeTime"`
//	WxUnsubscribeTime time.Time `orm:"column(wxUnsubscribeTime);null;" json:"wxUnsubscribeTime"`
//	WxCountry         string    `orm:"column(wxCountry);size(40);null;" json:"wxCountry"`
//	WxProvince        string    `orm:"column(wxProvince);size(40);null;" json:"wxProvince"`
//	WxCity            string    `orm:"column(wxCity);size(40);null;" json:"wxCity"`
//	WxHeadImgUrl      string    `orm:"column(wxHeadImgUrl);size(255);null;" json:"wxHeadImgUrl"`
//	WxSubscribeCount  int       `orm:"column(wxSubscribeCount);size(11);null;" json:"wxSubscribeCount"`
//	WxBrief           string    `orm:"column(wxBrief);size(600);null;" json:"wxBrief"`
//	BindDate          time.Time `orm:"column(BindDate);null;" json:"BindDate"`
//	BindWay           string    `orm:"column(BindWay);size(30);null;" json:"BindWay"`
//	MbrId             string    `orm:"column(mbrId);size(32);null;" json:"mbrId"`
//	AduitDate         time.Time `orm:"column(aduitDate);null;" json:"aduitDate"`
//	AduitPerson       string    `orm:"column(aduitPerson);size(30);null;" json:"aduitPerson"`
//	Status            string    `orm:"column(status);size(2);" json:"status"`
//	MbrName           string    `orm:"column(mbrName);size(64);null;" json:"mbrName"`
//	MbrType           string    `orm:"column(mbrType);size(10);null;" json:"mbrType"`
//	Mobile            string    `orm:"column(mobile);size(30);null;" json:"mobile"`
//	Idno              string    `orm:"column(idno);size(20);null;" json:"idno"`
//	BirthDate         string    `orm:"column(birthDate);size(8);null;" json:"birthDate"`
//	Addr              string    `orm:"column(addr);size(255);null;" json:"addr"`
//	CreateDate        time.Time `orm:"column(createDate);" json:"createDate"`
//	ChangeDate        time.Time `orm:"column(changeDate);null;" json:"changeDate"`
//	ApplyBrief        string    `orm:"column(ApplyBrief);size(255);null;" json:"apply_brief"`
//}
//
//func init() {
//	orm.RegisterModel(new(Wxsubscribe))
//}
//
//func ReadWxSubscribeList(query map[string]string, page int64, limit int64) (total int64, res []Wxsubscribe, err error) {
//	o := orm.NewOrm()
//	qs := o.QueryTable(new(Wxsubscribe))
//	cond := orm.NewCondition()
//	// query k=v
//	for k, v := range query {
//		k = strings.Replace(k, ".", "__", -1)
//		//		qs = qs.Filter(k, v)
//		if k == "WxNickName" {
//			cond = cond.And(k+"__icontains", v)
//		} else if k == "begin" {
//			cond = cond.And("WxSubscribeTime__gte", v)
//		} else if k == "end" {
//			cond = cond.And("WxSubscribeTime__lte", v)
//		} else if k == "status" {
//			cond = cond.And(k, v)
//		} else if k == "key" {
//			cond = cond.AndCond(cond.And("mbrName__icontains", v).Or("mobile__icontains", v))
//		}
//	}
//	qs = qs.SetCond(cond)
//	qs = qs.OrderBy("-createDate")
//	if total, err = qs.Count(); err == nil {
//		offset := (page - 1) * limit
//		if _, err := qs.Limit(limit, offset).All(&res); err == nil {
//			return total, res, nil
//		}
//	}
//
//	return 0, nil, err
//}
//
//func UpdateSubscribeByUID(item *Wxsubscribe) error {
//	o := orm.NewOrm()
//	params := make(orm.Params)
//	if item.MbrId != "" {
//		params["mbrId"] = item.MbrId
//	}
//	_, err := o.QueryTable("wxsubscribe").Filter("uid", item.Uid).Update(params)
//	if err != nil {
//		utils.Error(err)
//	}
//	return err
//}
//
//type SendMsgModel struct {
//	ClientId     string `json:"clientId"`
//	Timestamp    int64  `json:"timestamp"`
//	Sign         string `json:"sign"`
//	MbrId        string `json:"mbrId"`
//	MbrOwner     string `json:"mbrOwner"`
//	TplCode      string `json:"tplCode"`
//	FormattedMsg string `json:"formattedMsg"`
//}
//
//type TplMsg_OPENTM405904851 struct {
//	First    string `json:"first"`
//	Keyword1 string `json:"keyword1"`
//	Keyword2 string `json:"keyword2"`
//	Remark   string `json:"remark"`
//}
//
//type UpdateOpenIDModel struct {
//	MbrID  string `json:"mbrID"`
//	OpenID string `json:"openID"`
//}
//
//func BindCardByUID(item *Wxsubscribe) (string, error) {
//	o := orm.NewOrm()
//	filter := new(Wxsubscribe)
//	filter.Uid = item.Uid
//	err := o.Read(filter)
//	if err != nil {
//		utils.Error(err)
//	}
//
//	//验证会员卡号是否有效
//	serverAddress := beego.AppConfig.String("serveraddr")
//	url := serverAddress + "wxopenapi/CheckMbrID?mbrID=" + filter.MbrId
//	body, err := doGet(url)
//	if err != nil {
//		utils.Error(err)
//		return "", err
//	}
//	if string(body) != "success" {
//		err = errors.New("会员卡号无效或已被绑定")
//		utils.Error(err)
//		return "会员卡号无效或已被绑定", nil
//	}
//	//更新商场库会员openid
//	url = serverAddress + "wxopenapi/UpdateOpenID"
//	uptOpenID := new(UpdateOpenIDModel)
//	uptOpenID.MbrID = filter.MbrId
//	uptOpenID.OpenID = filter.WxOpenId
//	_, err = doPost(url, uptOpenID)
//	if err != nil {
//		utils.Error(err)
//		return "", err
//	}
//	if string(body) != "success" {
//		err = errors.New("同步商场数据失败")
//		utils.Error(err)
//		return "同步商场数据失败", nil
//	}
//
//	params := make(orm.Params)
//	params["status"] = "aa"
//	params["BindWay"] = "后台绑定"
//	params["BindDate"] = time.Now()
//	_, err = o.QueryTable("wxsubscribe").Filter("uid", item.Uid).Update(params)
//	if err != nil {
//		utils.Error(err)
//		return "", err
//	}
//
//	tpl := new(TplMsg_OPENTM405904851)
//	tpl.First = "您好，会员卡绑定成功"
//	tpl.Keyword1 = filter.MbrId
//	tpl.Keyword2 = tool.TimeToStr(time.Now(), "MM月dd日 HH时mm分")
//	tpl.Remark = "恭喜您成为天赐时代绿卡会员，自此您将获得我们为您提供的购物积分、积分换礼、特价通知、停车折价、免费杂志、周年庆受邀等服务"
//	str, _ := json.Marshal(tpl)
//
//	msg := new(SendMsgModel)
//	msg.ClientId = "web"
//	msg.FormattedMsg = string(str)
//	msg.MbrId = filter.MbrId
//	msg.MbrOwner = filter.ComID
//	msg.Sign = ""
//	msg.Timestamp = time.Now().Unix()
//	msg.TplCode = "OPENTM405904851"
//	str, _ = json.Marshal(msg)
//
//	url = serverAddress + "wxopenapi/SendTplMsg"
//	_, err = doPost(url, msg)
//	if err != nil {
//		utils.Error(err)
//		return "", err
//	}
//	return "success", nil
//}
//
//func doPost(url string, i interface{}) ([]byte, error) {
//	str, _ := json.Marshal(i)
//	resp, err := http.Post(url, "application/json", strings.NewReader(string(str)))
//	if err != nil {
//		utils.Error(err)
//		return nil, err
//	}
//
//	defer resp.Body.Close()
//	body, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		utils.Error(err)
//		return nil, err
//	}
//	utils.Info(fmt.Sprintf("url:%s,result:%s", url, string(body)))
//	return body, nil
//}
//
//func doGet(url string) ([]byte, error) {
//	resp, err := http.Get(url)
//	if err != nil {
//		utils.Error(err)
//		return nil, err
//	}
//	defer resp.Body.Close()
//	body, err := ioutil.ReadAll(resp.Body)
//	if err != nil {
//		utils.Error(err)
//		return nil, err
//	}
//	utils.Info(fmt.Sprintf("url:%s,result:%s", url, string(body)))
//	return body, nil
//}
