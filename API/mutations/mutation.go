package mutations

import (
	"../db"
	"github.com/graphql-go/graphql"
)

// GetRootFields returns all the available mutations.
func GetRootFields(database db.Database) graphql.Fields {
	return graphql.Fields{
		"createUser": GetCreateUserMutation(database),
	}
}