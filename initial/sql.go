package initial

import (
	"fmt"
	"time"
	_ "wx_server_go/models"
	_ "wx_server_go/models/cus"
	_ "wx_server_go/models/wx"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/ddliao/go-lib/slog"
	_ "github.com/go-sql-driver/mysql"
)

func InitSql() {
	user := beego.AppConfig.String("mysqluser")
	pass := beego.AppConfig.String("mysqlpass")
	host := beego.AppConfig.String("mysqlhost")
	port, err := beego.AppConfig.Int("mysqlport")
	dbname := beego.AppConfig.String("mysqldb")

	if nil != err {
		port = 3306
	}
	orm.RegisterDriver("mysql", orm.DRMySQL)
	if beego.AppConfig.String("runmode") == "dev" {
		orm.Debug = true
		//		host = "localhost"
		//		pass = "123456"
	} else {
		//		conn = fmt.Sprintf("%s:%s@/%s?charset=utf8", user, pass, dbname)
	}
	conn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&timeout=3s&readTimeout=3s&writeTimeout=3s", user, pass, host, port, dbname)
	slog.Info(conn)
	conn = conn + "&loc=Local"
	orm.DefaultTimeLoc = time.Local
	err = orm.RegisterDataBase("default", "mysql", conn, 2000, 1000)
	if err != nil {
		slog.Error("RegisterDataBase", err)
	}
	//	orm.RunSyncdb("default", false, true) // true 改成false，如果表存在则会给出提示，如果改成false则不会提示 ， 这句话没有会报主键不存在的错误
}
