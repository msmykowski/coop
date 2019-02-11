package resolvers_test

import (
	"api/db"
	"api/db/models"
	"api/graphql/resolvers"
	"github.com/graphql-go/graphql"
	"testing"
)

func TestJobDefinitionCreate(t *testing.T) {
	jd := resolvers.JobDefinition{DB: db.Open()}
	m := map[string]interface{}{"description": "description"}

	params := graphql.ResolveParams{Args: m}
	jd.Create(params)

	var jobDefinition models.JobDefinition
	jd.DB.Last(&jobDefinition)

	if jobDefinition.Description != m["description"] {
		t.Error(
      "JobDefinition Create:",
      "expected:",
      "description =",
      m["description"],
      "|",
      "received:",
      "description =",
      jobDefinition.Description,
    )
	}
}
