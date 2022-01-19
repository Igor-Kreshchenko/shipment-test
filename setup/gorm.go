package setup

import (
	"github.com/Igor-Kreshchenko/shipment-test/models"
	"gorm.io/gorm"

	"gorm.io/driver/sqlite"
)

func ConnectDataBase() (*gorm.DB, error) {
	var db *gorm.DB

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	err = db.AutoMigrate(&models.Shipment{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
