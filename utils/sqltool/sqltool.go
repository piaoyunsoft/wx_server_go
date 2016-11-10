package sqltool

import "github.com/astaxie/beego/orm"

func GetQueryBuilder() orm.QueryBuilder {
	qb, _ := orm.NewQueryBuilder("mysql")
	return qb
}

func PageQuery(qb orm.QueryBuilder, result interface{}, page int, limit int) (total int64, err error) {
	o := orm.NewOrm()
	sql := qb.String()
	if total, err = o.Raw(sql).QueryRows(result); err == nil {
		offset := (page - 1) * limit
		qb.Limit(limit).Offset(offset)
		sql = qb.String()
		if _, err = o.Raw(sql).QueryRows(result); err == nil {
			return total, nil
		} else {
			return 0, err
		}
	} else {
		return total, err
	}
}
