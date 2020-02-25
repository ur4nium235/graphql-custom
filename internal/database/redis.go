package database

import (
	"github.com/garyburd/redigo/redis"
)


/**
 *
 * @author: hoangtq
 * @timeCreate: 25/02/2020 14:32
 * To change this template use File | Settings | Editor | File and Code Template | Includes
 * */

func NewClientRedis(url, pass string) *redis.Pool {
	if len(url) <= 0 || len(pass) <= 0 {
		panic("input is required")
	}

	redisPool := &redis.Pool{
		Dial:         func() (redis.Conn, error) {
			return redis.Dial("tcp", url, redis.DialPassword(pass))
		},
	}

	conn := redisPool.Get()
	defer  conn.Close()

	// Test the connection
	_, err := conn.Do("PING")
	if err != nil {
		panic(err)
	}
	return redisPool
}

