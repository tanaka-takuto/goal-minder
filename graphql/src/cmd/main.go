package main

import (
	"log"
	"net/http"
	"os"

	gmGraphql "github.com/tanaka-takuto/goal-minder/adapter/graphql"
	"github.com/tanaka-takuto/goal-minder/adapter/graphql/resolver"
	"github.com/tanaka-takuto/goal-minder/infra/db"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "3000"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(gmGraphql.NewExecutableSchema(gmGraphql.Config{Resolvers: &resolver.Resolver{
		DB: db.Con,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
