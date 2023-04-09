package mygorm

import (
	"github.com/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New(url string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	err = migrate(db)
	if err != nil {
		return nil, errors.Wrap(err, "failed to migrate")
	}
	return db, nil
}

func migrate(db *gorm.DB) error {
	return db.AutoMigrate(&User{})
}
