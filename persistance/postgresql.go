package persistance

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = ""
	password = ""
	dbname   = "contacts"
)

func Insert(tableName string, columns []string, args ...interface{}) int {

	db, err := sql.Open("postgres", connectionString())
	CheckError(err)

	defer db.Close()

	var id int
	statement := prepareInsertStatement(tableName, columns)

	err = db.QueryRow(statement, args...).Scan(&id)
	CheckError(err)

	return id
}

func prepareInsertStatement(tableName string, columns []string) string {
	statement := fmt.Sprintf(`INSERT INTO "%s"("%s") VALUES (`, tableName, strings.Join(columns[:], "\",\""))

	var positions []string
	for pos := range columns {
		positions = append(positions, fmt.Sprintf("$%d", pos+1))
	}

	statement = statement + strings.Join(positions, ",") + ") RETURNING id"

	return statement
}

func connectionString() string {
	return fmt.Sprintf("host=%s port=%d dbname=%s sslmode=disable", host, port, dbname)
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
