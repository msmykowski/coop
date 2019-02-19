package resolvers_test

import (
	"api/db"
	"api/db/models"
	"api/graphql/resolvers"
	"github.com/graphql-go/graphql"
	"testing"
	"time"
)

func TestJobDefinitionCreate(t *testing.T) {
	jd := resolvers.JobDefinition{DB: db.Open()}
	m := map[string]interface{}{"description": "description", "executeAt": time.Now(), "executeEvery": 1440}

	params := graphql.ResolveParams{Args: m}
	jd.Create(params)

	var jobDefinition models.JobDefinition
	jd.DB.Last(&jobDefinition)

	var jobExecution models.JobExecution
	jd.DB.Last(&jobExecution)

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

	if jobDefinition.ExecuteEvery != m["executeEvery"] {
		t.Error(
			"JobDefinition Create:",
			"expected:",
			"executeEvery =",
			m["executeEvery"],
			"|",
			"received:",
			"executeEvery =",
			jobDefinition.ExecuteEvery,
		)
	}

	if jobExecution.ExecuteAt.UTC() != m["executeAt"].(time.Time).UTC() {
		t.Error(
			"JobDefinition Create:",
			"expected:",
			"executeAt =",
			m["executeAt"].(time.Time).UTC(),
			"|",
			"received:",
			"executeAt =",
			jobExecution.ExecuteAt.UTC(),
		)
	}

	if jobExecution.JobDefinitionID != jobDefinition.ID {
		t.Error(
			"JobDefinition Create:",
			"expected:",
			"JobDefinitionID =",
			jobDefinition.ID,
			"|",
			"received:",
			"executeAt =",
			jobExecution.JobDefinitionID,
		)
	}
}

func TestJobDefinitionShow(t *testing.T) {
	jd := resolvers.JobDefinition{DB: db.Open()}

	expected := models.JobDefinition{Description: "description"}
	jd.DB.Create(&expected)

	m := map[string]interface{}{"id": expected.ID}
	params := graphql.ResolveParams{Args: m}
	actual, err := jd.Show(params)

	if err != nil {
		t.Error("Error:", err)
	}

	description := actual.(models.JobDefinition).Description

	if expected.Description != description {
		t.Error(
			"JobDefinition Create:",
			"expected:",
			"description =",
			description,
			"|",
			"received:",
			"description =",
			expected.Description,
		)
	}
}
