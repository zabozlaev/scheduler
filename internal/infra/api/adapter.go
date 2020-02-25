package api

import (
	"github.com/graphql-go/graphql"
	"scheduler/internal/domain"
)

type reqBody struct {
	Query string `json:"query"`
}


type Root struct {
	Query *graphql.Object
}

type Adapter struct {
	GqlSchema *graphql.Schema
}

func NewAdapter(service domain.Service) *Adapter {
	resolver := Resolver{service:service}

	// Create a new Root that describes our base query set up. In this
	// example we have a user query that takes one argument called name
	root := Root{
		Query: graphql.NewObject(
			graphql.ObjectConfig{
				Name: "Query",
				Fields: graphql.Fields{
					"classes": &graphql.Field{
						Description: "Classes for certain group and a day",
						Type: graphql.NewList(graphql.NewNonNull(ClassType)),
						Args: graphql.FieldConfigArgument{
							"group": &graphql.ArgumentConfig{
								Type: graphql.NewNonNull(graphql.String),
								Description: "Group number",
							},
							"day": &graphql.ArgumentConfig{
								Type:         graphql.NewNonNull(DayEnumType),
								Description: "Day",
							},
						},
						Resolve: resolver.ClassesResolver,
					},
					"schedule": &graphql.Field{
						Type:              ScheduleType,
						Args:              graphql.FieldConfigArgument{
							"group": &graphql.ArgumentConfig{
								Type:         graphql.NewNonNull(graphql.String),
								Description:  "Group number",
							},
						},
						Resolve:           resolver.ScheduleResolver,
						Description:       "Schedule for certain group",
					},
				},
			},
		),
	}
	sc, err := graphql.NewSchema(
		graphql.SchemaConfig{Query: root.Query},
	)

	if err != nil {
		panic(err)
	}

	return &Adapter{GqlSchema: &sc}
}