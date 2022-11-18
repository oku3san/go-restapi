package testdata

import "github.com/oku3san/go-restapi/models"

var ArticleTestData = []models.Article{
  models.Article{
    ID:       1,
    Title:    "firstPost",
    Contents: "This is my first blog",
    UserName: "saki",
    NiceNum:  2,
  },
  {
    ID:       2,
    Title:    "2nd",
    Contents: "second blog post",
    UserName: "saki",
    NiceNum:  4,
  },
}
