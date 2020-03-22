package queries

import (
	"database/sql"
	"databasework/models"
	"log"
)

func CreateComment(comment *models.Comment)  {
	_, err := db.Exec(`INSERT INTO comment VALUES ($1, $2, $3, $4)`,
		comment.Comment_author, comment.Comment_post, comment.Comment_content, comment.Comment_created)
	if err != nil {
		log.Println(err)
	}else{
		log.Println("Ok!")
	}
}

func QueriesGetAllComment(PostID string ) *[]models.Comment {
	comment := &[]models.Comment{}
	err := db.Select(
		comment,
		`SELECT  *
			   FROM comment
			   where comment_post = $1`,
		PostID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		log.Println(err)
		return nil
	}

	return comment
}
