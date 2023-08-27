package main

import (
	"context"
	"github.com/99designs/gqlgen/graphql"
	"github.com/jpdel518/go-graphql-gateway/gateway/middleware"
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

	srv := handler.NewDefaultServer(internal.NewExecutableSchema(internal.Config{Resolvers: graph.NewResolver(context.Background())}))

	// middleware
	srv.AroundRootFields(func(ctx context.Context, next graphql.RootResolver) graphql.Marshaler {
		log.Println("before RootResolver")
		res := next(ctx)
		// log.Println(res)
		defer log.Println("after RootResolver")
		return res
	})
	srv.AroundOperations(func(ctx context.Context, next graphql.OperationHandler) graphql.ResponseHandler {
		log.Println("before OperationHandler")
		res := next(ctx)
		// log.Println(res)
		defer log.Println("after OperationHandler")
		return res
	})
	srv.AroundResponses(func(ctx context.Context, next graphql.ResponseHandler) *graphql.Response {
		log.Println("before ResponseHandler")
		res := next(ctx)
		// log.Println(res)
		defer log.Println("after ResponseHandler")
		return res
	})
	srv.AroundFields(func(ctx context.Context, next graphql.Resolver) (interface{}, error) {
		log.Println("before Resolver")
		res, err := next(ctx)
		// log.Println(res)
		defer log.Println("after Resolver")
		return res, err
	})

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", middleware.Auth(srv))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
