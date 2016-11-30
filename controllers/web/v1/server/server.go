package server

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"wx_server_go/constants"
	"wx_server_go/controllers/api"

	"github.com/astaxie/beego"
)

type ServerController struct {
	v1.BaseController
}

// @router /wxorder [get]
func (this *ServerController) GetAll() {
	orderid := this.GetString("orderid")
	serveraddr := beego.AppConfig.String("serveraddr")
	url := serveraddr + "wx/wxopenapi/QueryChargeOrder?orderid=" + orderid

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	this.Data["json"] = v1.ResData(constants.Success, string(body))
	this.ServeJSON()
}
