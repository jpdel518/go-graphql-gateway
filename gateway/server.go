package main

import (
	"github.com/jpdel518/go-graphql-gateway/gateway/utils"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/jpdel518/go-graphql-gateway/gateway/graph"
	"github.com/jpdel518/go-graphql-gateway/gateway/internal"
)

const defaultPort = "8080"

func init() {
	utils.LoggingSettings(os.Getenv("LOG_FILE"))
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(internal.NewExecutableSchema(internal.Config{Resolvers: &graph.Resolver{}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
