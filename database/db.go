package database

import (
	"fmt"
	"github.com/Augani/lepora/database/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//options for connecting to postgres database with gorm
type DatabaseOptions struct {
	Host string
	Port int
	UserName string
	Password string
	DatabaseName string
	SSLMode string
}



func InitDatabase(options DatabaseOptions) (*gorm.DB, error) {
	dbUrl := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s", options.UserName, options.Password, options.Host, options.Port, options.DatabaseName, options.SSLMode)
	db, err := gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	//automigrate logs model
	db.AutoMigrate(&models.Log{})
	return db, nil
}