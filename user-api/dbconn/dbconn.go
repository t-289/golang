package dbconn

import {
	"database/sql"
}

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "dev"
	dbPass := "123456"
	dbName := "dev"
	dbServer := "172.17.0.3"

	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@"+dbServer+"/"+dbName)
	if err != nil {
		panic(err.Error())
	}

	return db
}


func dbSelect(queryString string) {
	db := dbConn()

	selectDB, err := db.Query(queryString)

	if err != nil {
		panic(err.Error())
	}

	return selectDB
}
