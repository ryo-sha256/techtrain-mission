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
	// table for the list of active users, all the characters
	// and who owns which character
	config.DB_CONN.AutoMigrate(&models.User{}, &models.CreateCharacterRequest{}, &models.UserCharacter{})

	// router goes here
	r := routers.RouterBuilder()
	r.Run()
}
