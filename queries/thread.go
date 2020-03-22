package queries

import (
	"database/sql"
	"databasework/models"

	"log"
)


func QueriesGetAllThreads() *[]models.Thread {
	thread := &[]models.Thread{}
	err := db.Select(
		thread,
		`SELECT  *
			   FROM thread
			   `)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}
		log.Println(err)
		return nil
	}

	return thread
}

func CreateThread(thread *models.Thread)  {
	_, err := db.Exec(`INSERT INTO thread VALUES ($1, $2)`,
		thread.Thread_created, thread.Thread_title)
	if err != nil {
		log.Println(err)
	}else{
		log.Println("Ok!")
	}

}
