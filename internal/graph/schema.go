package graph

import "github.com/graphql-go/graphql"

var locationType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Location",
	Fields: graphql.Fields{
		"id":        &graphql.Field{Type: graphql.String},
		"latitude":  &graphql.Field{Type: graphql.Float},
		"longitude": &graphql.Field{Type: graphql.Float},
		"comment":   &graphql.Field{Type: graphql.String},
		"title":     &graphql.Field{Type: graphql.String},
		"category":  &graphql.Field{Type: graphql.String},
		"createdAt": &graphql.Field{Type: graphql.String},
	},
})

func CreateSchema(r *Resolver) (*graphql.Schema, error) {
	queryType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"locations": &graphql.Field{
				Type:    graphql.NewList(locationType),
				Resolve: r.ResolveLocations,
			},
			"location": &graphql.Field{
				Type: locationType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{Type: graphql.String},
				},
				Resolve: r.ResolveLocationByID,
			},
		},
	})

	mutationType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"createLocation": &graphql.Field{
				Type: locationType,
				Args: graphql.FieldConfigArgument{
					"latitude":  &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Float)},
					"longitude": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.Float)},
					"comment":   &graphql.ArgumentConfig{Type: graphql.String},
					"title":     &graphql.ArgumentConfig{Type: graphql.String},
					"category":  &graphql.ArgumentConfig{Type: graphql.String},
				},
				Resolve: r.CreateLocation,
			},
		},
	})

	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query:    queryType,
		Mutation: mutationType,
	})

	return &schema, err
}
