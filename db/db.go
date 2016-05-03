package db

import (
    "database/sql"
    _ "github.com/lib/pq"
    "log"
)

var db *sql.DB
var err error

func Connect() {
    db, err = sql.Open("postgres", "user=pqgotest dbname=pqgotest sslmode=verify-full")
    if err != nil {
        log.Fatal(err)
    }

    _, err = db.Exec("create table if not exists users(email text)")
    if err != nil {
        log.Fatal(err)
    }
}
