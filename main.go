package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	dbhost  string = "localhost"
	dbproto string = "tcp"
	dbport  int    = 3306
	dbname  string = "testing"
	dbuser  string = "testing"
	dbpass  string = "password"
)

type user struct {
	id        int       `mysql:"id"`
	email     string    `mysql:"email"`
	firstName string    `mysql:"firstName"`
	lastName  string    `mysql:"lastName"`
	createdAt time.Time `mysql:"createdAt"`
	lastLogin time.Time `mysql:"lastLogin"`
}

var err error
var db *sql.DB

func init() {
	validateConnection()
	initializeDatabase()

	time.Sleep(30)
	dropTable()
}

func main() {

}

func dropTable() {
	stmt := `DROP TABLE IF EXISTS users;`

	_, err := db.Exec(stmt)
	if err != nil {
		panic(err)
	}
}

func initializeDatabase() {
	stmt := `
	CREATE TABLE users (
		id INT NOT NULL AUTO_INCREMENT,
		email VARCHAR(255),
		firstName VARCHAR(255),
		lastName VARCHAR(255),
		createdAt TIME,
		lastLogin TIME,
		PRIMARY KEY (id)
	);
	`

	_, err := db.Exec(stmt)
	if err != nil {
		panic(err)
	}
}

func validateConnection() {
	cS := fmt.Sprintf("%v:%v@%v(%v:%v)/%v", dbuser, dbpass, dbproto, dbhost, dbport, dbname)
	log.Printf("Connecting to database on '%v'", dbhost)
	db, err = sql.Open("mysql", cS)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
}
