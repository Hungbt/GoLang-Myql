package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	name string `json:"name"`
}


func main() {
	db, err := sql.Open("mysql", "root:12345@tcp(127.0.0.1:3306)/sys")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	// GET database
	results, err := db.Query("select * from sys.users")
	if err != nil {
		panic(err.Error())
	}
	for results.Next() {
		var user User

		err = results.Scan(&user.name)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(user.name)
	}


	// --> insert database
	// insert, err := db.Query("INSERT INTO sys.users VALUES ('BuiTuanHung')")
	// if err != nil {
	// 	panic(err.Error())
	// }
	// defer insert.Close()
}
