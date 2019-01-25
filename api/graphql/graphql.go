package graphql

import (
	"api/db"
	"github.com/graphql-go/handler"
	"api/graphql/types"
	"api/graphql/resolvers"
	"github.com/graphql-go/graphql"
)

type GraphQL struct {
	Resolvers resolvers.Resolvers
	Types types.Types
}

func New(db *db.Connection) *handler.Handler {
	gql := GraphQL{
		Resolvers: resolvers.New(db),
		Types: types.New(),
	}

	schema, _ := gql.schema()

	return handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})
}

func (gQL *GraphQL) schema() (graphql.Schema, error) {
	queryType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"jobDefinition": &graphql.Field{
				Type: gQL.Types.JobDefinition,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.ID),
					},
				},
				Resolve: gQL.Resolvers.JobDefinition.Show,
			},
		},
	})

	mutationType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"createJobDefinition": &graphql.Field{
				Type: gQL.Types.JobDefinition,
				Args: graphql.FieldConfigArgument{
					"description": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: gQL.Resolvers.JobDefinition.Create,
			},
		},
	})

	schemaConfig := graphql.SchemaConfig{
		Query: queryType,
		Mutation: mutationType,
	}

	return graphql.NewSchema(schemaConfig)
}
