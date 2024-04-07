package repository

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

type authRedisRepo struct {
	redisClient *redis.Client
}

func NewAuthRedisRepository(redisClient *redis.Client) AuthRedisRepository {
	return &authRedisRepo{redisClient: redisClient}
}

type AuthRedisRepository interface {
	AddTokenToBlackList(ctx context.Context, jti string, expiresAt *time.Time) error
	IsTokenInBlackList(ctx context.Context, jti string) (bool, error)
}

func (repo *authRedisRepo) AddTokenToBlackList(ctx context.Context, jti string, expiresAt *time.Time) error {
	return repo.redisClient.Set(ctx, jti, true, expiresAt.Sub(time.Now())).Err()
}

func (repo *authRedisRepo) IsTokenInBlackList(ctx context.Context, jti string) (bool, error) {
	exists, err := repo.redisClient.Exists(ctx, jti).Result()
	if err != nil {
		return false, err
	}
	return exists == 1, nil
}
