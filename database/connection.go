package database

import (
	"github.com/WZWren/gossip-app-backend/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	// this line is where the SQL root account field goes - change to connect to a different db
	connection, err := gorm.Open(mysql.Open("root:rootwalla@/gossip"), &gorm.Config{})

	if err != nil {
		panic("could not connect to the database")
	}

	DB = connection

	// this syncs the golang instance with the mysql server.
	connection.AutoMigrate(&models.User{})
	connection.AutoMigrate(&models.Thread{})
	connection.AutoMigrate(&models.Comment{})
	connection.AutoMigrate(&models.Tab{})
}
