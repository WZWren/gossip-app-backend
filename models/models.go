package models

type User struct {
	Id       uint   `json:"user_id"`
	Name     string `json:"user_name" gorm:"unique"`
	Password []byte `json:"-"`
}

type Thread struct {
	Id          uint   `json:"thread_id"`
	UserId      uint   `json:"user_id"`
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
	ThreadId    uint   `json:"thread_id"`
	CommentSeq  uint   `json:"cmmt_seq"`
	Body        string `json:"cmmt_body"`
	DateCreated int64  `json:"cmmt_date"`
	DateUpdated int64  `json:"cmmt_upd"`
}

// tab stores the user bookmarks and ignores, etc.
// tabtype: bookmark / ignore
type Tab struct {
	UserId   uint   `json:"user_id"`
	ThreadId uint   `json:"thread_id"`
	TabType  string `json:"tab_type"`
}
