package model

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/vishals11/gouser/config"
)

var db *gorm.DB

func init() {
	var err error
	cfg := config.Get()
	connectionString := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", cfg.DBUserName, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)

	db, err = gorm.Open("mysql", connectionString)
	if err != nil {
		log.Fatalf("Failed connecting to database: %s", err)
	}

	log.Println("Connecting to database: successful")
	db.LogMode(true)
}
