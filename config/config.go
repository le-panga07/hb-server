package config

import (
	"database/sql"

	_ "hb-server/github.com/go-sql-driver/mysql"
)

//GetMySQLDB func
func GetMySQLDB() (db *sql.DB, err error) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "@Shubham07"
	dbName := "mediaAds"
	db, err = sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	return
}
