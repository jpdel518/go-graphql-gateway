package graph

import (
	"context"
	"github.com/jpdel518/go-graphql-gateway/gateway/graph/model"
	"github.com/jpdel518/go-graphql-gateway/gateway/infrastructure"
	"log"
	"sync"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	redisClient     *infrastructure.RedisClient
	userSubscribers map[string]chan<- *model.User
	mutex           sync.Mutex
}

func NewResolver(ctx context.Context) *Resolver {
	redisClient := infrastructure.NewRedisClient()
	if err := redisClient.TestConnection(ctx); err != nil {
		log.Fatalf("Error connecting to Redis: %v", err)
	}

	r := &Resolver{
		redisClient:     redisClient,
		userSubscribers: make(map[string]chan<- *model.User),
		mutex:           sync.Mutex{},
	}

	redisClient.SubscribeUserAdded(ctx, &r.mutex, r.userSubscribers)

	return r
}
