package models

import "github.com/go-xorm/xorm"

type Dictitem struct {
	Dictcode  string `xorm:"not null pk VARCHAR(3)"`
	Itemcode  string `xorm:"not null pk VARCHAR(10)"`
	Itemname  string `xorm:"not null VARCHAR(40)"`
	Status    string `xorm:"not null CHAR(2)"`
	Issysdict string `xorm:"CHAR(1)"`
}

type ReqDictitem struct {
	Dictcode  string `json:"dictcode"`
	Itemcode  string `json:"itemcode"`
	Itemname  string `json:"itemname"`
	Status    string `json:"status"`
	Issysdict string `json:"issysdict"`
}

type SeaDictitem struct {
	SeaModel
	Dictcode string `json:"dictcode"`
}

type DictitemModel struct {
	Dictcode  string `json:"dictcode"`
	Itemcode  string `json:"itemcode"`
	Itemname  string `json:"itemname"`
	Status    string `json:"status"`
	Issysdict string `json:"issysdict"`
}

func (this *SeaDictitem) where() *xorm.Session {
	session := x.NewSession().Table("dictitem").Alias("a")
	if this.Dictcode != "" {
		session.And("a.dictcode = ?", this.Dictcode)
	}
	return session
}

func (this *SeaDictitem) GetAll() ([]DictitemModel, error) {
	items := make([]DictitemModel, 0)
	if err := this.getAll(this.where, &items); err == nil {
		return items, nil
	} else {
		return nil, err
	}
}
