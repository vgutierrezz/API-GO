package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/vgutierrezz/goweb/internal/domain"
	"github.com/vgutierrezz/goweb/internal/user"
)

func main() {
	server := http.NewServeMux()
	db := user.DB{
		Users: []domain.User{
			{
				ID: 1, FirstName: "Victor",
				LastName: "Gutiérrez",
				Email:    "victor@example.com",
			}, {
				ID: 2, FirstName: "Ana",
				LastName: "López",
				Email:    "ana@example.com",
			}, {
				ID: 3, FirstName: "Luis",
				LastName: "Martínez",
				Email:    "luis@example.com",
			}},
		MaxUserID: 3,
	}

	logger := log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)
	repo := user.NewRepository(db, logger)
	service := user.NewService(logger, repo)

	ctx := context.Background()

	server.HandleFunc("/users", user.MakeEndpoint(ctx, service))

	fmt.Println("Server started in port 8080")
	log.Fatal(http.ListenAndServe(":8080", server))

}
