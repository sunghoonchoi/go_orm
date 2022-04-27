package orm

import (
	sqlite "db_sqlite3"
	"fmt"
	"strings"
)

func UpdateRows(ddlObject *DDLObject, targets []TargetColumn, conditions ...SearchConditon) error {
	searchValues := make([]interface{}, 0)

	dml := []string{}
	dml = append(dml, "UPDATE ")
	dml = append(dml, ddlObject.TableName)
	dml = append(dml, " SET ")
	for i, v := range targets {
		dml = append(dml, v.ColumnName)
		dml = append(dml, " = ")
		dml = append(dml, "?") //fmt.Sprintf("%v", v.Value)
		searchValues = append(searchValues, v.Value)
		if i != len(targets)-1 {
			dml = append(dml, ", ")
		}
	}

	dml = append(dml, " WHERE ")
	for i, v := range conditions {
		if i != 0 && v.LogicOp != "" {
			dml = append(dml, " ")
			dml = append(dml, string(v.LogicOp))
			dml = append(dml, " ")
		}
		dml = append(dml, v.ColumnName)
		dml = append(dml, string(v.Operator))
		dml = append(dml, "?")

		searchValues = append(searchValues, v.Value)
	}

	queryString := strings.Join(dml, "")
	fmt.Println("queryString = ", queryString)

	err2 := sqlite.Update(queryString, ddlObject.db, searchValues)
	if err2 != nil {
		fmt.Println("insert err = ", err2)
		return err2
	}

	fmt.Println("results = ", queryString)
	return nil
}
