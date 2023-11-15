package repository

import (
	"database/sql"
	"fmt"
	"strconv"
)

var port string = "5432"
var host string = "localhost"
var user string = "postgres"
var password string = "moonmk2004"
var dbname string = "edu"

var Db *sql.DB

func InitDataBase() error {
	portint, _ := strconv.ParseInt(port, 10, 64)
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, portint, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}
	fmt.Println("Successfully connected!")

	Db = db
	fmt.Print(Db)

	return nil
}
