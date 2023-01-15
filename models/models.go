package models

type User struct {
	Id       uint
	Name     string `gorm:"unique"`
	Password []byte
}
