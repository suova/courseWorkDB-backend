package queries

import (
	"database/sql"
	"databasework/models"
	"log"
)

func User(user *models.User) string {
	p := &models.Exist{}
	_ = db.Get(p, "SELECT COUNT(*) as count FROM  user_interface WHERE nickname = $1 or email = $2",
		user.Nickname, user.Email)

	if p.Count == "0" {
		println(user.Nickname)
		if user.Role == 3 || user.Role == 2 {
			_, err := db.Exec(`INSERT INTO user_interface VALUES ($1, $2, $3, $4, $5, $6, $7)`,
				user.Nickname, user.Email, user.Fullname, user.About, user.Country, user.Password, user.Role)
			if err != nil {
				log.Println(err)
			}else{
				log.Println("Ok!")
			}
		}else{
			_, err := db.Exec(`INSERT INTO user_interface VALUES ($1, $2, $3, $4, $5, $6)`,
				user.Nickname, user.Email, user.Fullname, user.About, user.Country, user.Password)
			if err != nil {
				log.Println(err)
			}else{
				log.Println("Ok!")
			}
			return "OK"
		}

	}
	return "EXIST"
}

func FindUser(nickname string) *models.User {
	findUser := &models.User{}
	err := db.Get(
		findUser,
		`SELECT *
			   FROM user_interface 
			   WHERE nickname = $1`,
		nickname)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}

		log.Println("Error of find user by nickname")
		return nil
	}

	return findUser
}

func ChangeRole(nickname string, role int) error {

	sqlStatement := `UPDATE user_interface
					SET role = $2
					WHERE nickname = $1;`
	_, err := db.Exec(sqlStatement, nickname, role)

	if err != nil {
		if err == sql.ErrNoRows {
			return err
		}

		log.Println("Error of find user by nickname")
		return err
	}

	return nil
}

func FindUsers() *[]models.User {

	findUsers := &[]models.User{}
	err := db.Select(
		findUsers,
		`SELECT  *
			   FROM user_interface`)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil
		}

		log.Println("Error of find users")
		return nil
	}

	return findUsers
}
