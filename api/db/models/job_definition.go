package models

import (
	"github.com/jinzhu/gorm"
)

type JobDefinition struct {
	gorm.Model
  Description string
	ExecuteAt int
	ExecuteEvery int `sql:"type:bigint;"`
}

func (jobDefinition *JobDefinition) Create(db *gorm.DB) error {
	tx := db.Begin()

	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Create(jobDefinition).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
