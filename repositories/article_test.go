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

  tests := []struct {
    testTitle string
    expected  models.Article
  }{
    {
      testTitle: "subtest1",
      expected: models.Article{
        ID:       1,
        Title:    "firstPost",
        Contents: "This is my first blog",
        UserName: "saki",
        NiceNum:  3,
      },
    }, {
      testTitle: "subtest2",
      expected: models.Article{
        ID:       2,
        Title:    "2nd",
        Contents: "second blog post",
        UserName: "saki",
        NiceNum:  4,
      },
    },
  }

  for _, test := range tests {
    t.Run(test.testTitle, func(t *testing.T) {
      got, err := repositories.SelectArticleDetail(db, test.expected.ID)
      if err != nil {
        t.Fatal(err)
      }

      if got.ID != test.expected.ID {
        t.Errorf("get %d but want %d\n", got.ID, test.expected.ID)
      }

      if got.Title != test.expected.Title {
        t.Errorf("get %s but want %s\n", got.Title, test.expected.Title)
      }

      if got.Contents != test.expected.Contents {
        t.Errorf("get %s but want %s\n", got.Contents, test.expected.Contents)
      }

      if got.UserName != test.expected.UserName {
        t.Errorf("get %s but want %s\n", got.UserName, test.expected.UserName)
      }

      if got.NiceNum != test.expected.NiceNum {
        t.Errorf("get %d but want %d\n", got.NiceNum, test.expected.NiceNum)
      }
    })
  }
}