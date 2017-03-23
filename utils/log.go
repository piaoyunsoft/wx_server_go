package utils

import (
	"fmt"

	"github.com/astaxie/beego"
)

func Trace(v ...interface{}) {
	fmt.Println(v)
	beego.Trace(v)
}

func Debug(v ...interface{}) {
	fmt.Println(v)
	beego.Debug(v)
}

func Info(v ...interface{}) {
	fmt.Println(v)
	beego.Info(v)
}

func Warn(v ...interface{}) {
	fmt.Println(v)
	beego.Warn(v)
}

func Error(v ...interface{}) {
	if len(v) >= 1 && v[0] != nil {
		fmt.Println(v)
		beego.Error(v)
	}
}

func Critical(v ...interface{}) {
	fmt.Println(v)
	beego.Critical(v)
}
