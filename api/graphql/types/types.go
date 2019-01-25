package types

import (
  "github.com/graphql-go/graphql"
)

type Types struct {
  JobDefinition *graphql.Object
}

func New() Types {
  return Types{
    JobDefinition: graphql.NewObject(graphql.ObjectConfig{
  		Name: "JobDefinition",
  		Fields: graphql.Fields{
  			"description": &graphql.Field{
  				Type: graphql.String,
  			},
  		},
  	}),
  }
}
