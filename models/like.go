package models

type Like struct {
	Like_author  string  `json:"author" db:"nickname"`
	Like_comment 	string  `json:"comment" db:"id_comment"`
}