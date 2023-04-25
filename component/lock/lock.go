package lock

import (
	"github.com/go-redis/redis"
	"time"
)

const (
	// 锁前缀
	prefix = "lock:"
)

// Lock 锁
type Lock interface {
	// 加锁
	Lock(second int) error
	// 解锁
	Unlock()
}

// redisLock 基于redis的分布时所
type redisLock struct {
	key         string
	redisClient redis.UniversalClient
}

// NewInstance
func NewInstance(redisClient redis.UniversalClient, key string) Lock {
	return &redisLock{
		key:         prefix + key,
		redisClient: redisClient,
	}
}

// Lock 加锁
func (lock *redisLock) Lock(second int) error {
	boolRes, err := lock.redisClient.SetNX(lock.key, 1, time.Second*time.Duration(second)).Result()

	if boolRes == false {
		return err
	}
	return nil
}

// Unlock 解锁
func (lock *redisLock) Unlock() {
	lock.redisClient.Del(lock.key)
}
