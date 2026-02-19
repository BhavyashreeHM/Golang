package db

import (
    "database/sql"
    "fmt"
    "log"

    _ "github.com/go-sql-driver/mysql"
)

func main() {
    // Replace with your own credentials
    dsn := "root:dreamfight@tcp(127.0.0.1:3306)/"

    // Open connection
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        log.Fatal("Error opening database:", err)
    }
    defer db.Close()

    // Test connection
    err = db.Ping()
    if err != nil {
        log.Fatal("Error connecting to database:", err)
    }

    fmt.Println("Successfully connected to MySQL!")
}