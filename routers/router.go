// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"wx_server_go/controllers/common"
	"wx_server_go/controllers/web/v1/charge"
	"wx_server_go/controllers/web/v1/cus"
	"wx_server_go/controllers/web/v1/park"
	"wx_server_go/controllers/web/v1/server"
	"wx_server_go/controllers/web/v1/sys"
	"wx_server_go/controllers/web/v1/users"
	"wx_server_go/controllers/web/v1/wx"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/web/v1",
		beego.NSNamespace("/user", beego.NSInclude(&users.UserController{})),
		beego.NSNamespace("/account", beego.NSInclude(&users.AccountController{})),
		beego.NSNamespace("/cus", beego.NSInclude(&cus.CusController{})),
		beego.NSNamespace("/subscribe", beego.NSInclude(&wx.WxSubscribeController{})),
		beego.NSNamespace("/wxtask", beego.NSInclude(&wx.WxTaskController{})),
		beego.NSNamespace("/cusmbr", beego.NSInclude(&cus.CusMbrController{})),
		beego.NSNamespace("/park", beego.NSInclude(&park.ParkController{})),
		beego.NSNamespace("/region", beego.NSInclude(&park.RegionController{})),
		beego.NSNamespace("/device", beego.NSInclude(&park.DeviceController{})),
		beego.NSNamespace("/traffic", beego.NSInclude(&park.TrafficController{})),
		beego.NSNamespace("/charge", beego.NSInclude(&charge.ChargeController{})),
		beego.NSNamespace("/vip", beego.NSInclude(&cus.VipController{})),
		beego.NSNamespace("/sys", beego.NSInclude(&sys.SysController{})),
		beego.NSNamespace("/module", beego.NSInclude(&sys.ModuleController{})),
		beego.NSNamespace("/server", beego.NSInclude(&server.ServerController{})),
	)
	nsCom := beego.NewNamespace("/com",
		beego.NSNamespace("/file", beego.NSInclude(&common.FileController{})),
	)
	beego.AddNamespace(ns)
	beego.AddNamespace(nsCom)
}
