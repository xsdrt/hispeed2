package cache

import "testing"

func TestRedisCache_Has(t *testing.T) {
	err := testRedisCache.Forget("foo") // make sure there is no foo in the cache by calling forget...
	if err != nil {
		t.Error(err)
	}

	inCache, err := testRedisCache.Has("foo")
	if err != nil {
		t.Error(err)
	}

	if inCache {
		t.Error("foo found in cache; foo shouldn't be there")
	}

	err = testRedisCache.Set("foo", "bar")
	if err != nil {
		t.Error(err)
	}

	inCache, err = testRedisCache.Has("foo")
	if err != nil {
		t.Error(err)
	}

	if !inCache {
		t.Error("foo not found in cache; should be there")
	}
}

func testRedisCache_Get(t *testing.T) {
	err := testRedisCache.Set("foo", "bar")
	if err != nil {
		t.Error(err)
	}

	x, err := testRedisCache.Get("foo")
	if err != nil {
		t.Error(err)
	}

	if x != "bar" {
		t.Error("not the correct vale from the cache")
	}
}
