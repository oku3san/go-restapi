package controllers_test

import (
  _ "github.com/go-sql-driver/mysql"
  "github.com/oku3san/go-restapi/controllers"
  "github.com/oku3san/go-restapi/controllers/testdata"
  "testing"
)

var aCon *controllers.ArticleController

func TestMain(m *testing.M) {
  ser := testdata.NewServiceMock()
  aCon = controllers.NewArticleController(ser)

  m.Run()
}
