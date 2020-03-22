package queries

import (
	"databasework/models"
	"log"
)

func Like(like *models.Like)  {
		_, err := db.Exec(`INSERT INTO likes VALUES ($1, $2)`,
			like.Like_author, like.Like_comment)
		if err != nil {
			log.Println(err)
		}else{
			log.Println("Ok!")
		}

}

func DisLike(like *models.Like)  {
	_, err := db.Exec(`DELETE 
							FROM likes
							WHERE nickname = $1 and id_comment= $2`,
		like.Like_author, like.Like_comment)
	if err != nil {
		log.Println(err)
	}else{
		log.Println("Ok!")
	}

}
