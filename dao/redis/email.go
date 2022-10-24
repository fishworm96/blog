package redis

import (
	"blog/models"
	"errors"
	"time"

	"github.com/go-redis/redis"
	"go.uber.org/zap"
)

func SaveCode(email, code string) (err error) {
	err = client.Set(email, code, time.Second*300).Err()
	return
}

func GetCode(p *models.EmailLogin) (err error) {
	val, err := client.Get(p.Email).Result()
	if err != nil {
		zap.L().Error("get email failed", zap.Any("val:", val), zap.Error(err))
		if errors.Is(err, redis.Nil) {
			return ErrorCodeInvalid
		}
		return
	}
	if val != p.Code {
		return ErrorCodeIncorrect
	}
	return
}
