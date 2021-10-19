package config

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"os"
)
var db *gorm.DB
var err error

func GetConn() *gorm.DB{
	// Please define your username and password for MySQL.
	host := os.Getenv("DATABASE_HOST")
	port := os.Getenv("DATABASE_PORT")
	user := os.Getenv("USERNAME_DB")
	pass := os.Getenv("PASSWORD_DB")
	dbname := os.Getenv("DATABASE_NAME")

	conn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", user, pass, host, port, dbname)
	log.Print(conn)
	db, err = gorm.Open("mysql", conn)
	// NOTE: See we’re using = to assign the global var
	// instead of := which would assign it only in this function

	if err != nil{
		log.Print("Connection Failed to Open")
	}else{
		log.Print("Connection Established")
	}

	return db
}