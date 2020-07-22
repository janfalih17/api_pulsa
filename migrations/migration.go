package migrations

import (
	"log"

	"github.com/janfalih17/api_pulsa/models"

	"github.com/janfalih17/api_pulsa/configs"
)

func Migrate() {
	db, err := configs.Connect()
	if err != nil {
		panic(err)
	}
	if db.HasTable(&models.User{}) && db.HasTable(&models.Status{}) != false {
		log.Fatal("Error Database Has Exist")
		return
	}
	db.AutoMigrate(&models.User{}, &models.Status{})
	db.Model(&models.User{}).AddForeignKey("status_id", "statuses(id)", "RESTRICT", "RESTRICT")
}
