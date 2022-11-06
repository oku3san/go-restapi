package handlers

import (
  "encoding/json"
  "errors"
  "github.com/oku3san/go-restapi/models"
  "io"
  "net/http"
  "strconv"
)

func HelloHandler(w http.ResponseWriter, req *http.Request) {
  io.WriteString(w, "Hello, world!\n")
}

func PostArticleHandler(w http.ResponseWriter, req *http.Request) {
  length, err := strconv.Atoi(req.Header.Get("Content-Length"))
  if err != nil {
    http.Error(w, "cannot get content length\n", http.StatusBadRequest)
    return
  }
  reqBodybuffer := make([]byte, length)

  if _, err := req.Body.Read(reqBodybuffer); !errors.Is(err, io.EOF) {
    http.Error(w, "fail to get request body\n", http.StatusBadRequest)
    return
  }
  defer req.Body.Close()

  var reqArticle models.Article
  if err := json.Unmarshal(reqBodybuffer, &reqArticle); err != nil {
    http.Error(w, "fail to decode json\n", http.StatusBadRequest)
    return
  }

  article := reqArticle
  jsonData, err := json.Marshal(article)
  if err != nil {
    http.Error(w, "fail to encode json\n", http.StatusInternalServerError)
    return
  }
  w.Write(jsonData)
}

func ArticleListHandler(w http.ResponseWriter, req *http.Request) {
  //queryMap := req.URL.Query()
  //
  //var page int
  //if p, ok := queryMap["page"]; ok && len(p) > 0 {
  //  var err error
  //  page, err = strconv.Atoi(p[0])
  //  if err != nil {
  //    http.Error(w, "Invalid query parameter", http.StatusBadRequest)
  //    return
  //  }
  //} else {
  //  page = 1
  //}
  //resString := fmt.Sprintf("Article List (page %d)\n", page)
  //io.WriteString(w, resString)
  articleList := []models.Article{models.Article1, models.Article2}
  jsonData, err := json.Marshal(articleList)
  if err != nil {
    http.Error(w, "fail to encode json\n", http.StatusInternalServerError)
    return
  }
  w.Write(jsonData)
}

func ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
  //articleID, err := strconv.Atoi(mux.Vars(req)["id"])
  //if err != nil {
  //  http.Error(w, "Invalid query parameter", http.StatusBadRequest)
  //  return
  //}
  //resString := fmt.Sprintf("Article No.%d\n", articleID)
  //io.WriteString(w, resString)
  article := models.Article1
  jsonData, err := json.Marshal(article)
  if err != nil {
    http.Error(w, "fail to encode json\n", http.StatusInternalServerError)
    return
  }
  w.Write(jsonData)
}

func PostNiceHandler(w http.ResponseWriter, req *http.Request) {
  article := models.Article1
  jsonData, err := json.Marshal(article)
  if err != nil {
    http.Error(w, "fail to encode json\n", http.StatusInternalServerError)
    return
  }
  w.Write(jsonData)
}

func PostCommentHandler(w http.ResponseWriter, req *http.Request) {
  comment := models.Comment1
  jsonData, err := json.Marshal(comment)
  if err != nil {
    http.Error(w, "fail to encode json\n", http.StatusInternalServerError)
    return
  }
  w.Write(jsonData)
}
