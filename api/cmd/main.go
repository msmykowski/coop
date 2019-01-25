package main

import (
	"api/db"
	"api/graphql"
	"net/http"
)

func main() {
	db := db.Open()

	defer db.Close()

	graphqlHandler := graphql.New(db)

	http.Handle("/", graphqlHandler)
	http.ListenAndServe(":8080", nil)
}
