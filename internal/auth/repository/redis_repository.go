package repository

import (
	"context"
	"github.com/redis/go-redis/v9"
	"time"
)

type authRedisRepo struct {
	redisClient *redis.Client
}

func NewAuthRedisRepository(redisClient *redis.Client) *authRedisRepo {
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
	return repo.redisClient.Get(ctx, jti).Bool()
}
