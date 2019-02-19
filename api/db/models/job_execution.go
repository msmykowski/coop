package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type JobExecution struct {
	gorm.Model
  ExecuteAt time.Time `sql:"type:timestamp without time zone"`
	JobDefinitionID  uint
}
