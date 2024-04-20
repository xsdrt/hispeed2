package cache

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

type Cache interface { // List all the functions must implement to satisfy this interface (go's interface system is so simple its great!!)
	Has(string) (bool, error)              // Does the cache have something (string/the key)
	Get(string) (interface{}, error)       // Use the string/key and use an  interface since a cache can store/return anything, also a potenial error..
	Set(string, interface{}, ...int) error // Set what we want to store (string/key;The ...int is where we set the expirery for the cache I.E. 60 secs or 25,000secs etc...
	Forget(string) error                   // Take it out of cache by key(whatever you named the value/key)
	EmptyByMatch(string) error             // where we can forget everything in the cache by a pattern I.E. everything that starts with letter a , or whatever...
	Empty() error                          // Just empty everything no paremeters, just everything...
}

type RedisCache struct {
	Conn   *redis.Pool
	Prefix string // use prefix(unique);case two or more app(s) use redis w/same key i.e. tempid then Forget func and all cached keys gone now...
}

type Entry map[string]interface{} // a map of serialized item(s) to pull out and deserialize...

// A simple check if a given key exist in redis cache for refernce on how works TODO: Need to implement the hash...
func (c *RedisCache) Has(str string) (bool, error) {
	key := fmt.Sprintf("%s:%s", c.Prefix, str) // prepend prefix to the key(str), so i.e. <prefix>:<whatever the user gave>
	conn := c.Conn.Get()
	defer conn.Close()

	ok, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		return false, err
	}

	return ok, nil
}
