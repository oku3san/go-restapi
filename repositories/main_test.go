package repositories_test

import (
  "database/sql"
  "fmt"
  _ "github.com/go-sql-driver/mysql"
  "os"
  "testing"
)

var testDB *sql.DB

func setup() error {
  dbUser := "root"
  dbPassword := "pass"
  dbDatabase := "sampledb"
  dbConn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)

  var err error
  testDB, err = sql.Open("mysql", dbConn)
  if err != nil {
    return err
  }
  return nil
}

func teardown() {
  testDB.Close()
}

func TestMain(m *testing.M) {
  err := setup()
  if err != nil {
    os.Exit(1)
  }

  m.Run()

  teardown()
}
