package database

import (
	"errors"
	"github.com/jinzhu/gorm"
	"gocms/pkg/logger"
	"reflect"
	"strconv"
)

type DataBase struct {
	*gorm.DB
}

func (that *DataBase) Where(query interface{}, param ...interface{}) *DataBase {
	return &DataBase{
		that.DB.Where(query, param),
	}
}

func (that *DataBase) WhereCombOr(query interface{}) *DataBase {

	t := reflect.TypeOf(query)
	v := reflect.ValueOf(query)

	flag := false
	db := that.DB
	for i := 0; i < t.NumField(); i++ {
		where := t.Field(i).Name
		param := fieldValueStr(v.Field(i))
		if flag {
			db = db.Or(where+"=?", param)
			continue
		}
		flag = true
		db = db.Where(where+"=?", param)
	}
	return &DataBase{db}
}

func fieldValueStr(value reflect.Value) (res string) {

	kind := value.Kind()

	switch kind {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:

		res = strconv.FormatInt(value.Int(), 10)

	case reflect.Float64, reflect.Float32:

		res = strconv.FormatFloat(value.Float(), 'f', 10, 64)

	case reflect.String:

		res = value.String()

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:

		res = strconv.FormatUint(value.Uint(), 10)

	case reflect.Bool:

		res = strconv.FormatBool(value.Bool())

	default:

		logger.PanicError(errors.New("unsupported type "+kind.String()), "fieldValueStr db.go:40", true)
	}
	return
}
