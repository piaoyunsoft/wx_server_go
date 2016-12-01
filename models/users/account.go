package users

import (
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

const tablename = "account"

func (this *Account) TableName() string {
	return models.TableName(tablename)
}

func init() {
	orm.RegisterModel(new(Account))
}

func GetAccounts(query map[string]string, page int, size int) (total int64, res []Account, err error) {
	filterFunc := func(qs orm.QuerySeter) orm.QuerySeter {
		for k, v := range query {
			k = strings.Replace(k, ".", "__", -1)
			//			qs = qs.Filter(k, v)
			if k == "accountName" {
				qs = qs.Filter("AccountName__icontains", v)
			} else if k == "mobile" {
				qs = qs.Filter("Mobile__icontains", v)
			} else if k == "status" {
				qs = qs.Filter("Status", v)
			}
		}
		return qs
	}

	if total, err := sqltool.PageQuery_QS(new(Account), filterFunc, &res, page, size); err == nil {
		return total, res, nil
	} else {
		return 0, nil, err
	}
}

func CreateAccount(item *Account) error {
	if id, err := sqltool.NewId(tablename); err == nil {
		item.Unicode = utils.Leftpad(string(id), 12, 0)
		return sqltool.Create(item)
	} else {
		return err
	}
}

func UpdateAccount(item *Account) error {
	return sqltool.Update(item)
}

func DelAccount(unicode string) error {
	filter := Account{Unicode: unicode}
	return sqltool.Delete(&filter)
}
