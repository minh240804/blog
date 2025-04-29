package model

import (

	// "fmt"

	_ "github.com/lib/pq"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// const (
// 	host     = "localhost"
// 	port     = 5432
// 	user     = "postgres"
// 	password = "123"
// 	dbname   = "DemoBlog"
// )

var DB *gorm.DB

func DbConnect() error {
	var err error
	dsn := "host=localhost user=postgres password=123 dbname=BlogFinal port=5432 sslmode=disable TimeZone='UTC+7'"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	return migrateBlog(DB)
}

func migrateBlog(db *gorm.DB) error {
	err := db.AutoMigrate(&User{})
	if err != nil {
		return err
	}
	err = db.AutoMigrate(&Category{})
	if err != nil {
		return err
	}
	err = db.AutoMigrate(&Blog{})
	if err != nil {
		return err
	}
	err = db.AutoMigrate(&Comment{})
	if err != nil {
		return err
	}
	err = db.AutoMigrate(&Report{})
	if err != nil {
		return err
	}
	err = db.AutoMigrate(&Vote{})
	if err != nil {
		return err
	}
	return nil
}
