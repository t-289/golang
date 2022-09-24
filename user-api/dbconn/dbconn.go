package dbconn

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func dbConn() (*sql.DB, error) {
	dbDriver := "mysql"
	dbUser := "dev"
	dbPass := "123456"
	dbName := "dev"
	dbServer := "172.17.0.3:3306"

	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@"+"tcp("+dbServer+")/"+dbName)
	if err != nil {
		fmt.Println("Erro ao conectar")
	}

	return db, nil
}

func DBSelect(queryString string) (*sql.Rows, error) {
	db, err := dbConn()

	if err != nil {
		fmt.Println("Erro ao executar o select")
	}

	selectDB, err := db.Query(queryString)

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	return selectDB, err
}

func DBDelete(queryString string) (int64, error) {
	db, err := dbConn()

	if err != nil {
		fmt.Println("Erro ao executar o select")
	}

	deleteDB, err := db.Exec(queryString)

	if err != nil {
		panic(err.Error())
	}

	result, _ := deleteDB.RowsAffected()

	defer db.Close()

	return result, err
}
