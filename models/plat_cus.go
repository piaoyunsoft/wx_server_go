package models

import "github.com/go-xorm/xorm"

type Platcus struct {
	Cusid   string `xorm:"not null pk VARCHAR(10)" json:"cusID"`
	Cusname string `xorm:"not null VARCHAR(30)" json:"cusName"`
}

type SeaPlatcus struct {
	SeaModel
	Cusid   string `json:"cusID"`
	Cusname string `json:"cusName"`
}

func (this *SeaPlatcus) where() *xorm.Session {
	session := x.NewSession().Table("platcus").Alias("a")
	if this.Cusid != "" {
		session.And("a.cusID = ?", this.Cusid)
	}
	return session
}

func (this *SeaPlatcus) GetAll() ([]Platcus, error) {
	items := make([]Platcus, 0)
	if err := this.getAll(this.where, &items); err == nil {
		return items, nil
	} else {
		return nil, err
	}
}
