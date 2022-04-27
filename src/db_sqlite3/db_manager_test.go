package db_sqlite3

import (
	"log"
	"testing"
)

//sqlite3 test
func Test_InitDB(t *testing.T) {
	db, err := InitDB(".hello.db")
	if err != nil {
		log.Fatal(err)
		t.Fatal("db open fail")

	}

	createTableQuery := `
		CREATE TABLE IF NOT EXISTS TEST_TABLE (
		id integer PRIMARY KEY autoincrement,
		userId text,
		password text,
		UNIQUE (id, userId)
		)
	`
	err = Create(createTableQuery, db)
	if err != nil {
		log.Fatal(err)
		t.Fatal("create table error")
	}
}

func Test_Insert(t *testing.T) {
	db, err := InitDB(".hello.db")
	if err != nil {
		log.Fatal(err)
		t.Fatal("db open fail")
	}

	insertQuery := "INSERT INTO TEST_TABLE (userId,password) VALUES (?,?)"
	id := "choi"
	password := "qwer1234!@#$"

	err = Insert(insertQuery, db, id, password)
	if err != nil {
		log.Fatal(err)
		t.Fatal("insert error")
	}
}

func Test_Select(t *testing.T) {
	db, err := InitDB(".hello.db")
	if err != nil {
		log.Fatal(err)
		t.Fatal("db open fail")
	}

	selectQuery := "SELECT * FROM TEST_TABLE WHERE UserId = $1"

	_, err = SelectOne(selectQuery, db, 1)
	if err != nil {
		log.Fatal(err)
		t.Fatal("select error")
	}
}
