package redis_client

import (
	"errors"

	"github.com/go-redis/redis"
)

//RedisClinet global redis client variable
var RedisClinet Redis

//Redis redis clinet wrapper
type Redis struct {
	Client *redis.Client
}

//NewRedis init global redis client
func NewRedis(opts *redis.Options) error {
	client := redis.NewClient(opts)

	pong, err := client.Ping().Result()
	if err != nil {
		return err
	}
	if pong != "PONG" {
		return errors.New("unsuccessfull redis ping")
	}
	RedisClinet.Client = client
	return nil
}
