package resolvers

import (
    "api/db"
    "api/db/models"
    "github.com/graphql-go/graphql"
)

type JobDefinition struct {
  db *db.Connection
}

func (c *JobDefinition) Show(params graphql.ResolveParams) (interface{}, error) {
    var jobDefinition models.JobDefinition
    err := c.db.Where("id = ?", params.Args["id"].(string)).First(&jobDefinition)

    if err.Error != nil {
      return nil, err.Error
    } else {
        return jobDefinition, nil
    }
}

func (c *JobDefinition) Create(params graphql.ResolveParams) (interface{}, error) {
  var jobDefinition models.JobDefinition
  jobDefinition.Description = params.Args["description"].(string)
  err := c.db.Create(&jobDefinition)

  if err.Error != nil {
    return nil, err.Error
  } else {
      return jobDefinition, nil
  }
}
