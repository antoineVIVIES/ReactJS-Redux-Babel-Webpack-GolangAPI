package queries

import (
	"../types"
	"../db"

	"github.com/graphql-go/graphql"
	_ "github.com/go-sql-driver/mysql"

)

// GetUserQuery returns the queries available against user type.
func GetUserQuery(database db.Database) *graphql.Field {
	return &graphql.Field{
		Type: graphql.NewList(types.UserType),
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			var users []types.User
			return users, nil
		},
	}
}