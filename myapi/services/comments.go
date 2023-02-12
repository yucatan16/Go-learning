package services

import (
	"github.com/yucatty/Go-learning/myapi/models"
	"github.com/yucatty/Go-learning/myapi/repositories"
)

func (s *MyAppService) PostCommentService(comment models.Comment) (models.Comment, error) {
	newComment, err := repositories.InsertComment(s.db, comment)
	if err != nil {
		return models.Comment{}, err
	}

	return newComment, nil
}
