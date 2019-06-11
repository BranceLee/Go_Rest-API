package models

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"github.com/jinzhu/gorm"
	"os"
	"fmt"
)

var db *gorm.DB

func init() {
	// 将.env 文件读取，并放到环境中
	e := godotenv.Load()
	if e != nil {
		fmt.Print(e)
	}

	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")

	dbURrl := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s",dbHost, username, dbName, password)
	fmt.Println(dbURrl)

	conn, err := gorm.Open("postgres",dbURrl)
	if err != nil {
		fmt.Print(err)
	}

	db = conn

	// create the table in the databa
	db.Debug().AutoMigrate(&Account{}) // Database migration
	// defer db.Close()
}

func GetDB() *gorm.DB{
	return db
}