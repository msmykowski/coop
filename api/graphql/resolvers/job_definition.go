package resolvers

import (
	"api/db"
	"api/db/models"
	"github.com/graphql-go/graphql"
)

type JobDefinition struct {
	DB *db.Connection
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

func (c *JobDefinition) Create(params graphql.ResolveParams) (interface{}, error) {
	var jobDefinition models.JobDefinition
	jobDefinition.Description = params.Args["description"].(string)
	err := c.DB.Create(&jobDefinition)

	if err.Error != nil {
		return nil, err.Error
	} else {
		return jobDefinition, nil
	}
}
