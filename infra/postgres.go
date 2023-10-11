package infra

import "database/sql"

func ConnectWithDatabase() *sql.DB {
	connStr := "user=admin password=root dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err.Error())
	}
	return db
}
