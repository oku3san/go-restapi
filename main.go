package main

import (
  "github.com/gorilla/mux"
  "github.com/oku3san/go-restapi/handlers"
  "log"
  "net/http"
)

func main() {

  r := mux.NewRouter()

  r.HandleFunc("/hello", handlers.HelloHandler).Methods(http.MethodGet)
  r.HandleFunc("/article", handlers.PostArticleHandler).Methods(http.MethodPost)
  r.HandleFunc("/article/list", handlers.ArticleListHandler).Methods(http.MethodGet)
  r.HandleFunc("/article1", handlers.ArticleDetailHandler).Methods(http.MethodGet)
  r.HandleFunc("/article/nice", handlers.PostNiceHandler).Methods(http.MethodPost)
  r.HandleFunc("/comment", handlers.PostCommentHandler).Methods(http.MethodPost)

  log.Println("server start at port 8080")
  err := http.ListenAndServe(":8080", r)
  log.Fatal(err)
}
