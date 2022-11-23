package main

import (
  "database/sql"
  "fmt"
  _ "github.com/go-sql-driver/mysql"
  "github.com/gorilla/mux"
  "github.com/oku3san/go-restapi/controllers"
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
  }
  defer db.Close()

  ser := services.NewMyAppService(db)
  con := controllers.NewMyAppController(ser)

  r := mux.NewRouter()

  r.HandleFunc("/hello", con.HelloHandler).Methods(http.MethodGet)
  r.HandleFunc("/article", con.PostArticleHandler).Methods(http.MethodPost)
  r.HandleFunc("/article/list", con.ArticleListHandler).Methods(http.MethodGet)
  r.HandleFunc("/article/{id:[0-9]+}", con.ArticleDetailHandler).Methods(http.MethodGet)
  r.HandleFunc("/article/nice", con.PostNiceHandler).Methods(http.MethodPost)
  r.HandleFunc("/comment", con.PostCommentHandler).Methods(http.MethodPost)

  log.Println("server start at port 8080")
  err2 := http.ListenAndServe(":8080", r)
  log.Fatal(err2)
}
