package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/sangwoo89118/credential"
)

func main() {
	db, err := sql.Open("mysql", credential.UserName+":"+credential.Password+"@tcp(https://sangwoo.me)/phpmyadmin")

	if err != nil {
		panic(err.Error())
	}

	fmt.Println("database connected")

	defer db.Close()
}
