package graph

import (
	"github.com/jpdel518/go-graphql-gateway/gateway/graph/model"
	"sync"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	userSubscribers map[string]chan<- *model.User
	mutex           sync.Mutex
}

func NewResolver() *Resolver {
	return &Resolver{
		userSubscribers: map[string]chan<- *model.User{},
		mutex:           sync.Mutex{},
	}
}
