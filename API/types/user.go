package types

import (
	"github.com/graphql-go/graphql"
	"fmt"
)

// User type definition.
type User struct {
	ID        int    `db:"id" json:"id"`
	Firstname string `db:"firstname" json:"firstname"`
	Lastname  string `db:"lastname" json:"lastname"`
	Email     string `db:"email" json:"email"`
	Password  string `db:"password" json:"password"`
	Token	  string `db:"token" json:"token"`
}

// UserType is the GraphQL schema for the user type.
var UserType = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"id":         &graphql.Field{Type: graphql.Int},
		"firstname":  &graphql.Field{Type: graphql.String},
		"lastname":   &graphql.Field{Type: graphql.String},
		"email":	  &graphql.Field{Type: graphql.String},
		"password":	  &graphql.Field{Type: graphql.String},
		"token":	  &graphql.Field{Type: graphql.String},
		"roles": 	  &graphql.Field{
			Type: graphql.NewList(RoleType),
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				var roles []Role

				userID := params.Source.(User).ID
				fmt.Println(userID)
				// Implement logic to retrieve user associated roles from user id here.

				return roles, nil
			},
		},
	},
})