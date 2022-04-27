//delete from people where id = 1"
package orm

import (
	sqlite "db_sqlite3"
	"fmt"
	"strings"
)

func DeleteRows(ddlObject *DDLObject, conditions ...SearchConditon) error {
	// results := make([]interface{}, 0)

	// SELECT * FROM TEST_TABLE WHERE UserId = $1
	searchValues := make([]interface{}, 0)

	dml := []string{}
	dml = append(dml, "DELETE FROM ")

	dml = append(dml, ddlObject.TableName)
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
	fmt.Println("[DEBUG]delete query = ", queryString)

	err2 := sqlite.Delete(queryString, ddlObject.db, searchValues)
	if err2 != nil {
		fmt.Println("insert err = ", err2)
		return err2
	}

	return nil
}

func DeleteAll(ddlObject *DDLObject) error {
	dml := []string{}
	dml = append(dml, "DELETE FROM ")
	dml = append(dml, ddlObject.TableName)

	queryString := strings.Join(dml, "")
	fmt.Println("[DEBUG]delete query = ", queryString)

	err2 := sqlite.Delete(queryString, ddlObject.db, nil)
	if err2 != nil {
		fmt.Println("insert err = ", err2)
		return err2
	}

	return nil
}
