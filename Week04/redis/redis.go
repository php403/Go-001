// +build wireinject
package redis

import (
	"fmt"
	"github.com/go-redsync/redsync"
	"github.com/gomodule/redigo/redis"
	"github.com/google/wire"
	"time"
)


var ProviderSet = wire.NewSet(NewRedis, NewOptionsRedis)


// Options is  configuration of database
type Options struct {
	Host string
	Port int
	User string
	PassWd string
	DbName string
}
var pool *redis.Pool
var redisLock *redsync.Redsync

func NewOptionsRedis() *Options {
	return &Options{"139.9.141.27", 3306, "pg719", "pg719@1996", "demo"}
}

func NewRedis(o *Options) error  {
	pool = &redis.Pool{
		MaxIdle:     20,
		IdleTimeout: 240 * time.Second,
		MaxActive:   50,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", fmt.Sprintf("%s:%s", o.Host, o.Port))
			if err != nil {
				return nil, err
			}
			if o.PassWd != "" {
				if _, err := c.Do("AUTH", o.PassWd); err != nil {
					c.Close()
					return nil, err
				}
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}

	redisLock = redsync.New([]redsync.Pool{pool})
	return nil
}

func GetRedisLock(key string, expireTime time.Duration) *redsync.Mutex  {
	return redisLock.NewMutex(key, redsync.SetExpiry(expireTime))
}