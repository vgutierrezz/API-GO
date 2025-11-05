package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/vgutierrezz/goweb/internal/user"
	"github.com/vgutierrezz/goweb/pkg/bootstrap"
)

func main() {
	server := http.NewServeMux()

	// Initialize components
	db := bootstrap.NewDB()

	logger := bootstrap.NewLogger()
	repo := user.NewRepository(db, logger)
	service := user.NewService(logger, repo)

	ctx := context.Background()

	server.HandleFunc("/users", user.MakeEndpoint(ctx, service))

	fmt.Println("Server started in port 8080")
	log.Fatal(http.ListenAndServe(":8080", server))

}
