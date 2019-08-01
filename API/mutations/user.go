package mutations

import (
	"../types"
	"fmt"
	"../db"
	"crypto/sha512"
	"github.com/graphql-go/graphql"
	"encoding/hex"
)

// GetCreateUserMutation creates a new user and returns it.
func GetCreateUserMutation(database db.Database) *graphql.Field {
	return &graphql.Field{
		Type: types.UserType,
		Args: graphql.FieldConfigArgument{
			"firstname": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"lastname": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"email": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
			"password": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(params graphql.ResolveParams) (interface{}, error) {
			passwd := sha512.Sum512([]byte(params.Args["password"].(string)))

			user := &types.User{
				Firstname: params.Args["firstname"].(string),
				Lastname:  params.Args["lastname"].(string),
				Email:     params.Args["email"].(string),
				Password:  hex.EncodeToString(passwd[:]),
			}
			fmt.Println(user)

      // Add your user in database here

			return user, nil
		},
	}
}