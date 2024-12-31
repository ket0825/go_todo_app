package testutil

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/go-redis/redis/v8"
)

func OpenRedisForTest(t *testing.T) *redis.Client {
	t.Helper()

	host := "127.0.0.1"
	port := 36379
	if _, defined := os.LookupEnv("CI"); defined {
		// CI 환경에서는 6379 포트로 연결합니다.
		port = 6379
	}
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", host, port),
		Password: "",
		DB:       0, // use default DB
	})
	if err := client.Ping(context.Background()).Err(); err != nil {
		t.Fatalf("failed to connect redis: %s", err)
	}
	return client

}
