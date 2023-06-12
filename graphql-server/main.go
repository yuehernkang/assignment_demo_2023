package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/TikTokTechImmersion/assignment_demo_2023/http-server/kitex_gen/rpc"
	"github.com/TikTokTechImmersion/assignment_demo_2023/http-server/kitex_gen/rpc/imservice"
	"github.com/TikTokTechImmersion/assignment_demo_2023/http-server/proto_gen/api"
	"github.com/graphql-go/graphql"
)

var cli imservice.Client

func main() {
	// Define the GraphQL schema
	fields := graphql.Fields{
		"hello": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return sendMessage()
			},
		},
	}

	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	schema, err := graphql.NewSchema(schemaConfig)

	if err != nil {
		log.Fatalf("Failed to create GraphQL schema: %v", err)
	}

	// Define the GraphQL handler function
	handler := func(w http.ResponseWriter, r *http.Request) {
		// Parse the GraphQL query from the request body
		queryParams := r.URL.Query()
		query := queryParams.Get("query")

		// Execute the query
		result := graphql.Do(graphql.Params{
			Schema:        schema,
			RequestString: query,
		})

		// Return the query result as JSON
		json.NewEncoder(w).Encode(result)
	}

	// Set up the HTTP server
	http.HandleFunc("/graphql", handler)
	log.Println("GraphQL server is running on http://localhost:6969/graphql")
	log.Fatal(http.ListenAndServe(":6969", nil))

}

func sendMessage(ctx context.Context) {
	var req api.SendRequest

	resp, err := cli.Send(ctx, &rpc.SendRequest{
		Message: &rpc.Message{
			Chat:   req.Chat,
			Text:   req.Text,
			Sender: req.Sender,
		},
	})
}
