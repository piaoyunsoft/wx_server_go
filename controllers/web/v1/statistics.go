package v1

import (
	"wx_server_go/constants"
	"wx_server_go/models"
)

type StatisticsController struct {
	BaseController
}

// @router / [get]
func (this *StatisticsController) GetAll() {
	if rs, err := models.GetStatistics(); err == nil {
		this.Data["json"] = ResData(constants.Success, rs)
	} else {
		this.Data["json"] = ResData(constants.DBError, nil)
	}
	this.ServeJSON()
}
