package queries

import (
	"database/sql"
	"databasework/models"
	"log"
)

func CreatePost(post *models.Post)  {
	_, err := db.Exec(`INSERT INTO post VALUES ($1, $2, $3, $4, $5)`,
		post.Post_author, post.Post_thread, post.Post_created, post.Post_title, post.Post_content)
	if err != nil {
		log.Println(err)
	}else{
		log.Println("Ok!")
	}

}

func QueriesGetAllPost(ForumID string ) *[]models.Post {
	post := &[]models.Post{}
	err := db.Select(
		post,
		`SELECT  *
			   FROM post
			   where post_thread = $1`,
			   ForumID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		log.Println(err)
		return nil
	}

	return post
}

func QueriesGetOnePost(PostID string ) *[]models.Post {
	post := &[]models.Post{}
	err := db.Select(
		post,
		`SELECT  *
			   FROM post
			   where post_id = $1`,
		PostID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		log.Println(err)
		return nil
	}

	return post
}

func QueriesDeleteOnePost(PostID string ) bool {
	_, err := db.Exec(
		`DELETE 
			   FROM post
			   WHERE post_id = $1`,
		PostID)
	if err != nil {
		if err == sql.ErrNoRows {
			return false
		}
		log.Println(err)
		return false
	}

	return true
}

