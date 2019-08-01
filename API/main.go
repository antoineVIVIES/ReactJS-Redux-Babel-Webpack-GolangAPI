package main

import (
	"log"
	"net/http"
	"./queries"
	"./mutations"
	"./db"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

func main() {
	database = db.OpenDB()
	schemaConfig := graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name:   "RootQuery",
			Fields: queries.GetRootFields(database),
		}),
		Mutation: graphql.NewObject(graphql.ObjectConfig{
			Name:   "RootMutation",
			Fields: mutations.GetRootFields(database),
		}),
	}

	schema, err := graphql.NewSchema(schemaConfig)

	if err != nil {
		log.Fatalf("Failed to create new schema, error: %v", err)
	}

	httpHandler := handler.New(&handler.Config{
		Schema: &schema})
	http.Handle("/",  httpHandler)
	log.Print("ready: listening...\n")

	http.ListenAndServe(":8383", nil)
}