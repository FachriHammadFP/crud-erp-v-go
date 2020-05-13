package koneksi

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDb() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/db_name")

	if err != nil {
		return nil, err
	}
	return db, nil
}
