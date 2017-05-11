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
	"wx_server_go/controllers/web/v1/server"

	"wx_server_go/controllers/web/v1"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/web/v1",
		//beego.NSNamespace("/user", beego.NSInclude(&users.UserController{})),
		beego.NSNamespace("/account", beego.NSInclude(&v1.AccountController{})),
		beego.NSNamespace("/cus", beego.NSInclude(&v1.CusController{})),
		beego.NSNamespace("/subscribe", beego.NSInclude(&v1.WxSubscribeController{})),
		beego.NSNamespace("/wxtask", beego.NSInclude(&v1.WxTaskController{})),
		//beego.NSNamespace("/cusmbr", beego.NSInclude(&cus.CusMbrController{})),
		beego.NSNamespace("/charge", beego.NSInclude(&v1.ChargeController{})),
		beego.NSNamespace("/vip", beego.NSInclude(&v1.VipController{})),
		beego.NSNamespace("/sys", beego.NSInclude(&v1.SysController{})),
		beego.NSNamespace("/module", beego.NSInclude(&v1.ModuleController{})),
		beego.NSNamespace("/server", beego.NSInclude(&server.ServerController{})),
		beego.NSNamespace("/statistics", beego.NSInclude(&v1.StatisticsController{})),
		beego.NSNamespace("/vipcls", beego.NSInclude(&v1.VipClsController{})),
		beego.NSNamespace("/wxchargelist", beego.NSInclude(&v1.WxChargeListController{})),
		beego.NSNamespace("/coupon", beego.NSInclude(&v1.CouponController{})),

		//beego.NSNamespace("/park", beego.NSInclude(&park.ParkController{})),
		//beego.NSNamespace("/region", beego.NSInclude(&park.RegionController{})),
		//beego.NSNamespace("/device", beego.NSInclude(&park.DeviceController{})),
		//beego.NSNamespace("/traffic", beego.NSInclude(&park.TrafficController{})),
	)
	nsCom := beego.NewNamespace("/com",
		beego.NSNamespace("/file", beego.NSInclude(&common.FileController{})),
	)
	beego.AddNamespace(ns)
	beego.AddNamespace(nsCom)
}
