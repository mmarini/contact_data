package persistance

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = ""
	password = ""
	dbname   = "contacts"
)

func Insert(statement string, args ...interface{}) int {

	db, err := sql.Open("postgres", connectionString())
	CheckError(err)

	defer db.Close()

	var id int
	err = db.QueryRow(statement, args...).Scan(&id)
	CheckError(err)

	return id
}

func connectionString() string {
	return fmt.Sprintf("host=%s port=%d dbname=%s sslmode=disable", host, port, dbname)
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
