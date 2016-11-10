package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["wx_server_go/controllers/web/v1/wx:WxSubscribeController"] = append(beego.GlobalControllerRouter["wx_server_go/controllers/web/v1/wx:WxSubscribeController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

	beego.GlobalControllerRouter["wx_server_go/controllers/web/v1/wx:WxTaskController"] = append(beego.GlobalControllerRouter["wx_server_go/controllers/web/v1/wx:WxTaskController"],
		beego.ControllerComments{
			Method: "GetAll",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			Params: nil})

}
