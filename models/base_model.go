package models

import (
	"fmt"
	"time"

	"errors"

	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	"github.com/jdongdong/go-lib/slog"
)

var (
	x     *xorm.Engine
	DbCfg struct {
		Host, Port, User, Pwd, Db string
	}
)

type Model struct {
}

type SeaModel struct {
	PageIndex int `json:"pageIndex"`
	PageSize  int `json:"pageSize"`
}

type TreeModel struct {
	Key      string      `json:"key"`
	Title    string      `json:"title"`
	Type     string      `json:"type"`
	Checked  bool        `json:"checked"`
	Children []TreeModel `json:"children"`
}

type MenuModel struct {
	Key   string      `json:"key"`
	Href  string      `json:"href"`
	Name  string      `json:"name"`
	Icon  string      `json:"icon"`
	Child []MenuModel `json:"child"`
}

func init() {
	//tables = append(tables, new(SysUser))

	LoadConfig()

	err := NewEngine()
	slog.Error(err)
}

func LoadConfig() {
	DbCfg.Host = beego.AppConfig.DefaultString("mysqlhost", "localhost")
	DbCfg.Port = beego.AppConfig.DefaultString("mysqlport", "3306")
	DbCfg.User = beego.AppConfig.DefaultString("mysqluser", "root")
	DbCfg.Pwd = beego.AppConfig.String("mysqlpass")
	DbCfg.Db = beego.AppConfig.String("mysqldb")
}

func getEngine() (*xorm.Engine, error) {
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true", DbCfg.User, DbCfg.Pwd, DbCfg.Host, DbCfg.Port, DbCfg.Db)
	slog.Info(connStr)
	return xorm.NewEngine("mysql", connStr)
}

func SetEngine() (err error) {
	x, err = getEngine()
	if err != nil {
		return fmt.Errorf("fail to connect to database: %v", err)
	}
	x.SetMapper(core.GonicMapper{})
	x.TZLocation = time.Local
	x.SetMaxOpenConns(2000)
	x.SetMaxIdleConns(1000)
	x.DB().SetConnMaxLifetime(time.Second * 5)
	x.ShowExecTime(true)

	x.ShowSQL(true)
	x.Logger().SetLevel(core.LOG_INFO)
	return nil
}

func NewEngine() (err error) {
	if err = SetEngine(); err != nil {
		return err
	}
	return nil
}

func toLike(s string) string {
	return fmt.Sprintf("%%%s%%", s)
}

func toPaging(pageIndex, pageSize int) (limit, start int) {
	if pageIndex <= 0 {
		pageIndex = 1
	}
	if pageSize <= 0 {
		pageSize = 10
	}
	limit = pageSize
	start = (pageIndex - 1) * pageSize
	return limit, start
}

type getSession func() *xorm.Session

func (this *SeaModel) getPagingSel(session getSession, sel string, bean interface{}, item interface{}) (int64, error) {
	total, err := session().Count(bean)
	if err == nil {
		err = session().
			Limit(toPaging(this.PageIndex, this.PageSize)).
			Select(sel).
			Find(item)
		if err == nil {
			return total, nil
		} else {
			slog.Error(err)
			return 0, err
		}
	} else {
		slog.Error(err)
		return 0, err
	}
}

func (this *SeaModel) getPaging(session getSession, bean interface{}, item interface{}) (int64, error) {
	total, err := session().Count(bean)
	if err == nil {
		err = session().
			Limit(toPaging(this.PageIndex, this.PageSize)).
			Find(item)
		if err == nil {
			return total, nil
		} else {
			slog.Error(err)
			return 0, err
		}
	} else {
		slog.Error(err)
		return 0, err
	}
}

func (this *SeaModel) getAll(session getSession, item interface{}) error {
	if err := session().Find(item); err == nil {
		return nil
	} else {
		slog.Error(err)
		return err
	}
}

func (this *SeaModel) getOne(session getSession, item interface{}) error {
	has, err := session().Get(item)
	if err != nil {
		return err
	} else if !has {
		err = errors.New("has no data")
		slog.Error(err)
		return err
	} else {
		return nil
	}
}

func (this *SeaModel) getOneSel(session getSession, sel string, item interface{}) error {
	has, err := session().Select(sel).Get(item)
	if err != nil {
		slog.Error(err)
		return err
	} else if !has {
		err = errors.New("has no data")
		slog.Error(err)
		return err
	} else {
		return nil
	}
}

func doPost(url string, i interface{}) ([]byte, error) {
	str, _ := json.Marshal(i)
	resp, err := http.Post(url, "application/json", strings.NewReader(string(str)))
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
	slog.Info(fmt.Sprintf("url:%s,body:%+s, result:%s", url, string(str), string(body)))
	return body, nil
}

func doGet(url string) ([]byte, error) {
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
	slog.Info(fmt.Sprintf("url:%s,result:%s", url, string(body)))
	return body, nil
}
