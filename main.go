package main

import (
  "github.com/oku3san/go-restapi/handlers"
  "log"
  "net/http"
)

func main() {

  http.HandleFunc("/hello", handlers.HelloHandler)
  http.HandleFunc("/article", handlers.PostArticleHandler)
  http.HandleFunc("/article/list", handlers.ArticleListHandler)
  http.HandleFunc("/article1", handlers.ArticleDetailHandler)
  http.HandleFunc("/article/nice", handlers.PostNiceHandler)
  http.HandleFunc("/comment", handlers.PostCommentHandler)

  log.Println("server start at port 8080")
  err := http.ListenAndServe(":8080", nil)
  log.Fatal(err)
}
