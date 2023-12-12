package repository

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
	"time"
)

var dbase *gorm.DB

func Init() *gorm.DB {
	db, err := gorm.Open("postgres", "host=localhost port=5436 postgres_password=qwerty  sslmode=disable")

	if err != nil {
		log.Fatal(err)
	}

	return db

}

func GetDB() *gorm.DB {
	if dbase == nil {
		dbase = Init()
		var sleep time.Duration
		for dbase == nil {
			sleep = sleep * 2
			fmt.Printf("database if unavivalve,wait a %d sec.\n ", sleep)
			time.Sleep(sleep * time.Second)
			dbase = Init()
		}
	}

	return dbase
}
