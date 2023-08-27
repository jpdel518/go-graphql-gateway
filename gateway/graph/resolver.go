package graph

import (
	"context"
	"encoding/json"
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
	pubsub := redisClient.SubscribeUserAdded(ctx)

	r := &Resolver{
		redisClient:     redisClient,
		userSubscribers: make(map[string]chan<- *model.User),
		mutex:           sync.Mutex{},
	}

	go func() {
		// このGoルーチンは、loopを抜けることなく常に実行される
		pubsubCh := pubsub.Channel()

		// チャンネルからメッセージを受信
		for msg := range pubsubCh {
			// 受信したメッセージはJSON形式
			// JSONを構造体に変換
			user := &model.User{}
			if err := json.Unmarshal([]byte(msg.Payload), user); err != nil {
				log.Printf("Error unmarshalling response body: %v", err)
				continue
			}

			// 購読しているクライアントにRedisから受信したメッセージを送信
			r.mutex.Lock()
			for _, ch := range r.userSubscribers {
				ch <- user
			}
			r.mutex.Unlock()
		}
	}()

	return r
}
