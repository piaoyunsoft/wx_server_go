package model

import (
	"time"
)

type Account struct {
	Unicode     string    `xorm:"not null pk CHAR(12)"`
	Password    string    `xorm:"VARCHAR(255)"`
	Status      string    `xorm:"not null CHAR(2)"`
	Mobile      string    `xorm:"not null VARCHAR(20)"`
	Userid      string    `xorm:"VARCHAR(20)"`
	Accountname string    `xorm:"VARCHAR(30)"`
	Fromsys     string    `xorm:"VARCHAR(6)"`
	Fromdeptid  string    `xorm:"VARCHAR(50)"`
	Remark      string    `xorm:"VARCHAR(100)"`
	Vlddtm      time.Time `xorm:"DATETIME"`
}

type Alipaycb struct {
	Id               int64  `xorm:"BIGINT(20)"`
	PaymentType      string `xorm:"not null VARCHAR(1)"`
	Subject          string `xorm:"not null VARCHAR(256)"`
	TradeNo          string `xorm:"not null VARCHAR(32)"`
	BuyerEmail       string `xorm:"not null VARCHAR(100)"`
	GmtCreate        string `xorm:"not null VARCHAR(19)"`
	NotifyType       string `xorm:"not null VARCHAR(20)"`
	Quantity         string `xorm:"not null VARCHAR(10)"`
	OutTradeNo       string `xorm:"not null VARCHAR(64)"`
	NotifyTime       string `xorm:"not null VARCHAR(19)"`
	SellerId         string `xorm:"not null VARCHAR(16)"`
	TradeStatus      string `xorm:"VARCHAR(32)"`
	IsTotalFeeAdjust string `xorm:"VARCHAR(1)"`
	TotalFee         string `xorm:"VARCHAR(20)"`
	GmtPayment       string `xorm:"VARCHAR(19)"`
	SellerEmail      string `xorm:"VARCHAR(100)"`
	GmtClose         string `xorm:"VARCHAR(19)"`
	Price            string `xorm:"VARCHAR(20)"`
	BuyerId          string `xorm:"VARCHAR(16)"`
	NotifyId         string `xorm:"default '' VARCHAR(100)"`
	UseCoupon        string `xorm:"VARCHAR(2)"`
}

type Datapushtask struct {
	Id         int64     `xorm:"BIGINT(20)"`
	Tasktype   string    `xorm:"not null VARCHAR(30)"`
	Tasklinkid string    `xorm:"not null VARCHAR(120)"`
	Op         string    `xorm:"VARCHAR(12)"`
	Targetid   string    `xorm:"VARCHAR(20)"`
	Createdate time.Time `xorm:"DATETIME"`
	Pushdate   time.Time `xorm:"DATETIME"`
	Pushtimes  int       `xorm:"default 0 INT(11)"`
	Remark     string    `xorm:"default '' VARCHAR(2000)"`
	Doneflg    string    `xorm:"default 'f' CHAR(1)"`
}

type Dictindex struct {
	Dictcode     string `xorm:"not null pk CHAR(3)"`
	Dictname     string `xorm:"not null VARCHAR(40)"`
	Issysdict    string `xorm:"CHAR(1)"`
	Itemcodelen  int    `xorm:"INT(11)"`
	Dicttypename string `xorm:"VARCHAR(40)"`
	Status       string `xorm:"not null CHAR(2)"`
	Picfileid    string `xorm:"VARCHAR(255)"`
}

type Dictitem struct {
	Dictcode  string `xorm:"not null pk VARCHAR(3)"`
	Itemcode  string `xorm:"not null pk VARCHAR(10)"`
	Itemname  string `xorm:"not null VARCHAR(40)"`
	Status    string `xorm:"not null CHAR(2)"`
	Issysdict string `xorm:"CHAR(1)"`
}

type Platcus struct {
	Cusid       string `xorm:"not null pk VARCHAR(10)"`
	Cusname     string `xorm:"not null VARCHAR(30)"`
	Cusdes      string `xorm:"VARCHAR(255)"`
	Memcode     string `xorm:"VARCHAR(8)"`
	Regionid    string `xorm:"VARCHAR(30)"`
	Addr        string `xorm:"VARCHAR(60)"`
	Postcode    string `xorm:"CHAR(6)"`
	Tel         string `xorm:"VARCHAR(30)"`
	Entpcls     string `xorm:"CHAR(3)"`
	Taxno       string `xorm:"VARCHAR(30)"`
	Licenceno   string `xorm:"VARCHAR(30)"`
	Legalperson string `xorm:"VARCHAR(30)"`
	Bizscope    string `xorm:"VARCHAR(255)"`
	Status      string `xorm:"not null CHAR(2)"`
	Createdate  string `xorm:"not null CHAR(14)"`
	Createby    string `xorm:"VARCHAR(20)"`
	Auditdate   string `xorm:"not null CHAR(14)"`
	Auditperson string `xorm:"VARCHAR(10)"`
	Signkey     string `xorm:"VARCHAR(100)"`
}

type Platcuswx struct {
	Cusid          string `xorm:"not null pk VARCHAR(10)"`
	Wxname         string `xorm:"not null VARCHAR(32)"`
	Wxloginid      string `xorm:"VARCHAR(40)"`
	Wxloginpwd     string `xorm:"VARCHAR(400)"`
	Appid          string `xorm:"VARCHAR(40)"`
	Appsecret      string `xorm:"VARCHAR(400)"`
	Serverurl      string `xorm:"VARCHAR(100)"`
	Servertoken    string `xorm:"VARCHAR(100)"`
	Encodingaeskey string `xorm:"VARCHAR(100)"`
	Status         string `xorm:"CHAR(2)"`
	Srvflg         string `xorm:"CHAR(1)"`
	Srvbegdate     string `xorm:"CHAR(14)"`
	Srvenddate     string `xorm:"CHAR(14)"`
	WepayMchid     string `xorm:"VARCHAR(32)"`
	WepayKey       string `xorm:"VARCHAR(400)"`
	WepayNotifyurl string `xorm:"VARCHAR(100)"`
}

type Platcuswxtask struct {
	Id           int64     `xorm:"BIGINT(20)"`
	Fromcusid    string    `xorm:"VARCHAR(12)"`
	Towxopenid   string    `xorm:"not null VARCHAR(40)"`
	Tocusmbrid   string    `xorm:"VARCHAR(48)"`
	Tocusmbrname string    `xorm:"VARCHAR(200)"`
	Toaccountid  string    `xorm:"VARCHAR(48)"`
	Tplcode      string    `xorm:"VARCHAR(16)"`
	Tplid        string    `xorm:"VARCHAR(100)"`
	Type         string    `xorm:"VARCHAR(20)"`
	Msgbody      string    `xorm:"TEXT"`
	Createdtm    time.Time `xorm:"DATETIME"`
	Status       string    `xorm:"CHAR(2)"`
	Senddtm      string    `xorm:"CHAR(14)"`
	Wxmsgid      string    `xorm:"VARCHAR(50)"`
	Receivedtm   string    `xorm:"CHAR(10)"`
	Remark       string    `xorm:"VARCHAR(200)"`
	Url          string    `xorm:"VARCHAR(200)"`
}

type PowertargetModule struct {
	Id       int    `xorm:"not null pk INT(11)"`
	Pid      int    `xorm:"INT(11)"`
	Name     string `xorm:"VARCHAR(50)"`
	Sort     int    `xorm:"INT(11)"`
	Type     string `xorm:"not null CHAR(8)"`
	Target   string `xorm:"not null VARCHAR(255)"`
	Remark   string `xorm:"VARCHAR(255)"`
	Status   string `xorm:"CHAR(2)"`
	Actionid int    `xorm:"INT(11)"`
	Icon     string `xorm:"VARCHAR(20)"`
}

type Sequence struct {
	Tablename string `xorm:"not null pk VARCHAR(64)"`
	Id        int64  `xorm:"default 1 BIGINT(20)"`
}

type Vipcls struct {
	Comid     string `xorm:"not null pk VARCHAR(60)"`
	Vipclsid  string `xorm:"not null pk VARCHAR(10)"`
	Vipclsdes string `xorm:"not null VARCHAR(60)"`
	Status    string `xorm:"not null CHAR(2)"`
}

type Vipgiftexch struct {
	Exchid        string    `xorm:"not null pk VARCHAR(32)"`
	Comid         string    `xorm:"VARCHAR(10)"`
	Mbrid         string    `xorm:"VARCHAR(32)"`
	Wxopenid      string    `xorm:"VARCHAR(40)"`
	Giftcode      string    `xorm:"not null VARCHAR(32)"`
	Giftname      string    `xorm:"VARCHAR(255)"`
	Getway        string    `xorm:"VARCHAR(10)"`
	Mailaddr      string    `xorm:"VARCHAR(255)"`
	Mailpsnname   string    `xorm:"VARCHAR(255)"`
	Mailpsnmobile string    `xorm:"VARCHAR(24)"`
	Exchqty       string    `xorm:"DECIMAL(10,2)"`
	Usescore      int       `xorm:"INT(11)"`
	Status        string    `xorm:"CHAR(2)"`
	Mailstatus    string    `xorm:"CHAR(2)"`
	Createdate    time.Time `xorm:"DATETIME"`
	Changedate    time.Time `xorm:"DATETIME"`
}

type Vipgiftlist struct {
	Giftcode    string    `xorm:"not null pk VARCHAR(32)"`
	Cusid       string    `xorm:"VARCHAR(10)"`
	Gifttype    string    `xorm:"VARCHAR(10)"`
	Giftamt     string    `xorm:"DECIMAL(10,2)"`
	Giftname    string    `xorm:"VARCHAR(255)"`
	Getway      string    `xorm:"VARCHAR(10)"`
	Getbrief    string    `xorm:"VARCHAR(255)"`
	Vlddays     time.Time `xorm:"DATE"`
	Giftpic     string    `xorm:"VARCHAR(255)"`
	Scoreneed   int       `xorm:"INT(11)"`
	Stkqty      string    `xorm:"DECIMAL(10,2)"`
	Deptid      string    `xorm:"VARCHAR(12)"`
	Begdate     time.Time `xorm:"DATETIME"`
	Enddate     time.Time `xorm:"DATETIME"`
	Status      string    `xorm:"CHAR(2)"`
	Makeperson  string    `xorm:"VARCHAR(20)"`
	Makedate    time.Time `xorm:"DATETIME"`
	Auditperson string    `xorm:"VARCHAR(20)"`
	Auditdate   time.Time `xorm:"DATETIME"`
}

type Wxchargelist struct {
	Id         string `xorm:"not null pk VARCHAR(32)"`
	Name       string `xorm:"not null VARCHAR(60)"`
	Comid      string `xorm:"not null VARCHAR(60)"`
	Vipclsid   string `xorm:"not null VARCHAR(10)"`
	Payamt     string `xorm:"not null DECIMAL(10,2)"`
	Getrealamt string `xorm:"DECIMAL(10,2)"`
	Getgiftamt string `xorm:"DECIMAL(10,2)"`
	Status     string `xorm:"CHAR(2)"`
}

type Wxchargeodr struct {
	Odrid       string    `xorm:"not null pk VARCHAR(64)"`
	Payptf      string    `xorm:"VARCHAR(12)"`
	Subuid      string    `xorm:"VARCHAR(48)"`
	Comid       string    `xorm:"VARCHAR(60)"`
	Wxopenid    string    `xorm:"VARCHAR(40)"`
	Mbrid       string    `xorm:"VARCHAR(64)"`
	Amt         string    `xorm:"DECIMAL(10,2)"`
	Payptfodrid string    `xorm:"VARCHAR(64)"`
	Paytime     time.Time `xorm:"DATETIME"`
	Status      string    `xorm:"CHAR(2)"`
	Errmsg      string    `xorm:"VARCHAR(128)"`
	Createdate  time.Time `xorm:"DATETIME"`
	Changedate  time.Time `xorm:"DATETIME"`
	Getrealamt  string    `xorm:"DECIMAL(10,2)"`
	Getgiftamt  string    `xorm:"DECIMAL(10,2)"`
}

type Wxpaycb struct {
	Id                 int64  `xorm:"BIGINT(20)"`
	Appid              string `xorm:"not null VARCHAR(32)"`
	MchId              string `xorm:"not null VARCHAR(32)"`
	DeviceInfo         string `xorm:"VARCHAR(32)"`
	NonceStr           string `xorm:"not null VARCHAR(32)"`
	Sign               string `xorm:"not null VARCHAR(32)"`
	SignType           string `xorm:"VARCHAR(32)"`
	ResultCode         string `xorm:"not null VARCHAR(16)"`
	ErrCode            string `xorm:"VARCHAR(32)"`
	ErrCodeDes         string `xorm:"VARCHAR(128)"`
	Openid             string `xorm:"not null VARCHAR(128)"`
	IsSubscribe        string `xorm:"CHAR(1)"`
	TradeType          string `xorm:"not null VARCHAR(16)"`
	BankType           string `xorm:"VARCHAR(16)"`
	TotalFee           int    `xorm:"not null INT(11)"`
	SettlementTotalFee int    `xorm:"INT(11)"`
	FeeType            string `xorm:"VARCHAR(8)"`
	CashFee            int    `xorm:"not null INT(11)"`
	CashFeeType        string `xorm:"VARCHAR(16)"`
	CouponFee          int    `xorm:"INT(11)"`
	CouponCount        int    `xorm:"INT(11)"`
	CouponType         int    `xorm:"INT(11)"`
	CouponIds          string `xorm:"VARCHAR(255)"`
	CouponFees         string `xorm:"VARCHAR(255)"`
	TransactionId      string `xorm:"not null VARCHAR(32)"`
	OutTradeNo         string `xorm:"not null VARCHAR(32)"`
	Attach             string `xorm:"VARCHAR(128)"`
	TimeEnd            string `xorm:"not null VARCHAR(14)"`
}

type Wxsubscribe struct {
	Uid               string    `xorm:"not null pk VARCHAR(48)"`
	Wxopenid          string    `xorm:"not null unique(Index_wxAttTmp) VARCHAR(40)"`
	Wxunionid         string    `xorm:"VARCHAR(48)"`
	Comwxid           string    `xorm:"not null unique(Index_wxAttTmp) VARCHAR(100)"`
	Comid             string    `xorm:"VARCHAR(60)"`
	Wxnickname        string    `xorm:"VARCHAR(100)"`
	Wxsex             string    `xorm:"CHAR(1)"`
	Subscribed        string    `xorm:"CHAR(1)"`
	Wxsubscribetime   time.Time `xorm:"not null DATETIME"`
	Wxunsubscribetime time.Time `xorm:"DATETIME"`
	Wxcountry         string    `xorm:"VARCHAR(40)"`
	Wxprovince        string    `xorm:"VARCHAR(40)"`
	Wxcity            string    `xorm:"VARCHAR(40)"`
	Wxheadimgurl      string    `xorm:"VARCHAR(255)"`
	Wxsubscribecount  int       `xorm:"INT(11)"`
	Wxbrief           string    `xorm:"VARCHAR(600)"`
	Binddate          time.Time `xorm:"DATETIME"`
	Bindway           string    `xorm:"VARCHAR(30)"`
	Mbrid             string    `xorm:"VARCHAR(32)"`
	Aduitdate         time.Time `xorm:"DATETIME"`
	Aduitperson       string    `xorm:"VARCHAR(30)"`
	Status            string    `xorm:"not null CHAR(2)"`
	Mbrname           string    `xorm:"VARCHAR(64)"`
	Mbrtype           string    `xorm:"VARCHAR(10)"`
	Mobile            string    `xorm:"VARCHAR(30)"`
	Idno              string    `xorm:"VARCHAR(20)"`
	Birthdate         string    `xorm:"CHAR(8)"`
	Addr              string    `xorm:"VARCHAR(255)"`
	Createdate        time.Time `xorm:"not null DATETIME"`
	Changedate        time.Time `xorm:"DATETIME"`
	Applybrief        string    `xorm:"VARCHAR(255)"`
	Vipclsid          string    `xorm:"VARCHAR(10)"`
}

type Wxtplmsgcustpl struct {
	Id      int    `xorm:"not null pk INT(11)"`
	Cusid   string `xorm:"VARCHAR(10)"`
	Tplid   string `xorm:"VARCHAR(100)"`
	Tplcode string `xorm:"VARCHAR(16)"`
}

type Wxtplmsgtpllib struct {
	Tplcode    string `xorm:"not null pk VARCHAR(16)"`
	Title      string `xorm:"VARCHAR(20)"`
	Msgdtl     string `xorm:"VARCHAR(200)"`
	Sample     string `xorm:"VARCHAR(200)"`
	Useexample string `xorm:"VARCHAR(200)"`
	Status     string `xorm:"CHAR(2)"`
}
