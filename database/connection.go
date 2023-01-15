package database

import (
	"github.com/WZWren/gossip-app-backend/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	connection, err := gorm.Open(mysql.Open("root:rootwalla@/gossip"), &gorm.Config{})

	if err != nil {
		panic("could not connect to the database")
	}

	DB = connection

	connection.AutoMigrate(&models.User{})
	connection.AutoMigrate(&models.Thread{})
	connection.AutoMigrate(&models.Comment{})
}
