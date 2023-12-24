package repository

import (
	"APIADMIN/config"
	"database/sql"
	"fmt"

	// "APIADMIN/config"
	//  "os"
	"strconv"
)

// var port string = os.Getenv("POSTGRES_PORT")
// var host string = os.Getenv("POSTGRES_HOST")
// var user string = os.Getenv("POSTGRES_USER")
// var password string = os.Getenv("POSTGRES_PASSWORD")
// var dbname string = os.Getenv("POSTGRES_DBNAME")

// var port string = "5432"
// var host string = "localhost"
// var user string = "postgres"
// var password string = "moonmk2004"
// var dbname string = "edu"

var Db *sql.DB

func InitDataBase() error {
	config := config.GetPostgres()
	fmt.Println("host: ",config.HostName)
	portint, _ := strconv.ParseInt(config.Port, 10, 64)
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		config.HostName, portint, config.User, config.Password, config.Dbname)

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
