package dbconn

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func dbConn() (*sql.DB, error) {
	dbDriver := "mysql"
	dbUser := "dev"
	dbPass := "123456"
	dbName := "dev"
	dbServer := "172.17.0.3"

	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@"+dbServer+"/"+dbName)
	if err != nil {
		panic(err.Error())
	}

	return db, nil
}

func DBSelect(queryString string) (*sql.Rows, error) {
	db, err := dbConn()

	if err != nil {
		panic(err.Error())
	}

	selectDB, err := db.Query(queryString)

	if err != nil {
		panic(err.Error())
	}

	return selectDB, err
}
