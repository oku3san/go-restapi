package controllers

import (
  "encoding/json"
  "github.com/oku3san/go-restapi/apperrors"
  "github.com/oku3san/go-restapi/controllers/services"
  "github.com/oku3san/go-restapi/models"
  "net/http"
)

type CommentController struct {
  service services.CommentServicer
}

func NewCommentController(s services.CommentServicer) *CommentController {
  return &CommentController{service: s}
}

func (c *CommentController) PostCommentHandler(w http.ResponseWriter, req *http.Request) {
  var reqComment models.Comment
  if err := json.NewDecoder(req.Body).Decode(&reqComment); err != nil {
    err = apperrors.ReqBodyDecodeFailed.Wrap(err, "bad request body")
    http.Error(w, "fail to decode json\n", http.StatusBadRequest)
    return
  }

  comment, err := c.service.PostCommentService(reqComment)
  if err != nil {
    http.Error(w, "fail internal exec\n", http.StatusInternalServerError)
    return
  }
  json.NewEncoder(w).Encode(comment)
}
