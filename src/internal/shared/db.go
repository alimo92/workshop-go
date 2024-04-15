package shared

import (
	"fmt"
	"workshop/src/internal/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"log"
)

var DB *gorm.DB

func InitRds() {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		model.Env.PostgresHost,
		model.Env.PostgresUser,
		model.Env.PostgresPassword,
		model.Env.PostgresDb,
		model.Env.PostgresPort,
		model.Env.PostgresSslMode)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("Failed to connect to db %s", err.Error())
	} else {
		log.Println("Connection to dabase succeeded")
	}

	err = db.AutoMigrate(&model.User{}, &model.YaraRule{})

	if err != nil {
		log.Fatalf("AutoMigrate failed %s", err.Error())
	}

	DB = db
}
