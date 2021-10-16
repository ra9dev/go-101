package main

import (
	"context"
	"lectures-6/internal/http"
	"log"
)

func main() {
	srv := http.NewServer(context.Background(), ":8080")
	if err := srv.Run(); err != nil {
		log.Println(err)
	}

	srv.WaitForGracefulTermination()
}
