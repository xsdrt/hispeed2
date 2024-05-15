package cache

import (
	"log"
	"os"
	"testing"
	"time"

	"github.com/alicebob/miniredis/v2"
	"github.com/dgraph-io/badger/v3"
	"github.com/gomodule/redigo/redis"
)

var testRedisCache RedisCache
var testBadgerCache BadgerCache

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

	_ = os.RemoveAll("./testdata/tmp/badger") // before starting a new test make sure old test DB is deleted...

	// create a badger database(test) for a test cache...
	if _, err := os.Stat("./testdata/tmp"); os.IsNotExist(err) {
		err := os.Mkdir("./testdata/tmp", 0755)
		if err != nil {
			log.Fatal(err)
		}
	}
	err = os.Mkdir("./testdata/tmp/badger", 0755)
	if err != nil {
		log.Fatal(err)
	}

	db, _ := badger.Open(badger.DefaultOptions("./testdata/tmp/badger")) // cretae DB if doesn't exist...
	testBadgerCache.Conn = db

	os.Exit(m.Run())

}
