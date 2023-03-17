package model

import(
	_ "github.com/lib/pq"
	"log"
	"fmt"
	"database/sql"
)

const (
    host     = "localhost"
    port     = 5432
    user     = "postgres"
    password = "1543137"
    dbname   = "GoConnectDB"
)

var con *sql.DB;

func Connect() *sql.DB{
dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",host, port, user, password, dbname)
	db, err := sql.Open("postgres", dsn)
if err != nil {
    log.Fatal(err)
}
fmt.Println("Database Connected Succesfully")
con = db
return db
}

