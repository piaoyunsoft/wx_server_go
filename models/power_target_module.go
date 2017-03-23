package models

import "github.com/go-xorm/xorm"

type PowertargetModule struct {
	Id       int    `xorm:"not null pk INT(11)"`
	Pid      int    `xorm:"INT(11)"`
	Name     string `xorm:"VARCHAR(50)"`
	Sort     int    `xorm:"INT(11)"`
	Type     string `xorm:"not null CHAR(8)"`
	Target   string `xorm:"not null VARCHAR(255)"`
	Remark   string `xorm:"VARCHAR(255)"`
	Status   string `xorm:"CHAR(2)"`
	Actionid int    `xorm:"INT(11)"`
	Icon     string `xorm:"VARCHAR(20)"`
}

type ReqPowertargetModule struct {
	Id       int    `json:"id"`
	Pid      int    `json:"pid"`
	Name     string `json:"name"`
	Sort     int    `json:"sort"`
	Type     string `json:"type"`
	Target   string `json:"url"`
	Remark   string `json:"remark"`
	Status   string `json:"status"`
	Actionid int    `json:"-"`
	Icon     string `json:"icon"`
}

type SeaPowertargetModule struct {
	SeaModel
	Status string `json:"status"`
}

type PowertargetModuleModel struct {
	Id       int    `json:"id"`
	Pid      int    `json:"pid"`
	Name     string `json:"name"`
	Sort     int    `json:"sort"`
	Type     string `json:"type"`
	Target   string `json:"url"`
	Remark   string `json:"remark"`
	Status   string `json:"status"`
	Actionid int    `json:"-"`
	Icon     string `json:"icon"`
}

func (this *SeaPowertargetModule) where() *xorm.Session {
	session := x.NewSession().Table("powertarget_module").Alias("a")
	if this.Status != "" {
		session.And("a.status = ?", this.Status)
	}
	return session
}

func (this *SeaPowertargetModule) GetAll() ([]PowertargetModuleModel, error) {
	items := make([]PowertargetModuleModel, 0)
	if err := this.getAll(this.where, &items); err == nil {
		return items, nil
	} else {
		return nil, err
	}
}
