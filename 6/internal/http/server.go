package http

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Server struct {
	ctx         context.Context
	idleConnsCh chan struct{}

	Address string
}

func NewServer(ctx context.Context, address string) *Server {
	return &Server{
		ctx:         ctx,
		Address:     address,
		idleConnsCh: make(chan struct{}),
	}
}

func adder(w http.ResponseWriter, r *http.Request) {
	numbers := r.URL.Path[1:] // "123"

	sum := 0
	for _, nStr := range numbers {
		n, err := strconv.Atoi(string(nStr)) // "1" -> 1
		if err != nil {
			fmt.Fprintf(w, "Got err: %v", err)
			return
		}

		sum += n
	}

	fmt.Fprintf(w, "Sum of path: %d", sum)
}

func (s *Server) Run() error {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world!"))
	})
	mux.HandleFunc("/123", adder)

	srv := &http.Server{
		Addr:         s.Address,
		Handler:      mux,
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
