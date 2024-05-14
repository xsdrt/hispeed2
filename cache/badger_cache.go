package cache

import "github.com/dgraph-io/badger/v3"

// From BadgerDB Repository:"BadgerDB is an embeddable, persistent and fast key-value (KV) database"
// "written in pure Go. It is the underlying database for Dgraph, a fast, distributed graph database"
type BadgerCache struct {
	Conn   *badger.DB
	Prefix string
}
