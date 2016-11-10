package v1

type ErrorController struct {
	BaseController
}

func (c *ErrorController) Error404() {
	c.Data["json"] = ResCode(404)
	c.ServeJSON()
}

func (c *ErrorController) Error501() {
	c.Data["json"] = ResCode(501)
	c.ServeJSON()
}

func (c *ErrorController) ErrorDb() {
	c.Data["json"] = ResCode(100)
	c.ServeJSON()
}
