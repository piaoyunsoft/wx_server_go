package sqltool

import (
	"wx_server_go/utils"

	"github.com/astaxie/beego/orm"
)

type IdStruct struct {
	Newid int `orm:"column(newid)"`
}

func GetQueryBuilder() orm.QueryBuilder {
	qb, _ := orm.NewQueryBuilder("mysql")
	return qb
}

func PageQuery_QS(tableName interface{}, fun func(orm.QuerySeter) orm.QuerySeter, res interface{}, page int, size int) (total int64, err error) {
	qs := GetQuerySeter(tableName)
	qs = fun(qs)
	if total, err = qs.Count(); err == nil {
		offset := (page - 1) * size
		if _, err := qs.Limit(size, offset).All(res); err == nil {
			return total, nil
		}
	}
	utils.Error(err)
	return 0, err
}

func PageQuery_QB(qb orm.QueryBuilder, result interface{}, page int, limit int) (total int64, err error) {
	o := orm.NewOrm()
	sql := qb.String()
	if total, err = o.Raw(sql).QueryRows(result); err == nil {
		offset := (page - 1) * limit
		qb.Limit(limit).Offset(offset)
		sql = qb.String()
		if _, err = o.Raw(sql).QueryRows(result); err == nil {
			return total, nil
		}
	}
	utils.Error(err)
	return 0, err
}

func GetQuerySeter(tableName interface{}) orm.QuerySeter {
	o := orm.NewOrm()
	qs := o.QueryTable(tableName)
	return qs
}

func Query_QS(tableName interface{}, fun func(orm.QuerySeter) orm.QuerySeter, res interface{}) error {
	qs := GetQuerySeter(tableName)
	qs = fun(qs)
	_, err := qs.All(res)
	return err
}

func Update(item interface{}) error {
	o := orm.NewOrm()
	_, err := o.Update(item)
	if err != nil {
		utils.Error(err)
	}
	return err
}

func Delete(item interface{}) error {
	o := orm.NewOrm()
	_, err := o.Delete(item)
	if err != nil {
		utils.Error(err)
	}
	return err
}

func Create(item interface{}) error {
	o := orm.NewOrm()
	if _, err := o.Insert(item); err == nil {
		return nil
	} else {
		utils.Error(err)
		return err
	}
}

func NewId(tablename string) (int, error) {
	o := orm.NewOrm()
	var id IdStruct
	err := o.Raw("call pro_sequence(?)", tablename).QueryRow(&id)
	if err != nil {
		utils.Error(err)
	}
	return id.Newid, err
}
