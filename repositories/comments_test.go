package repositories_test

import (
  _ "github.com/go-sql-driver/mysql"
  "github.com/oku3san/go-restapi/models"
  "github.com/oku3san/go-restapi/repositories"
  "testing"
)

func TestSelectCommentList(t *testing.T) {
  articleID := 1
  got, err := repositories.SelectCommentList(testDB, articleID)
  if err != nil {
    t.Fatal(err)
  }

  for _, comment := range got {
    if comment.ArticleID != articleID {
      t.Errorf("want comment of articleID %d but got ID %d\n", articleID, comment.ArticleID)
    }
  }
}

func TestInsertComment(t *testing.T) {
  comment := models.Comment{
    ArticleID: 1,
    Message:   "CommentInsertTest",
  }

  expectedCommentID := 3
  newComment, err := repositories.InsertComment(testDB, comment)
  if err != nil {
    t.Error(err)
  }
  if newComment.CommentID != expectedCommentID {
    t.Errorf("new comment id is expected %d but got %d\n", expectedCommentID, newComment.CommentID)
  }

  t.Cleanup(func() {
    const sqlStr = `
			delete from comments
			where message = ?
		`
    const sqlStrForReset = `
      alter table comments AUTO_INCREMENT = 1;
    `
    testDB.Exec(sqlStr, comment.Message)
    testDB.Exec(sqlStrForReset)
  })
}
