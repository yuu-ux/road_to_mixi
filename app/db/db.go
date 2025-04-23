package db

import (
    "database/sql"
    "os"
    "fmt"
)
import _ "github.com/go-sql-driver/mysql"

func New() (*sql.DB, error) {
    user := os.Getenv("USER")
    password := os.Getenv("PASSWORD")
    host := os.Getenv("HOST")
    port := os.Getenv("PORT")
    database := os.Getenv("DATABASE")
    uri := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, password, host, port, database)
    db, err := sql.Open("mysql", uri)
    if err != nil {
        return nil, err
    }
    return db, nil
}
