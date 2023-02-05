package repositories_test

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/yucatty/Go-learning/myapi/models"
	"github.com/yucatty/Go-learning/myapi/repositories"
)

func TestInsertComment(t *testing.T) {
	comment := models.Comment{
		ArticleID: 1,
		Message:   "hello",
	}

	expectedCommentID := 3

	newComment, err := repositories.InsertComment(testDB, comment)
	if err != nil {
		t.Error(err)
	}

	if expectedCommentID != newComment.CommentID {
		t.Errorf("new comment id is expected %d but got %d\n", expectedCommentID, newComment.CommentID)
	}

	t.Cleanup(func() {
		const sqlStr = `
			delete from comments
			where message = ?
		`
		testDB.Exec(sqlStr, comment.Message)
	})
}

func TestSelectCommentList(t *testing.T) {
	articleID := 1

	got, err := repositories.SelectCommentList(testDB, articleID)
	if err != nil {
		t.Error(err)
	}

	for _, comment := range got {
		if comment.ArticleID != articleID {
			t.Errorf("want comment of articleID %d but got ID %d\n", articleID, comment.ArticleID)
		}
	}
}
