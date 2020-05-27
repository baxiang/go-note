package dialect

import (
	"fmt"
	"reflect"
)

type mysql struct {
}
var _ Dialect =(*mysql)(nil)

func init(){
	RegisterDialect("mysql",&mysql{})
}

func (m *mysql)DataTypeOf(t reflect.Value)string{
	switch t.Kind() {
	case reflect.Bool:
		return "bool"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uintptr:
		return "integer"
	case reflect.Int64, reflect.Uint64:
		return "bigint"
	case reflect.Float32, reflect.Float64:
		return "real"
	case reflect.String:
		return "text"
	case reflect.Array, reflect.Slice:
		return "blob"
	}
	panic(fmt.Sprintf("invalid sql type %s (%s)", t.Type().Name(), t.Kind()))
}
func (m *mysql)TableExistSQL(tableName string)(string, []interface{}){
	args := []interface{}{"test",tableName}
	return "SELECT table_name FROM INFORMATION_SCHEMA.TABLES WHERE table_schema = ? AND table_name = ?", args
}

