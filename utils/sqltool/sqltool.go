package sqltool

import (
	"strings"

	"time"

	"github.com/astaxie/beego/orm"
	"github.com/ddliao/go-lib/slog"
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
	slog.Error(err)
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
	slog.Error(err)
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

func Update(item interface{}, cols ...string) error {
	o := orm.NewOrm()
	_, err := o.Update(item, cols...)
	if err != nil {
		slog.Error(err)
	}
	return err
}

func Delete(item interface{}) error {
	o := orm.NewOrm()
	_, err := o.Delete(item)
	if err != nil {
		slog.Error(err)
	}
	return err
}

func Create(item interface{}) error {
	o := orm.NewOrm()
	if _, err := o.Insert(item); err == nil {
		return nil
	} else {
		slog.Error(err)
		return err
	}
}

func NewId(tablename string) (int, error) {
	o := orm.NewOrm()
	var id IdStruct
	err := o.Raw("call pro_sequence(?)", tablename).QueryRow(&id)
	if err != nil {
		slog.Error(err)
	}
	return id.Newid, err
}

func SqlFormat(to_match_str string) string {
	temp := strings.Replace(to_match_str, "--", "", -1)
	temp = strings.Replace(to_match_str, "'", "", -1)
	return temp
}

func StrToTime(s string, format string) time.Time {
	tempFormat := strings.Replace(format, "yyyy", "2006", -1)
	tempFormat = strings.Replace(tempFormat, "MM", "01", -1)
	tempFormat = strings.Replace(tempFormat, "dd", "02", -1)
	tempFormat = strings.Replace(tempFormat, "HH", "15", -1)
	tempFormat = strings.Replace(tempFormat, "mm", "04", -1)
	tempFormat = strings.Replace(tempFormat, "ss", "05", -1)
	loc, _ := time.LoadLocation("Local")
	formatTime, _ := time.ParseInLocation(tempFormat, s, loc)
	return formatTime
}

func TimeToStr(t time.Time, format string) string {
	tempFormat := strings.Replace(format, "yyyy", "2006", -1)
	tempFormat = strings.Replace(tempFormat, "MM", "01", -1)
	tempFormat = strings.Replace(tempFormat, "dd", "02", -1)
	tempFormat = strings.Replace(tempFormat, "HH", "15", -1)
	tempFormat = strings.Replace(tempFormat, "mm", "04", -1)
	tempFormat = strings.Replace(tempFormat, "ss", "05", -1)
	formatTime := t.Format(tempFormat)
	return formatTime
}
