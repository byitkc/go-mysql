package main

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"os"
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
var timeFormat = "2006-01-02 15:04:05.000000"

func init() {
	var u = user{
		email:     "brandon@byitkc.com",
		firstName: "Brandon",
		lastName:  "Young",
		createdAt: time.Now(),
		lastLogin: time.Now(),
	}
	defer dropTable()
	validateConnection()
	initializeDatabase()
	storeUser(u)
	user := retrieveUser("brandon@byitkc.com")
	fmt.Printf("%+v", user)

	log.Println("sleeping for 5 seconds")
	time.Sleep(time.Duration(5) * time.Second)
}

func main() {

}

func dropTable() {
	log.Printf("Cleaning up database")
	stmt := `DROP TABLE IF EXISTS users;`

	_, err := db.Exec(stmt)
	if err != nil {
		panic(err)
	}
	log.Printf("Successfully cleaned up database")
}

func initializeDatabase() {
	log.Printf("Initializing Database")
	f, err := os.Open("sql/init.sql")
	if err != nil {
		panic(err)
	}
	stmt, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(string(stmt))
	if err != nil {
		panic(err)
	}
	log.Printf("Successfully Initialized Database")
}

func storeUser(u user) {
	log.Println("Inserting test user")
	stmt := fmt.Sprintf(`
	INSERT INTO users(email, firstName, lastName, createdAt, lastLogin)
	VALUES ('%v', '%v', '%v', '%v', '%v');
	`, u.email, u.firstName, u.lastName, encodeTime(u.createdAt), encodeTime(u.lastLogin))
	log.Println(stmt)

	_, err = db.Exec(stmt)
	if err != nil {
		panic(err)
	}
	log.Println("Successfully inserted test user")
}

// retrieveUser will query the database for a user matching the provided email and
// will return the user
func retrieveUser(email string) user {
	var u user
	log.Println("Retrieving user from table")
	stmt := fmt.Sprintf(`SELECT * FROM users WHERE email = '%v'`, email)

	log.Println(stmt)
	err := db.QueryRow(stmt).Scan(&u.id, &u.email, &u.firstName, &u.lastName, &u.createdAt, &u.lastLogin)
	if err != nil {
		panic(err)
	}

	log.Println("Successfully retrieved user from table")
	return u
}

// valuedatConnect establishes a connection to the database and confirmes that we
// can reach the database as expected.
func validateConnection() {
	cS := fmt.Sprintf("%v:%v@%v(%v:%v)/%v?parseTime=true", dbuser, dbpass, dbproto, dbhost, dbport, dbname)
	log.Printf("Connecting to database on '%v'", dbhost)
	db, err = sql.Open("mysql", cS)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	log.Printf("Connected to database on '%v'", dbhost)
}

// encodeTime is used for encoding a time.Time object into a format that follows
// the standard format for storage in MySQL.
func encodeTime(time time.Time) string {
	return time.Format(timeFormat)
}
