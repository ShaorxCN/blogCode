package redisUtil

import (
	"log"
	"testing"

	"github.com/garyburd/redigo/redis"
)

func TestGetRedisClient(t *testing.T) {
	client := GetRedisClient("redis://:password@host:port/index", 0)
	conn := client.GetConn()
	log.Println(client)
	defer conn.Close()
	val, err := redis.String(conn.Do("GET", "js"))
	if err != nil {
		t.Fail()
	}

	t.Log(val)

}
