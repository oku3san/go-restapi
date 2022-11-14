package repositories_test

import (
  "database/sql"
  "fmt"
  _ "github.com/go-sql-driver/mysql"
  "github.com/oku3san/go-restapi/models"
  "github.com/oku3san/go-restapi/repositories"
  "testing"
)

func TestSelectArticleDetail(t *testing.T) {

  dbUser := "root"
  dbPassword := "pass"
  dbDatabase := "sampledb"
  dbConn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)

  db, err := sql.Open("mysql", dbConn)
  if err != nil {
    t.Fatal(err)
  }
  defer db.Close()

  expected := models.Article{
    ID:       1,
    Title:    "firstPost",
    Contents: "This is my first blog",
    UserName: "saki",
    NiceNum:  3,
  }

  got, err := repositories.SelectArticleDetail(db, expected.ID)
  if err != nil {
    t.Fatal(err)
  }

  if got.ID != expected.ID {
    t.Errorf("get %d but want %d\n", got.ID, expected.ID)
  }

  if got.Title != expected.Title {
    t.Errorf("get %s but want %s\n", got.Title, expected.Title)
  }

  if got.Contents != expected.Contents {
    t.Errorf("get %s but want %s\n", got.Contents, expected.Contents)
  }

  if got.UserName != expected.UserName {
    t.Errorf("get %s but want %s\n", got.UserName, expected.UserName)
  }

  if got.NiceNum != expected.NiceNum {
    t.Errorf("get %d but want %d\n", got.NiceNum, expected.NiceNum)
  }
}
