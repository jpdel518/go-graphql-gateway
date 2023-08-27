package infrastructure

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/jpdel518/go-graphql-gateway/gateway/graph/model"
	"github.com/redis/go-redis/v9"
	"log"
	"os"
)

const (
	redisUserAddedSubscription = "users"
)

type RedisClient struct {
	client     redis.UniversalClient
	userPubsub *redis.PubSub
}

func NewRedisClient() *RedisClient {
	// redisClient := redis.NewClient(&redis.Options{
	// 	Addr:     fmt.Sprintf("%v:%v", os.Getenv("REDIS_ENDPOINT"), "6379"),
	// 	Password: "",
	// 	DB:       0,
	// })
	// redisをクラスターで利用する場合は、redis.NewUniversalClientを利用する
	redisClient := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs:    []string{fmt.Sprintf("%v:%v", os.Getenv("REDIS_ENDPOINT"), "6379")},
		Password: "",
		DB:       0,
	})

	return &RedisClient{
		client: redisClient,
	}
}

func (r *RedisClient) TestConnection(ctx context.Context) error {
	// Redisに接続できるか確認
	result, err := r.client.Ping(ctx).Result()
	if err != nil {
		log.Printf("Error pinging Redis: %v", err)
		return err
	}
	log.Printf("Redis ping result: %v", result)
	return nil
}

func (r *RedisClient) SubscribeUserAdded(ctx context.Context) *redis.PubSub {
	// チャンネルを購読
	pubsub := r.client.Subscribe(ctx, redisUserAddedSubscription)
	r.userPubsub = pubsub

	return pubsub
}

func (r *RedisClient) PublishUserAdded(ctx context.Context, user *model.User) error {
	// ユーザー追加イベントをRedisにpublish
	userJson, err := json.Marshal(user)
	if err != nil {
		log.Printf("Error marshalling user: %v", err)
		return err
	}
	if err := r.client.Publish(ctx, redisUserAddedSubscription, userJson).Err(); err != nil {
		log.Printf("Error publishing user: %v", err)
		return err
	}
	return nil
}
