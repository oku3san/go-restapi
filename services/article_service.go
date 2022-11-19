package services

import (
  "github.com/oku3san/go-restapi/models"
  "github.com/oku3san/go-restapi/repositories"
)

func GetArticleService(articleID int) (models.Article, error) {
  db, err := connectDB()
  if err != nil {
    return models.Article{}, err
  }
  
  article, err := repositories.SelectArticleDetail(db, articleID)
  if err != nil {
    return models.Article{}, err
  }
  commentList, err := repositories.SelectCommentList(db, articleID)
  if err != nil {
    return models.Article{}, err
  }

  article.CommentList = append(article.CommentList, commentList...)

  return article, nil
}