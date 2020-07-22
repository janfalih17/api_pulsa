package configs

import (
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
)

func Connect() (*gorm.DB, error) {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	setting := os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") + "@/" + os.Getenv("DB_NAME") + "?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(os.Getenv("DB_DRIVER"), setting)
	if err != nil {
		return nil, err
	}
	return db, nil
}
