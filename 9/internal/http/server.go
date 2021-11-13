package http

import (
	"context"
	"github.com/go-chi/chi"
	lru "github.com/hashicorp/golang-lru"
	"lecture-9/internal/store"
	"log"
	"net/http"
	"time"
)

type Server struct {
	ctx         context.Context
	idleConnsCh chan struct{}
	store       store.Store
	cache       *lru.TwoQueueCache

	Address string
}

func NewServer(ctx context.Context, opts ...ServerOption) *Server {
	srv := &Server{
		ctx:         ctx,
		idleConnsCh: make(chan struct{}),
	}

	for _, opt := range opts {
		opt(srv)
	}

	return srv
}

func (s *Server) basicHandler() chi.Router {
	r := chi.NewRouter()

	categoriesResource := NewCategoryResource(s.store, s.cache)
	r.Mount("/categories", categoriesResource.Routes())

	return r
}

func (s *Server) Run() error {
	srv := &http.Server{
		Addr:         s.Address,
		Handler:      s.basicHandler(),
		ReadTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 30,
	}
	go s.ListenCtxForGT(srv)

	log.Println("[HTTP] Server running on", s.Address)
	return srv.ListenAndServe()
}

func (s *Server) ListenCtxForGT(srv *http.Server) {
	<-s.ctx.Done() // блокируемся, пока контекст приложения не отменен

	if err := srv.Shutdown(context.Background()); err != nil {
		log.Printf("[HTTP] Got err while shutting down^ %v", err)
	}

	log.Println("[HTTP] Proccessed all idle connections")
	close(s.idleConnsCh)
}

func (s *Server) WaitForGracefulTermination() {
	// блок до записи или закрытия канала
	<-s.idleConnsCh
}
