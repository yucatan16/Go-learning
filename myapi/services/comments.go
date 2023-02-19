package services

import (
	"github.com/yucatty/Go-learning/myapi/apperrors"
	"github.com/yucatty/Go-learning/myapi/models"
	"github.com/yucatty/Go-learning/myapi/repositories"
)

func (s *MyAppService) PostCommentService(comment models.Comment) (models.Comment, error) {
	newComment, err := repositories.InsertComment(s.db, comment)
	if err != nil {
		err = apperrors.InsertDataFailed.Wrap(err, "fail to record data")
		return models.Comment{}, err
	}

	return newComment, nil
}
