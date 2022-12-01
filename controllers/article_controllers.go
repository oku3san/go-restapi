package controllers

import (
  "encoding/json"
  "github.com/gorilla/mux"
  "github.com/oku3san/go-restapi/apperrors"
  "github.com/oku3san/go-restapi/controllers/services"
  "github.com/oku3san/go-restapi/models"
  "io"
  "net/http"
  "strconv"
)

type ArticleController struct {
  service services.ArticleServicer
}

func NewArticleController(s services.ArticleServicer) *ArticleController {
  return &ArticleController{service: s}
}

func (c *ArticleController) HelloHandler(w http.ResponseWriter, req *http.Request) {
  io.WriteString(w, "Hello, world!\n")
}

func (c *ArticleController) PostArticleHandler(w http.ResponseWriter, req *http.Request) {
  var reqArticle models.Article
  if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
    err = apperrors.ReqBodyDecodeFailed.Wrap(err, "bad request body")
    apperrors.ErrorHandler(w, req, err)
    return
  }

  article, err := c.service.PostArticleService(reqArticle)
  if err != nil {
    apperrors.ErrorHandler(w, req, err)
    return
  }
  json.NewEncoder(w).Encode(article)
}

func (c *ArticleController) ArticleListHandler(w http.ResponseWriter, req *http.Request) {
  queryMap := req.URL.Query()

  var page int
  if p, ok := queryMap["page"]; ok && len(p) > 0 {
    var err error
    page, err = strconv.Atoi(p[0])
    if err != nil {
      err = apperrors.BadParam.Wrap(err, "queryparam must be number")
      http.Error(w, "Invalid query parameter", http.StatusBadRequest)
      return
    }
  } else {
    page = 1
  }

  articleList, err := c.service.GetArticleListService(page)
  if err != nil {
    http.Error(w, "Invalid query parameter", http.StatusBadRequest)
    return
  }
  json.NewEncoder(w).Encode(articleList)
}

func (c *ArticleController) ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
  articleID, err := strconv.Atoi(mux.Vars(req)["id"])
  if err != nil {
    err = apperrors.BadParam.Wrap(err, "queryparam must be number")
    http.Error(w, "Invalid query parameter", http.StatusBadRequest)
    return
  }

  article, err := c.service.GetArticleService(articleID)
  if err != nil {
    http.Error(w, "fail internal exec\n", http.StatusInternalServerError)
    return
  }
  json.NewEncoder(w).Encode(article)
}

func (c *ArticleController) PostNiceHandler(w http.ResponseWriter, req *http.Request) {
  var reqArticle models.Article
  if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
    err = apperrors.ReqBodyDecodeFailed.Wrap(err, "bad request body")
    http.Error(w, "fail to decode json\n", http.StatusBadRequest)
    return
  }

  article, err := c.service.PostNiceService(reqArticle)
  if err != nil {
    http.Error(w, "fail internal exec\n", http.StatusInternalServerError)
    return
  }
  json.NewEncoder(w).Encode(article)
}
