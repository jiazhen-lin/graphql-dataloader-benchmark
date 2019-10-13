package main

import (
	"log"
	"net/http"

	gql "github.com/jiazhen-lin/graphql-dataloader-benchmark/gql"
)

func main() {
	gql.CreateTestData()
	http.Handle("/query", gql.GetHandler(true))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
