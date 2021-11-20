package http

import (
	lru "github.com/hashicorp/golang-lru"
	"lecture-10/internal/message_broker"
	"lecture-10/internal/store"
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

func WithBroker(broker message_broker.MessageBroker) ServerOption {
	return func(srv *Server) {
		srv.broker = broker
	}
}
