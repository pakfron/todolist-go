package server

import (
	"fmt"
	"os"
	"todolist-go/pkg/database"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func MigrationDB() {

	godotenv.Load("../.env")
	dbname := os.Getenv("MYSQL_DBNAME")
	user := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	host := os.Getenv("MYSQL_HOST")

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, dbname)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(database.Todolist{})

	InstanceDB, _ := db.DB()
	_ = InstanceDB.Close()
}

func GetDB() *gorm.DB {
	godotenv.Load("../.env")
	dbname := os.Getenv("MYSQL_DBNAME")
	user := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	host := os.Getenv("MYSQL_HOST")

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, dbname)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return db
}

func CloseDB(db *gorm.DB) error {
	InstanceDB, _ := db.DB()
	return InstanceDB.Close()
}
