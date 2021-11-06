package main

import (
	"context"
	"lecture-8/internal/http"
	"lecture-8/internal/store/postgres"
)

func main() {
	urlExample := "postgres://localhost:5432/goods"
	store := postgres.NewDB()
	if err := store.Connect(urlExample); err != nil {
		panic(err)
	}
	defer store.Close()

	srv := http.NewServer(context.Background(), ":8080", store)
	if err := srv.Run(); err != nil {
		panic(err)
	}

	srv.WaitForGracefulTermination()
}
