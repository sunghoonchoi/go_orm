package orm

import (
	sqlite "db_sqlite3"
	"fmt"
	"strings"
)

func SearchRows(ddlObject *DDLObject, conditions ...SearchConditon) ([]interface{}, error) {
	// searchResults := make([]interface{}, 0)
	searchValues := make([]interface{}, 0)

	dml := []string{}
	dml = append(dml, "SELECT * FROM ")

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

	results, err2 := sqlite.Select(queryString, ddlObject.db, searchValues, ddlObject.DefineElements, ddlObject.DefineKinds)
	if err2 != nil {
		fmt.Println("select err = ", err2)
		return nil, err2
	}

	return results, nil
}
