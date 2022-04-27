package orm

import (
	"database/sql"
	sqlite "db_sqlite3"
	"fmt"
	"reflect"
	"strings"
)

type DDLObject struct {
	TableName      string
	DefineElements []string
	DefineKinds    []string
	PrimaryKeys    []int
	db             *sql.DB
}

func RegisterInOrmEngine(object interface{}, dbPath string, args ...int) (*DDLObject, error) {
	var ddlObject *DDLObject = new(DDLObject)

	tag := reflect.TypeOf(object)
	ddlObject.TableName = tag.Name()

	for _, index := range args {
		Contains(ddlObject.PrimaryKeys, index)
		ddlObject.PrimaryKeys = append(ddlObject.PrimaryKeys, index)
	}

Loop:
	for i := 0; i < tag.NumField(); i++ {

		fieldName := tag.Field(i).Name
		sqlType := GetSqlTypeOf(tag.Field(i).Type)

		if sqlType == "SKIP" {
			continue Loop
		}

		ddlObject.DefineElements = append(ddlObject.DefineElements, fieldName)
		ddlObject.DefineKinds = append(ddlObject.DefineKinds, sqlType)

	}
	fmt.Println("values :", ddlObject.DefineElements)
	fmt.Println("columens :", ddlObject.DefineKinds)
	fmt.Println("primary keys :", ddlObject.PrimaryKeys)

	_, err := CreateTable(ddlObject, dbPath)
	return ddlObject, err
}

func GetColumnNameAtIndex(ddlObject *DDLObject, index int) string {
	columnName := ddlObject.DefineElements[index]
	return columnName
}

func CreateTable(ddlObject *DDLObject, dbPath string) (string, error) {

	ddl := []string{}
	ddl = append(ddl, "CREATE TABLE IF NOT EXISTS")
	ddl = append(ddl, " ")
	ddl = append(ddl, ddlObject.TableName)
	ddl = append(ddl, " ")
	ddl = append(ddl, "(")
	for i := range ddlObject.DefineKinds {
		ddl = append(ddl, "'")
		ddl = append(ddl, ddlObject.DefineElements[i])
		ddl = append(ddl, "'")
		ddl = append(ddl, " ")
		ddl = append(ddl, ddlObject.DefineKinds[i])
		ddl = append(ddl, ",")
	}
	ddl = append(ddl, "PRIMARY KEY(")
	for j := range ddlObject.PrimaryKeys {
		ddl = append(ddl, "'")
		ddl = append(ddl, ddlObject.DefineElements[j])
		ddl = append(ddl, "'")
		if j != (len(ddlObject.PrimaryKeys) - 1) {
			ddl = append(ddl, ",")
		}
	}
	ddl = append(ddl, ")")
	ddl = append(ddl, ")")

	db, err := sqlite.InitDB(dbPath)
	if err != nil {
		return "", err
	} else {
		ddlObject.db = db
	}

	queryString := strings.Join(ddl, "")

	err2 := sqlite.Create(queryString, db)
	if err != nil {
		return "", err2
	} else {
		fmt.Println("create table success")
	}

	return queryString, nil
}
