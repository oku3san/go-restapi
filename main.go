package main

import (
  "database/sql"
  "fmt"
  _ "github.com/go-sql-driver/mysql"
  "github.com/gorilla/mux"
  "github.com/oku3san/go-restapi/handlers"
  "github.com/oku3san/go-restapi/models"
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

  const sqlStr = `
    select title, contents, username, nice
    from articles;
  `
  rows, err := db.Query(sqlStr)
  if err != nil {
    fmt.Println(err)
    return
  }
  defer rows.Close()

  articleArray := make([]models.Article, 0)
  for rows.Next() {
    var article models.Article
    err := rows.Scan(&article.Title, &article.Contents, &article.UserName, &article.NiceNum)

    if err != nil {
      fmt.Println(err)
    } else {
      articleArray = append(articleArray, article)
    }
  }

  fmt.Printf("%+v\n", articleArray)

  r := mux.NewRouter()

  r.HandleFunc("/hello", handlers.HelloHandler).Methods(http.MethodGet)
  r.HandleFunc("/article", handlers.PostArticleHandler).Methods(http.MethodPost)
  r.HandleFunc("/article/list", handlers.ArticleListHandler).Methods(http.MethodGet)
  r.HandleFunc("/article/{id:[0-9]+}", handlers.ArticleDetailHandler).Methods(http.MethodGet)
  r.HandleFunc("/article/nice", handlers.PostNiceHandler).Methods(http.MethodPost)
  r.HandleFunc("/comment", handlers.PostCommentHandler).Methods(http.MethodPost)

  log.Println("server start at port 8080")
  err2 := http.ListenAndServe(":8080", r)
  log.Fatal(err2)
}
