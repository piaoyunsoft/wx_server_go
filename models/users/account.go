package users

import (
	"errors"
	"strconv"
	"strings"
	"time"
	"wx_server_go/models"
	"wx_server_go/utils"
	"wx_server_go/utils/sqltool"

	"github.com/astaxie/beego/orm"
)

type Account struct {
	Unicode     string    `orm:"column(Unicode);pk" json:"unicode"`
	Password    string    `orm:"column(Password)" json:"password"`
	Status      string    `orm:"column(Status)" json:"status"`
	Mobile      string    `orm:"column(Mobile)" json:"mobile"`
	UserId      string    `orm:"column(UserId)" json:"userId"`
	AccountName string    `orm:"column(AccountName)" json:"accountName"`
	FromSys     string    `orm:"column(FromSys)" json:"-"`
	FromDeptId  string    `orm:"column(FromDeptId)" json:"cusId"`
	Remark      string    `orm:"column(Remark)" json:"remark"`
	VldDtm      time.Time `orm:"column(vldDtm)" json:"vldDtm"`
}

type AccountModel struct {
	Unicode     string    `orm:"column(Unicode);pk" json:"unicode"`
	Password    string    `orm:"column(Password)" json:"password"`
	Status      string    `orm:"column(Status)" json:"status"`
	Mobile      string    `orm:"column(Mobile)" json:"mobile"`
	UserId      string    `orm:"column(UserId)" json:"userId"`
	AccountName string    `orm:"column(AccountName)" json:"accountName"`
	FromSys     string    `orm:"column(FromSys)" json:"-"`
	FromDeptId  string    `orm:"column(FromDeptId)" json:"cusId"`
	Remark      string    `orm:"column(Remark)" json:"remark"`
	VldDtm      time.Time `orm:"column(vldDtm)" json:"vldDtm"`
	CusName     string    `orm:"column(cusName)" json:"cusName"`
}

const tablename = "account"

func (this *Account) TableName() string {
	return models.TableName(tablename)
}

func init() {
	orm.RegisterModel(new(Account))
}

func GetAccounts(query map[string]string, page int, size int) (total int64, res []Account, err error) {
	filterFunc := func(qs orm.QuerySeter) orm.QuerySeter {
		cond := orm.NewCondition()
		for k, v := range query {
			k = strings.Replace(k, ".", "__", -1)
			//			qs = qs.Filter(k, v)
			if k == "accountName" {
				cond = cond.And("AccountName__icontains", v)
			} else if k == "mobile" {
				cond = cond.And("Mobile__icontains", v)
			} else if k == "status" {
				cond = cond.And("Status", v)
			} else if k == "unicode" {
				cond = cond.And("Unicode", v)
			} else if k == "username" {
				cond = cond.AndCond(cond.And("AccountName", v).Or("Mobile", v))
			} else if k == "password" {
				cond = cond.And("Password", v)
			}
		}
		qs = qs.SetCond(cond)
		return qs
	}

	if total, err := sqltool.PageQuery_QS(new(Account), filterFunc, &res, page, size); err == nil {
		return total, res, nil
	} else {
		return 0, nil, err
	}
}

func GetPageAccounts(query map[string]string, page int, size int) (total int64, res []AccountModel, err error) {
	qb := sqltool.GetQueryBuilder()
	_w := "1=1"
	for k, v := range query {
		k = strings.Replace(k, ".", "__", -1)
		if k == "accountName" {
			_w += " and a.AccountName like '%" + sqltool.SqlFormat(v) + "%'"
		} else if k == "mobile" {
			_w += " and a.Mobile like '%" + sqltool.SqlFormat(v) + "%'"
		} else if k == "status" {
			_w += " and a.Status = '" + sqltool.SqlFormat(v) + "'"
		} else if k == "unicode" {
			_w += " and a.Unicode = '" + sqltool.SqlFormat(v) + "'"
		} else if k == "username" {
			_w += " (and a.AccountName = '" + sqltool.SqlFormat(v) + "' or Mobile = '" + sqltool.SqlFormat(v) + "'"
		} else if k == "password" {
			_w += " and a.Password = '" + sqltool.SqlFormat(v) + "'"
		}
	}
	qb.Select("a.*", "b.cusName").From("account as a").
		LeftJoin("platcus as b").On("b.cusID = a.FromDeptId").
		Where(_w).
		OrderBy("a.Unicode").
		Desc()

	if total, err = sqltool.PageQuery_QB(qb, &res, page, size); err == nil {
		return total, res, nil
	} else {
		return 0, nil, err
	}
}

func Login(username string, password string) (interface{}, error) {
	var query = make(map[string]string)
	query["username"] = username
	query["password"] = password
	query["status"] = "aa"
	if count, results, err := GetAccounts(query, 1, 1); err == nil {
		if count > 0 {
			return results[0], err
		} else {
			return nil, errors.New("no account data")
		}
	} else {
		return nil, err
	}
}

func GetAccountByUnicode(unicode string) (interface{}, error) {
	var query = make(map[string]string)
	query["unicode"] = unicode
	if count, results, err := GetAccounts(query, 1, 1); count > 0 && err == nil {
		return results[0], err
	} else {
		return nil, err
	}
}

func CheckAccount(query map[string]string) (bool, error) {
	unicode := ""
	for k, v := range query {
		if k == "unicode" {
			unicode = v
		}
	}
	filterFunc := func(qs orm.QuerySeter) orm.QuerySeter {
		for k, v := range query {
			k = strings.Replace(k, ".", "__", -1)
			//			qs = qs.Filter(k, v)
			if k == "accountName" {
				qs = qs.Filter("AccountName", v)
			} else if k == "mobile" {
				qs = qs.Filter("Mobile", v)
			}
		}
		return qs
	}

	var res []Account
	if err := sqltool.Query_QS(new(Account), filterFunc, &res); err == nil {
		if len(res) > 0 {
			for _, v := range res {
				if v.Unicode != unicode {
					return false, nil
				}
			}
			return true, nil
		} else {
			return true, nil
		}
	} else {
		return false, err
	}
}

func CreateAccount(item *Account) error {
	if id, err := sqltool.NewId(tablename); err == nil {
		item.Unicode = utils.Leftpad(strconv.Itoa(id), 12, 48)
		item.Password = "000000"
		return sqltool.Create(item)
	} else {
		return err
	}
}

func UpdateAccount(item *Account) error {
	o := orm.NewOrm()
	params := make(orm.Params)
	if item.Password != "" {
		params["Password"] = item.Password
	}
	if item.Status != "" {
		params["Status"] = item.Status
	}
	if item.Mobile != "" {
		params["Mobile"] = item.Mobile
	}
	if item.UserId != "" {
		params["UserId"] = item.UserId
	}
	if item.AccountName != "" {
		params["AccountName"] = item.AccountName
	}
	if item.FromSys != "" {
		params["FromSys"] = item.FromSys
	}
	if item.FromDeptId != "" {
		params["FromDeptId"] = item.FromDeptId
	}
	if item.Remark != "" {
		params["Remark"] = item.Remark
	}
	if !item.VldDtm.IsZero() {
		params["vldDtm"] = item.VldDtm
	}
	_, err := o.QueryTable("account").Filter("Unicode", item.Unicode).Update(params)
	if err != nil {
		utils.Error(err)
	}
	return err
}

func DelAccount(unicode string) error {
	filter := Account{Unicode: unicode}
	return sqltool.Delete(&filter)
}

func ResetPwd(item *Account) error {
	item.Password = "000000"
	return sqltool.Update(item, "Password")
}
