package models

import (
	"github.com/jinzhu/gorm"
)

type JobDefinition struct {
	gorm.Model
  Description string
	ExecuteEvery int `sql:"type:bigint;"`
	JobExecutions []JobExecution
}
