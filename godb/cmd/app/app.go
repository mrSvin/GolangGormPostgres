package app

import (
	"fmt"
	"gorm.io/gorm"
	"gormPostgres/godb/internal/models"
)

func Start(db *gorm.DB) {
	create(db, "Sava", "65488548")
	create(db, "Kesha", "23432432")
	//selectAndUpdate(db)
}
func create(db *gorm.DB, username string, numberCard string) {
	u := &models.User{
		Name:     username,
		Age:      30,
		IsVerify: true,
		Cards: []models.Card{models.Card{
			Number: numberCard,
			Type:   "VISA",
		}},
	}
	result := db.Create(u)
	if result.Error != nil {
		panic(result.Error)
	}

	fmt.Printf("User created with ID: %d", u.ID)
}

func selectAndUpdate(db *gorm.DB) {
	db.Model(&models.User{}).Where("name = ?", "Alex").Update("name", "Sava")

}
