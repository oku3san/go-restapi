package main

import (
  "database/sql"
  "fmt"
  _ "github.com/go-sql-driver/mysql"
  "os"
)

func main() {

  dbUser := os.Getenv("DBUSERNAME")
  dbPassword := os.Getenv("DBUSERPASS")
  dbDatabase := os.Getenv("DATABASE")
  dbConn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)
  db, err := sql.Open("mysql", dbConn)
  if err != nil {
    fmt.Println(err)
  }
  defer db.Close()

  //articleID := 1
  //const sqlStr = `
  //  select *
  //  from articles
  //  where article_id = ?;
  //`
  //row := db.QueryRow(sqlStr, articleID)
  //if err := row.Err(); err != nil {
  //  fmt.Println(err)
  //  return
  //}
  //
  //var article models.Article
  //var createdTime sql.NullTime
  //err = row.Scan(&article.ID, &article.Title, &article.Contents, &article.UserName, &article.NiceNum, &createdTime)
  //
  //if err != nil {
  //  fmt.Println(err)
  //  return
  //}
  //
  //if createdTime.Valid {
  //  article.CreatedAt = createdTime.Time
  //}
  //
  //fmt.Printf("%+v\n", article)
  //
  //articleInsert := models.Article{
  //  Title:    "insert test",
  //  Contents: "Can I insert data correctly?",
  //  UserName: "saki",
  //}
  //const sqlStrInsert = `
  //  insert into articles (title, contents, username, nice, created_at) values
  //  (?,?,?,0,now());
  //`
  //result, err := db.Exec(sqlStrInsert, articleInsert.Title, articleInsert.Contents, articleInsert.UserName)
  //if err != nil {
  //  fmt.Println(err)
  //  return
  //}
  //
  //fmt.Println(result.LastInsertId())
  //fmt.Println(result.RowsAffected())

  tx, err := db.Begin()
  if err != nil {
    fmt.Println(err)
    return
  }

  article_id := 1
  const sqlGetNice = `
    select nice
    from articles
    where article_id = ?;
  `

  row := tx.QueryRow(sqlGetNice, article_id)
  if err := row.Err(); err != nil {
    fmt.Println(err)
    tx.Rollback()
    return
  }

  var nicenum int
  err = row.Scan(&nicenum)
  if err != nil {
    fmt.Println(err)
    tx.Rollback()
    return
  }

  const sqlUpdateNice = `update articles set nice = ? where article_id = ?`
  _, err = tx.Exec(sqlUpdateNice, nicenum+1, article_id)
  if err != nil {
    fmt.Println(err)
    tx.Rollback()
    return
  }

  tx.Commit()

  //r := mux.NewRouter()
  //
  //r.HandleFunc("/hello", handlers.HelloHandler).Methods(http.MethodGet)
  //r.HandleFunc("/article", handlers.PostArticleHandler).Methods(http.MethodPost)
  //r.HandleFunc("/article/list", handlers.ArticleListHandler).Methods(http.MethodGet)
  //r.HandleFunc("/article/{id:[0-9]+}", handlers.ArticleDetailHandler).Methods(http.MethodGet)
  //r.HandleFunc("/article/nice", handlers.PostNiceHandler).Methods(http.MethodPost)
  //r.HandleFunc("/comment", handlers.PostCommentHandler).Methods(http.MethodPost)
  //
  //log.Println("server start at port 8080")
  //err2 := http.ListenAndServe(":8080", r)
  //log.Fatal(err2)
}
