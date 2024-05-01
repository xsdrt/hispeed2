package cache

import (
	"os"
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/gomodule/redigo/redis"
)

var testRedisCache RedisCache

func TestMain(m *testing.M) {
	s, err := miniredis.Run()
	if err != nil {
		panic((err))
	}
	defer s.Close()

	pool := redis.Pool{
		MaxIdle:     50,
		MaxActive:   1000,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", s.Addr()) //s.Addr gives the addrss for the temp in memory Redis server...
		},
	}

	testRedisCache.Conn = &pool
	testRedisCache.Prefix = "test-hispeed2" // The prfix to append (something not used anywhere else...)

	defer testRedisCache.Conn.Close()

	os.Exit(m.Run())

}
