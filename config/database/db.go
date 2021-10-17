package database

import (
	"fmt"

	"github.com/itp-backend/backend-a-co-create/config"
	"github.com/itp-backend/backend-a-co-create/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// SetupDBConn SetupDBConnection
func SetupDBConn() *gorm.DB {
	dbuser := config.Init().DBUsername
	dbpassword := config.Init().DBPassword
	dbhost := config.Init().DBHost
	dbport := config.Init().DBPort
	dbname := config.Init().DBName

	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=True&loc=Local", dbuser, dbpassword, dbhost, dbport, dbname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to create a connection to database")
	}

	fmt.Println("Connect to database...")

	// db migrate
	db.AutoMigrate(&model.Role{}, &model.User{}, &model.Enrollment{}, &model.Project{})

	return db
}

// CloseDBConn to close connection database
func CloseDBConn(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to close a connection from database")
	}
	dbSQL.Close()
}
