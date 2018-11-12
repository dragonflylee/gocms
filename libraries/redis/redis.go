package redis

import (
	"errors"
	"fmt"
	"log"
	"time"

	"gopkg.in/redis.v5"
)

const Nil = redis.Nil

const RetryCount = 5

type RedisConf struct {
	Host string
	Port int
	DB   int
}

type RedisPool struct {
	client *redis.Client
	conf   *RedisConf
}

func NewPool(conf *RedisConf) *RedisPool {
	if conf.Host == "" {
		panic(errors.New("redis config error"))
	}
	if conf.Port == 0 {
		conf.Port = 6379
	}
	var (
		client *redis.Client
		err    error
	)
	for i := 0; i < RetryCount; i++ {
		client = redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%d", conf.Host, conf.Port),
			Password: "",
			DB:       conf.DB,
		})
		_, err = client.Ping().Result()
		if err != nil {
			log.Printf("Failed to connect Redis Server: %v error:%s", conf, err.Error())
			time.Sleep(2 * time.Second)
			log.Printf("Retrying to connect")
		} else {
			return &RedisPool{client, conf}
		}
	}
	panic(err)
}

func (r *RedisPool) NewConn() *redis.Client {
	return r.client
}

func (r *RedisPool) IsNil(err error) bool {
	if err == redis.Nil {
		return true
	}
	return false
}
