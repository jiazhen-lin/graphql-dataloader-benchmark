package gql

import (
	"encoding/json"
	"net/http"

	graphql "github.com/graph-gophers/graphql-go"
)

var (
	Schema = `
		schema {
			query: Query
		}
		type Query {
			users: [User!]!
		}  
		type User {
			name: String!
			posts: [Post!]!
		}  
		type Post {
			text: String!
		}
	`
)

func GetHandler(loaderEnable bool) http.Handler {
	return &handler{
		schema:       graphql.MustParseSchema(Schema, &Resolver{}, graphql.MaxParallelism(50)),
		loaderEnable: loaderEnable,
	}
}

type handler struct {
	schema       *graphql.Schema
	loaderEnable bool
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var params struct {
		Query         string                 `json:"query"`
		OperationName string                 `json:"operationName"`
		Variables     map[string]interface{} `json:"variables"`
	}
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// add data loader in context
	ctx := Attach(r.Context(), h.loaderEnable)

	response := h.schema.Exec(ctx, params.Query, params.OperationName, params.Variables)
	responseJSON, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(responseJSON)
}
