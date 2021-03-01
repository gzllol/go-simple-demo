package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)


func GetSqlClient() *sql.DB {
	db, err := sql.Open("mysql", "root:example@tcp(127.0.0.1:23306)/test")
	if err != nil {
		fmt.Println(err)
	}
	return db
}

func TestInsertStudent(t *testing.T) {
	db := GetSqlClient()
	r, e := db.Exec("INSERT INTO student(name) VALUES (?)", "Tom")
	fmt.Println(r, e)
	if r != nil {
		id, _ := r.LastInsertId()
		ra, _ := r.RowsAffected()
		fmt.Println(id, ra, e)
	}
}
