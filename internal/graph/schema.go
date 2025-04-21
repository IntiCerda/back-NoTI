package graph

import "github.com/graphql-go/graphql"

var userType = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"id":    &graphql.Field{Type: graphql.String},
		"name":  &graphql.Field{Type: graphql.String},
		"email": &graphql.Field{Type: graphql.String},
	},
})

func CreateSchema(r *Resolver) (*graphql.Schema, error) {
	queryType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"user": &graphql.Field{
				Type: userType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{Type: graphql.String},
				},
				Resolve: r.ResolveUser,
			},
			"users": &graphql.Field{
				Type:    graphql.NewList(userType),
				Resolve: r.ResolveUsers,
			},
		},
	})

	mutationType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: graphql.Fields{
			"createUser": &graphql.Field{
				Type: userType,
				Args: graphql.FieldConfigArgument{
					"name":  &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
					"email": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				},
				Resolve: r.CreateUser,
			},
		},
	})

	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query:    queryType,
		Mutation: mutationType,
	})

	return &schema, err
}
