package cache

import (
	"time"

	"github.com/go-redis/redis"
)

type CacheStoreInterface interface {
	Get(key string) (string, error)
	Unmarshal(key string, object interface{}) error
	Set(key string, value interface{}) error
	Fetch(key string, fc func() interface{}) (string, error)
	Delete(key string) error
}

type RedisStoreInterface interface {
	Keys(pattern string) ([]string, error)
	GetByte(key string) ([]byte, error)
	Get(key string) (string, error)
	IncrBy(key string, value int64) (int64, error)
	DecrBy(key string, value int64) (int64, error)
	Unmarshal(key string, object interface{}) error
	Set(key string, value interface{}) error
	Fetch(key string, fc func() interface{}) (string, error)
	Delete(key string) error
	RPush(key string, value ...interface{}) error
	LPop(key string) (string, error)
	Del(key ...string) error
	LRem(key string, value interface{}) error
	LRange(key string) ([]string, error)
	LPush(key string, value interface{}) error
	LLen(key string) (int64, error)
	LIndex(key string, index int64) (string, error)
	HSet(key string, field string, value interface{}) error
	HMSet(key string, fields map[string]interface{}) error
	HGet(key string, field string) (string, error)
	HLen(key string) (int64, error)
	HDel(key string, field string) error
	HExists(key string, field string) (bool, error)
	HGetall(key string) (map[string]string, error)
	SIsMember(key string, field string) (bool, error)
	SMembers(key string) ([]string, error)
	SRem(key string, value ...interface{}) error
	Scan(cursor uint64, match string, count int64) ([]string, uint64, error)
	MGet(keys []string) ([]interface{}, error)
	SAdd(key string, value ...interface{}) error
	SCard(key string) (int64, error)
	SRandMember(key string) (string, error)
	ZRangeWithScores(key string, start, stop int64) ([]redis.Z, error)
	ZAdd(key string, members ...redis.Z) error
	Expire(key string, seconds time.Duration) (bool, error)
	Do(cmd string, key string, seconds string) error
}
