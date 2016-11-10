package initial

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

func InitLog() {
	//	beego.SetLevel(logs.LevelDebug)
	//	beego.SetLogFuncCall(true)
	//	logs.SetLogFuncCallDepth(6)
	beego.SetLogger(logs.AdapterMultiFile, `{"filename":"logs/pt.log"}`)
}
