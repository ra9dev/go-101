package http

import (
	lru "github.com/hashicorp/golang-lru"
	"lecture-9/internal/store"
)

type ServerOption func(srv *Server)

func WithAddress(address string) ServerOption {
	return func(srv *Server) {
		srv.Address = address
	}
}

func WithStore(store store.Store) ServerOption {
	return func(srv *Server) {
		srv.store = store
	}
}

func WithCache(cache *lru.TwoQueueCache) ServerOption {
	return func(srv *Server) {
		srv.cache = cache
	}
}
