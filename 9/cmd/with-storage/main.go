package main

import (
	"context"
	lru "github.com/hashicorp/golang-lru"
	"lecture-9/internal/http"
	"lecture-9/internal/store/postgres"
)

func main() {
	urlExample := "postgres://localhost:5432/goods"
	store := postgres.NewDB()
	if err := store.Connect(urlExample); err != nil {
		panic(err)
	}
	defer store.Close()

	cache, err := lru.New2Q(6)
	if err != nil {
		panic(err)
	}

	srv := http.NewServer(
		context.Background(),
		http.WithAddress(":8080"),
		http.WithStore(store),
		http.WithCache(cache),
	)
	if err := srv.Run(); err != nil {
		panic(err)
	}

	srv.WaitForGracefulTermination()
}
