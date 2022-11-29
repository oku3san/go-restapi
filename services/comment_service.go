package services

import (
  "github.com/oku3san/go-restapi/apperrors"
  "github.com/oku3san/go-restapi/models"
  "github.com/oku3san/go-restapi/repositories"
)

func (s *MyAppService) PostCommentService(comment models.Comment) (models.Comment, error) {
  newComment, err := repositories.InsertComment(s.db, comment)
  if err != nil {
    err = apperrors.InsertDataFailed.Wrap(err, "fail to record data")
    return models.Comment{}, err
  }
  return newComment, nil
}
