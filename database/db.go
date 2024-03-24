package database

import (
	models "Golang-FGA-FinalProject/model"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB(driver string) (db *gorm.DB, err error) {
	if driver == "mysql" {
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", GoDotEnvVariable("DB_USER"), GoDotEnvVariable("DB_PASS"), GoDotEnvVariable("DB_HOST"), GoDotEnvVariable("DB_PORT"), GoDotEnvVariable("DB_NAME"), GoDotEnvVariable("TIMEOUT"))
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			panic("fail to connect to database, error=" + err.Error())
		}
	} else if driver == "postgres" {
		dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", GoDotEnvVariable("DB_HOST"), GoDotEnvVariable("DB_PORT"), GoDotEnvVariable("DB_USER"), GoDotEnvVariable("DB_PASS"), GoDotEnvVariable("DB_NAME"))
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("driver %s not supported", driver)
	}

	log.Default().Println("DB connected successfully")
	db.AutoMigrate(&models.User{}, &models.SocialMedia{}, &models.Photo{}, &models.Comment{})
	return db, nil
}
