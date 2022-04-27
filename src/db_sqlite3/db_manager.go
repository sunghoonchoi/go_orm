package db_sqlite3

import (
	"database/sql"
	"fmt"
	"log"
	"reflect"

	_ "github.com/mattn/go-sqlite3"
)

type TEST_TABLE struct {
	Id       int
	UserId   string
	Password string
}

func InitDB(file string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func Create(query string, db *sql.DB) error {
	_, e := db.Exec(query)
	if e != nil {
		return e
	}
	return nil
}

func Insert(query string, db *sql.DB, args ...interface{}) error {
	tx, _ := db.Begin()
	stmt, _ := tx.Prepare(query)
	var err error
	for index, name := range args {
		_, err = stmt.Exec(name, index)
		checkError(err)
	}

	tx.Commit()
	return nil
}

func Insert2(query string, db *sql.DB, args []interface{}) error {
	tx, _ := db.Begin()
	stmt, _ := tx.Prepare(query)
	var err error
	// for index, name := range args {
	_, err = stmt.Exec(args...)
	checkError(err)
	// }

	tx.Commit()
	return nil
}

func Update(query string, db *sql.DB, args []interface{}) error {
	tx, _ := db.Begin()
	stmt, _ := tx.Prepare(query)
	var err error
	_, err = stmt.Exec(args...)
	checkError(err)
	// }

	tx.Commit()
	return nil
}

func Delete(query string, db *sql.DB, args []interface{}) error {
	tx, _ := db.Begin()
	stmt, _ := tx.Prepare(query)
	var err error
	_, err = stmt.Exec(args...)
	checkError(err)
	// }

	tx.Commit()
	return nil
}

func Select(query string, db *sql.DB, condition []interface{}, resultNames []string, resultTypes []string) ([]interface{}, error) {

	results := make([]interface{}, 0)
	fmt.Println("[QueryBefore]query = ", query)
	fmt.Println("[QueryBefore]condition = ", condition)

	row, err := db.Query(query, condition...)
	if err != nil {
		log.Fatal(err)
	}

	defer row.Close()
	for row.Next() { // Iterate and fetch the records from result cursor

		rowValues := make([]interface{}, 0)
		for i := range resultTypes {
			var value interface{}
			switch resultTypes[i] {
			case "NUMERIC":
				value = int(0)
				rowValues = append(rowValues, &value)
			case "REAL":
				value = float32(0.0)
				rowValues = append(rowValues, &value)
			case "BLOB":
				value = make([]byte, 1000)
				rowValues = append(rowValues, &value)
			case "TEXT":
				value = ""
				rowValues = append(rowValues, &value)
			}
		}
		row.Scan(rowValues...)

		//my loginc make map for the result and append it to the targetArray
		m := make(map[string]interface{})
		for i, v := range rowValues {
			value := reflect.ValueOf(v).Elem().Interface()

			m[resultNames[i]] = value
		}

		results = append(results, m)
	}

	return results, nil
}

func SelectOne(query string, db *sql.DB, condition interface{}) (TEST_TABLE, error) {

	var result TEST_TABLE

	rows := db.QueryRow(query, condition)
	err := rows.Scan(&result.Id, &result.UserId, &result.Password)
	if err != nil {
		return TEST_TABLE{}, err
	}

	return result, nil
}

func checkError(err error) {
	if err != nil {
		// panic(err)
		fmt.Println(err)
	}
}
