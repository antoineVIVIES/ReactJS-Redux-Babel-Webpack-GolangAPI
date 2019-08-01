package db

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type Database struct {
	db *sql.DB
}

func CreateQuery() {
	fmt.Println("aeazeaze")
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func OpenDB() {
	db, err := sql.Open("mysql", "test:root@/accrohome")
	checkErr(err)
	res, err := db.Query("SELECT * FROM users;")
	checkErr(err)
	res.Next()
	return db
}
