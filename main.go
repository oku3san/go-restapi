package main

import (
  "database/sql"
  "fmt"
  _ "github.com/go-sql-driver/mysql"
  "github.com/oku3san/go-restapi/controllers"
  "github.com/oku3san/go-restapi/routers"
  "github.com/oku3san/go-restapi/services"
  "log"
  "net/http"
  "os"
)

var (
  dbUser     = os.Getenv("DBUSERNAME")
  dbPassword = os.Getenv("DBUSERPASS")
  dbDatabase = os.Getenv("DATABASE")
  dbConn     = fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)
)

func main() {
  db, err := sql.Open("mysql", dbConn)
  if err != nil {
    fmt.Println("fail to connect DB")
    return
  }
  //defer db.Close()

  ser := services.NewMyAppService(db)
  con := controllers.NewMyAppController(ser)
  r := routers.NewRouter(con)

  log.Println("server start at port 8080")
  log.Fatal(http.ListenAndServe(":8080", r))
}
