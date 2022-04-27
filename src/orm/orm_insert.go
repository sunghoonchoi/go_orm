package orm

import (
	sqlite "db_sqlite3"
	"errors"
	"fmt"
	"reflect"
	"strings"
)

type DMLObject struct {
	DefineElements []string
	Values         []interface{}
}

func InsertToTable(ddlObject *DDLObject, object interface{}) error {
	var dmlObject = new(DMLObject)

	//step 1. check the object try to insert is fit to the ddl
	tag := reflect.TypeOf(object)
	if ddlObject.TableName != tag.Elem().Name() {
		fmt.Println("ddlObject.TableName", ddlObject.TableName)
		fmt.Println("tag.Name()", tag.Name())
		return errors.New("type error")
	}

	if ddlObject == nil {
		return errors.New("object to be inserted is null")
	}

	//TODO: primary key null check
	// now db_manager can handle this error

	//step 2. make query
	dml := []string{}
	dml = append(dml, "INSERT INTO ")

	dml = append(dml, ddlObject.TableName)
	dml = append(dml, " (")
	for i := range ddlObject.DefineElements {
		dml = append(dml, "'")
		dml = append(dml, ddlObject.DefineElements[i])
		dml = append(dml, "'")
		if i != (len(ddlObject.DefineElements) - 1) {
			dml = append(dml, ",")
		}
	}
	dml = append(dml, ") VALUES (")

	for i := range ddlObject.DefineElements {
		dml = append(dml, "?")
		val := GetAttrValue(object, ddlObject.DefineElements[i])
		fmt.Println("the result of GetAttrValue.Interface() ", val.Interface())
		dmlObject.Values = append(dmlObject.Values, val.Interface())

		if i != (len(ddlObject.DefineElements) - 1) {
			dml = append(dml, ",")
		}
	}
	dml = append(dml, ")")

	//step 3. run query
	queryString := strings.Join(dml, "")

	fmt.Println("[Debug] queryString = ", queryString)
	err := sqlite.Insert2(queryString, ddlObject.db, dmlObject.Values)
	if err != nil {
		fmt.Println("insert err = ", err)
		return err
	}

	return nil
}
