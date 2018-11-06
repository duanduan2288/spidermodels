package models

import (
	"fmt"
	"time"

	"github.com/garyburd/redigo/redis"
)

type RedisService struct {
	Addr        string
	Password    string
	MaxIdle     int
	IdleTimeout int
}

var RedisPools map[string]*redis.Pool

func Redis(name ...string) *redis.Pool {
	k := "default"
	if len(name) > 0 {
		k = name[0]
	}

	fmt.Println("RedisPools[k] %v,k:%v", RedisPools[k], k)

	if pool, ok := RedisPools[k]; ok {
		return pool
	}

	panic(fmt.Errorf("unkonw redis %s", k))

	return nil
}

func InitRedis(confs map[string]RedisService) {
	RedisPools := make(map[string]*redis.Pool)
	for k, v := range confs {
		RedisPools[k] = newRedis(v)
	}

	fmt.Printf("InitRedis: init redis done %v\n", RedisPools)
}

func newRedis(conf RedisService) *redis.Pool {
	fmt.Println(conf)
	return &redis.Pool{
		MaxIdle:     conf.MaxIdle,
		IdleTimeout: time.Duration(conf.IdleTimeout) * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", conf.Addr)
			if err != nil {
				return nil, err
			}
			if len(conf.Password) > 0 {
				if _, err := c.Do("AUTH", conf.Password); err != nil {
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
}
