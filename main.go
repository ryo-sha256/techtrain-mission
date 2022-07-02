package main

import (
	"fmt"

	config "techtrain-mission/config"
	models "techtrain-mission/models"
	routers "techtrain-mission/routers"

	_ "github.com/go-sql-driver/mysql"

	"github.com/jinzhu/gorm"
)

var err error

func main() {
	config.DB_CONN, err = gorm.Open("mysql", config.DBConfigURL(config.DBConfigBuilder()))
	if err != nil {
		fmt.Println("Status:", err)
	}

	defer config.DB_CONN.Close()

	// User : user table
	// CreateCharacterRequest : character table
	// UserCharacter : table to store the relation of user and character
	// RarerityList : table to store the list of character drop rate
	config.DB_CONN.AutoMigrate(
		&models.User{},
		&models.CreateCharacterRequest{},
		&models.UserCharacter{},
		&models.RarityList{})

	r := routers.RouterBuilder()
	r.Run()
}
