package utils

import "github.com/astaxie/beego"

func Trace(v ...interface{}) {
	beego.Trace(v)
}

func Debug(v ...interface{}) {
	beego.Debug(v)
}

func Info(v ...interface{}) {
	beego.Info(v)
}

func Warn(v ...interface{}) {
	beego.Warn(v)
}

func Error(v ...interface{}) {
	beego.Error(v)
}

func Critical(v ...interface{}) {
	beego.Critical(v)
}
