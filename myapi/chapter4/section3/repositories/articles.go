package repositories

import (
	"database/sql"

	"github.com/yucatty/Go-learning/myapi/models"
)

func InsertArtivle(db *sql.DB, article models.Article) (models.Article, error) {
	const sqlstr = `
		insert into articles (title, contents, username, nice, created_at) values
		(?, ?, ?, 0, now());
	`
	var newArticle models.Article
	newArticle.Title, newArticle.Contents, newArticle.UserName = article.Title, article.Contents, article.UserName

	result, err := db.Exec(sqlstr, article.Title, article.Contents, article.UserName)
	if err != nil {
		return newArticle, err
	}

	id, _ := result.LastInsertId()
	newArticle.ID = int(id)

	return newArticle, nil
}

func SelectArticleList(db *sql.DB, page int) ([]models.Article, error) {
	const sqlstr = `
		select article_id, title, contents, username, nice
		from articles
		limit ? offset ?
	`

	articleArray := make([]models.Article, 0)
	rows, err := db.Query(sqlstr, 5, (5 * (page - 1)))
	if err != nil {
		return []models.Article{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var article models.Article
		err := rows.Scan(&article.ID, &article.Title, &article.Contents, &article.UserName, &article.NiceNum)
		if err != nil {
			return nil, err
		} else {
			articleArray = append(articleArray, article)
		}
	}

	return articleArray, nil
}

func SelectArticleDetail(db *sql.DB, articleID int) (models.Article, error) {
	const sqlStr = `
		select *
		from articles
		where article_id = ?;
	`
	row := db.QueryRow(sqlStr, articleID)
	if err := row.Err(); err != nil {
		return models.Article{}, err
	}

	var article models.Article
	var createdTime sql.NullTime

	err := row.Scan(&article.ID, &article.Title, &article.Contents, &article.UserName, &article.NiceNum, &createdTime)
	if err != nil {
		return models.Article{}, err
	}

	if createdTime.Valid {
		article.CreatedAt = createdTime.Time
	}

	return article, nil
}

func UpdateNiceNum(db *sql.DB, articleID int) error {
	const sqlGetNice = `
		select nce
		from articles
		where article_id = ?;
	`
	const sqlUpdateNice = `
		update articles set nice = ? where article_id = ?
	`

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	row := tx.QueryRow(sqlGetNice, articleID)
	if err := row.Err(); err != nil {
		tx.Rollback()
		return err
	}

	var nicenum int
	err = row.Scan(&nicenum)
	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.Exec(sqlUpdateNice, nicenum+1, articleID)
	if err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
