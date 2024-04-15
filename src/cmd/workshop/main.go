package main

import (
	"log"
	"workshop/src/internal/api"
	"workshop/src/internal/model"
	"workshop/src/internal/shared"

	"gorm.io/gorm/clause"
)

func main() {
	shared.InitRds()
	api.InitRoutes()

	user := model.User{FirstName: "Ali", LastName: "Moro"}

	result := shared.DB.Clauses(clause.Returning{}).Select("*").Create(&user)

	if result.Error != nil {
		log.Fatalf("Failed to create user %s", result.Error)
	}

	result.First(&user)
	log.Println(user)
}
