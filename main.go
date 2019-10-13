package main

import (
	"log"
	"net/http"
	gql "test-data-loader/gql"
)

func main() {
	gql.CreateTestData()
	http.Handle("/query", gql.GetHandler(false))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
