package sqltool

import (
	"wx_server_go/utils"

	"github.com/astaxie/beego/orm"
)

func GetQueryBuilder() orm.QueryBuilder {
	qb, _ := orm.NewQueryBuilder("mysql")
	return qb
}

func PageQuery_QS(qs orm.QuerySeter, result interface{}, page int, limit int) (total int64, err error) {
	if total, err = qs.Count(); err == nil {
		offset := (page - 1) * limit
		if _, err := qs.Limit(limit, offset).All(result); err == nil {
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
