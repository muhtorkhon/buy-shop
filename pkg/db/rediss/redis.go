package rediss

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisDb struct {
	Rdb *redis.Client
}

func NewRedis(client *redis.Client) *RedisDb {
	return &RedisDb{
		Rdb: client,
	}
}

func (r *RedisDb) Ping() error {
	return r.Rdb.Ping(context.Background()).Err()
}

func (r *RedisDb) Set(ctx context.Context, key string, value interface{}) error {
	return r.Rdb.Set(ctx, key, value, 0).Err()
}

func (r *RedisDb) Get(ctx context.Context, key string) (string, error) {
	return r.Rdb.Get(ctx, key).Result()
}

func (r *RedisDb) Delete(ctx context.Context, key string) error {
	return r.Rdb.Del(ctx, key).Err()
}

func (r *RedisDb) Exists(ctx context.Context, key string) (bool, error) {
	result, err := r.Rdb.Exists(ctx, key).Result()
	if err != nil {
		return false, err
	}
	return result > 0, nil
}

func (r *RedisDb) SetEx(ctx context.Context, key string, value interface{}, duration time.Duration) error {
	return r.Rdb.SetEx(ctx, key, value, duration).Err()
}
