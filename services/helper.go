package services

import (
  "database/sql"
  "fmt"
  _ "github.com/go-sql-driver/mysql"
  "os"
)

var (
  dbUser     = os.Getenv("DBUSERNAME")
  dbPassword = os.Getenv("DBUSERPASS")
  dbDatabase = os.Getenv("DATABASE")
  dbConn     = fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)
)

func connectDB() (*sql.DB, error) {
  db, err := sql.Open("mysql", dbConn)
  if err != nil {
    return nil, err
  }
  return db, nil
}
