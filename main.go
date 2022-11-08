package main

import (
  "database/sql"
  "fmt"
  "github.com/gorilla/mux"
  "github.com/oku3san/go-restapi/handlers"
  "log"
  "net/http"
)

func main() {

  dbUser := "root"
  dbPassword := "pass"
  dbDatabase := "sampledb"
  dbConn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)
  db, err := sql.Open("mysql", dbConn)
  if err != nil {
    fmt.Println(err)
  }
  defer db.Close()

  if err := db.Ping(); err != nil {
    fmt.Println(err)
  } else {
    fmt.Println("connect to db")
  }

  r := mux.NewRouter()

  r.HandleFunc("/hello", handlers.HelloHandler).Methods(http.MethodGet)
  r.HandleFunc("/article", handlers.PostArticleHandler).Methods(http.MethodPost)
  r.HandleFunc("/article/list", handlers.ArticleListHandler).Methods(http.MethodGet)
  r.HandleFunc("/article/{id:[0-9]+}", handlers.ArticleDetailHandler).Methods(http.MethodGet)
  r.HandleFunc("/article/nice", handlers.PostNiceHandler).Methods(http.MethodPost)
  r.HandleFunc("/comment", handlers.PostCommentHandler).Methods(http.MethodPost)

  log.Println("server start at port 8080")
  err := http.ListenAndServe(":8080", r)
  log.Fatal(err)
}
