package models
type Post struct {
	Post_author  string  `json:"author" db:"post_author"`
	Post_thread  string   `json:"thread" db:"post_thread"`
	Post_created string  `json:"created" db:"post_created"`
	Post_title   string  `json:"title" db:"post_title"`
	Post_content string  `json:"content" db:"post_topic"`
	Post_id      string  `json:"id" db:"post_id"`
}

type FullPost struct {
	Post_author  string  `json:"author" db:"post_author"`
	Post_thread  string   `json:"thread" db:"post_thread"`
	Post_created string  `json:"created" db:"post_created"`
	Post_title   string  `json:"title" db:"post_title"`
	Post_content string  `json:"content" db:"post_topic"`
	Post_id      string  `json:"id" db:"post_id"`
	Comment_author  string  `json:"author" db:"comment_author"`
	Comment_post 	string  `json:"post" db:"comment_post"`
	Comment_created string  `json:"created" db:"comment_created"`
	Comment_content string  `json:"content" db:"comment_content"`
	Likes 			string  `json:"likes" db:"likes"`
	Comment_id      string  `json:"id" db:"id_comment"`
}
