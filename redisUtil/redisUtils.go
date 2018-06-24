package redisUtil

import (
	"log"
	"sync"
	"time"

	"github.com/garyburd/redigo/redis"
)

// 简单的client结构  区分不同server
type RedisClient struct {
	pool    *redis.Pool
	Address string
}

var (
	redisMap map[string]*RedisClient
	mapMutex *sync.RWMutex
)

const (
	// 参考服务端 应当小于服务端
	defaultTimeOut   = 60 * 10
	defaultMaxIdle   = 5
	defaultMaxActive = 20
)

func init() {
	redisMap = make(map[string]*RedisClient)
	mapMutex = new(sync.RWMutex)
}

// redis://:password@host:port/index
func newPool(redisURL string, timeout time.Duration) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     defaultMaxIdle,
		MaxActive:   defaultMaxActive,
		IdleTimeout: timeout,
		Dial: func() (redis.Conn, error) {
			c, err := redis.DialURL(redisURL)
			return c, err
		},
		// t是放回池子的时间，暂时不限制了
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

func GetRedisClient(address string, timeOut time.Duration) *RedisClient {
	var client *RedisClient
	mapMutex.RLock()
	client, ok := redisMap[address]
	mapMutex.RUnlock()
	if !ok {
		client = &RedisClient{Address: address, pool: newPool(address, timeOut)}
		mapMutex.Lock()
		redisMap[address] = client
		mapMutex.Unlock()
	}
	log.Println(client)
	return client
}

func (client *RedisClient) GetConn() redis.Conn {
	return client.pool.Get()
}
