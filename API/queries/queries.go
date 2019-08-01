package queries

import (
	"github.com/graphql-go/graphql"
	"../db"
)

// GetRootFields returns all the available queries.
func GetRootFields(db db.Database) graphql.Fields {
	return graphql.Fields{
		"users": GetUserQuery(db),
	}
}