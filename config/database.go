package config

import (
	"net/url"

	"itmx_test/service/entity"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var Db *gorm.DB

func InitDB() *gorm.DB {
	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Local")
	Db, err := gorm.Open(sqlite.Open("itmx.sqlite"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	Db.AutoMigrate(
		entity.Customer{},
	)

	return Db
}
