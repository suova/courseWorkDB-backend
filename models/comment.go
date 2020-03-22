package models
type Comment struct {
	Comment_author  string  `json:"author" db:"comment_author"`
	Comment_post 	string  `json:"post" db:"comment_post"`
	Comment_created string  `json:"created" db:"comment_created"`
	Comment_content string  `json:"content" db:"comment_content"`
	Likes 			string  `json:"likes" db:"likes"`
	Comment_id      string  `json:"id" db:"id_comment"`
}