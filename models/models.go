package models

type User struct {
	Id       uint   `json:"user_id"`
	Name     string `json:"user_name" gorm:"unique"`
	Password []byte `json:"-"`
}

// I can relegate the username to the frontend/functions
// in the backend. However, I don't see the point as each thread
// will only ever have 1 user. For simplicity, we store the user
// name with the thread and comment.

type Thread struct {
	Id          uint   `json:"thread_id"`
	UserId      uint   `json:"user_id"`
	UserName    string `json:"user_name"`
	TagId       uint   `json:"tag_id"`
	Title       string `json:"thread_title"`
	Body        string `json:"thread_body"`
	DateCreated int64  `json:"thread_date"`
	DateUpdated int64  `json:"thread_upd"`
	CommentNo   uint   `json:"thread_cmmt_no"`
}

type Comment struct {
	Id          uint   `json:"cmmt_id"`
	UserId      uint   `json:"user_id"`
	UserName    string `json:"user_name"`
	ThreadId    uint   `json:"thread_id"`
	CommentSeq  uint   `json:"cmmt_seq"`
	Body        string `json:"cmmt_body"`
	DateCreated int64  `json:"cmmt_date"`
	DateUpdated int64  `json:"cmmt_upd"`
}

// tab stores the user bookmarks and ignores, etc.
// tabtype: bookmark = 1 / ignore = 2
type Tab struct {
	UserId   uint `json:"user_id"`
	ThreadId uint `json:"thread_id"`
	TabType  byte `json:"tab_type"`
}
