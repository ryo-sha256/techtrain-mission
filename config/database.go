// config/database.go
// db configuration
package Config

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

var DB_CONN *gorm.DB

// Builder Pattern
// TODO: Singleton
type DBConfig struct {
	userName string
	dbName   string
	password string
	host     string
	port     string
}

func DBConfigBuilder() *DBConfig {
	// can be replaced with the std library one
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	userName := os.Getenv("DB_USER_NAME")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")

	dbConfig := DBConfig{
		userName: userName,
		password: password,
		dbName:   dbName,
		host:     host,
		port:     port,
	}

	return &dbConfig
}

func DBConfigURL(dbConfig *DBConfig) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.userName,
		dbConfig.password,
		dbConfig.host,
		dbConfig.port,
		dbConfig.dbName,
	)
}
