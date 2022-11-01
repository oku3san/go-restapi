package main

import (
  "github.com/gorilla/mux"
  "github.com/oku3san/go-restapi/handlers"
  "log"
  "net/http"
)

func main() {

  r := mux.NewRouter()

  r.HandleFunc("/hello", handlers.HelloHandler)
  r.HandleFunc("/article", handlers.PostArticleHandler)
  r.HandleFunc("/article/list", handlers.ArticleListHandler)
  r.HandleFunc("/article1", handlers.ArticleDetailHandler)
  r.HandleFunc("/article/nice", handlers.PostNiceHandler)
  r.HandleFunc("/comment", handlers.PostCommentHandler)

  log.Println("server start at port 8080")
  err := http.ListenAndServe(":8080", r)
  log.Fatal(err)
}
