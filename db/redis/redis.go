package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

var client *redis.ClusterClient

func Init() {
	client = redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:              []string{"192.168.1.1:1234"}, // ip:port
		NewClient:          nil,
		MaxRedirects:       0,
		ReadOnly:           false,
		RouteByLatency:     false,
		RouteRandomly:      false,
		ClusterSlots:       nil,
		Dialer:             nil,
		OnConnect:          nil,
		Username:           "",
		Password:           "",
		MaxRetries:         0,
		MinRetryBackoff:    0,
		MaxRetryBackoff:    0,
		DialTimeout:        0,
		ReadTimeout:        time.Second,
		WriteTimeout:       time.Second,
		PoolFIFO:           false,
		PoolSize:           100,
		MinIdleConns:       0,
		MaxConnAge:         0,
		PoolTimeout:        0,
		IdleTimeout:        time.Minute,
		IdleCheckFrequency: 0,
		TLSConfig:          nil,
	})

	if client == nil {
		panic("redis init failed")
		return
	}

	fmt.Println("redis init success")
}

func Set(ctx context.Context, key string, value interface{}, expireTime time.Duration) (result string, err error) {
	result, err = client.Set(ctx, key, value, expireTime).Result()
	return
}

// SetNX
// if not lock, return false
func SetNX(ctx context.Context, key string, value interface{}, expireTime time.Duration) (result bool, err error) {
	result, err = client.SetNX(ctx, key, value, expireTime).Result()
	return
}

func Get(ctx context.Context, key string) (result string, err error) {
	result, err = client.Get(ctx, key).Result()
	if err != nil {
		if err != redis.Nil {
		}
	}
	return
}

func GetInt(ctx context.Context, key string) (result int, err error) {

	result, err = client.Get(ctx, key).Int()
	if err != nil {
		if err != redis.Nil {
		}
	}
	return
}

// Incr redis value++
// result is the value after value++
// if the key not exist, return 1
func Incr(ctx context.Context, key string) (result int64, err error) {

	result, err = client.Incr(ctx, key).Result()

	return
}

// Decr redis value--
// result is the value after value--
// if the key not exist, return -1
func Decr(ctx context.Context, key string) (result int64, err error) {

	result, err = client.Decr(ctx, key).Result()

	return
}

func Del(ctx context.Context, key string) (aff int64, err error) {
	aff, err = client.Del(ctx, key).Result()
	if err != nil {
		if err != redis.Nil {
		}
	}
	return
}
