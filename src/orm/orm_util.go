package orm

import (
	"reflect"
)

type COMP_OPER string
type LOGIC_OPER string

const (
	BIGGER       COMP_OPER = ">"
	LESS         COMP_OPER = "<"
	EQUAL        COMP_OPER = "="
	NOT_EQUAL    COMP_OPER = "!="
	BIGGER_EQUAL COMP_OPER = ">="
	LESS_EQUAL   COMP_OPER = "<="
)

const (
	AND LOGIC_OPER = "AND"
	OR  LOGIC_OPER = "OR"
)

type SearchConditon struct {
	ColumnName string
	Value      interface{}
	Operator   COMP_OPER
	LogicOp    LOGIC_OPER
}

type TargetColumn struct {
	ColumnName string
	Value      interface{}
}

func Contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func GetSqlTypeOf(input reflect.Type) string {
	typeString := input.String()
	var sqlString string

	switch typeString {
	case "orm.OrmEngine":
		sqlString = "SKIP"
	case "[]uint8":
		sqlString = "BLOB"
	case "string":
		sqlString = "TEXT"
	case "int":
		sqlString = "NUMERIC"
	case "float":
		sqlString = "REAL"
	default:
		sqlString = "ERROR"
	}

	return sqlString
}

func GetAttrValue(obj interface{}, fieldName string) reflect.Value {
	pointToStruct := reflect.ValueOf(obj) // addressable
	curStruct := pointToStruct.Elem()
	if curStruct.Kind() != reflect.Struct {
		panic("not struct")
	}
	curField := curStruct.FieldByName(fieldName) // type: reflect.Value
	if !curField.IsValid() {
		panic("not found:" + fieldName)
	}
	return curField
}
