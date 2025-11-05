package bootstrap

import (
	"log"
	"os"

	"github.com/vgutierrezz/goweb/internal/domain"
	"github.com/vgutierrezz/goweb/internal/user"
)

func NewLogger() *log.Logger {
	return log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)
}

func NewDB() user.DB {
	return user.DB{
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
}
