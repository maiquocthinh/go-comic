package repository

import (
	"context"
	"github.com/maiquocthinh/go-comic/pkg/utils"
	"github.com/redis/go-redis/v9"
	"time"
)

const (
	resetPassCodePrefix = "rpc:"
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
	AddHashedCodeResetPassword(ctx context.Context, email, hashedCode string, expiresAt *time.Time) error
	GetHashedCodeResetPassword(ctx context.Context, email string) (string, error)
	RemoveHashedCodeResetPassword(ctx context.Context, email string) error
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

func (repo *authRedisRepo) AddHashedCodeResetPassword(ctx context.Context, email string, hashedCode string, expiresAt *time.Time) error {
	key := resetPassCodePrefix + utils.GetMD5Hash(email)
	return repo.redisClient.Set(ctx, key, hashedCode, expiresAt.Sub(time.Now())).Err()
}

func (repo *authRedisRepo) GetHashedCodeResetPassword(ctx context.Context, email string) (string, error) {
	key := resetPassCodePrefix + utils.GetMD5Hash(email)
	val, err := repo.redisClient.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return val, nil
}

func (repo *authRedisRepo) RemoveHashedCodeResetPassword(ctx context.Context, email string) error {
	key := resetPassCodePrefix + utils.GetMD5Hash(email)
	_, err := repo.redisClient.Del(ctx, key).Result()
	if err != nil {
		return err
	}
	return nil
}
