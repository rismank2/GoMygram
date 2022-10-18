package main

import (
	router "MyGram/routers"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "273161"
	dbname   = "mygram"
)

var (
	db *sql.DB

	err error
)

func main() {
	db, err = sql.Open("postgres", ConnectDb(host, user, password, dbname, port))
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Sukses Koneksi Ke Database")
	router.RunRoute(db)

}
func ConnectDb(host, user, password, name string, port int) string {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host,
		port,
		user,
		password,
		dbname)
	return psqlInfo
}
