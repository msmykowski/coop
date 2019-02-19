package resolvers

import (
	"api/db"
	"api/db/models"
	"github.com/graphql-go/graphql"
	"time"
)

type JobDefinition struct {
	DB *db.Connection
}

func (c *JobDefinition) Create(params graphql.ResolveParams) (interface{}, error) {
	var jobDefinition models.JobDefinition
	jobDefinition.Description = params.Args["description"].(string)
	jobDefinition.ExecuteEvery = params.Args["executeEvery"].(int)

	tx := c.DB.Begin()

	if err := tx.Error; err != nil {
		return nil, err
	}

	if err := tx.Create(&jobDefinition).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	var jobExecution models.JobExecution
	jobExecution.ExecuteAt = params.Args["executeAt"].(time.Time).UTC()
	jobExecution.JobDefinitionID = jobDefinition.ID

	if err := tx.Create(&jobExecution).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	return jobDefinition, nil
}

func (c *JobDefinition) Show(params graphql.ResolveParams) (interface{}, error) {
	var jobDefinition models.JobDefinition
	err := c.DB.Where("id = ?", params.Args["id"].(uint)).First(&jobDefinition)

	if err.Error != nil {
		return nil, err.Error
	} else {
		return jobDefinition, nil
	}
}
