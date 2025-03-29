package internal

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDatabase(DSN string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(DSN), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return db, nil
}
