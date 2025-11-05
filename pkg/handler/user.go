package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/vgutierrezz/goweb/internal/user"
	"github.com/vgutierrezz/goweb/pkg/transport"
)

func NewUserHttpServer(ctx context.Context, router *http.ServerMux, endpoints user.Endpoints) {

	router.HandleFunc("/users", UserServer(ctx, endpoints))
}

func UserServer(ctx context.Context, endpoints user.Endpoints) func(w http.ResponseWriter, r *http.Request) {

	return func(w http.ResponseWriter, r *http.Request) {
		transport := transport.NewTransport(w, r, ctx)

		switch r.Method {
		case http.MethodGet:
			GetAllUsers(ctx, s, w)
		case http.MethodPost:
			decode := json.NewDecoder(r.Body)
			var newUser CreateReq
			if err := decode.Decode(&newUser); err != nil {
				MsgResponse(w, http.StatusBadRequest, err.Error())
				return
			}
			PostUser(ctx, s, w, newUser)
		default:
			InvalidMethod(w)
		}
	}
}

func InvalidMethod(w http.ResponseWriter) {
	status := http.StatusNotFound
	w.WriteHeader(status) //Sobreescribe status
	fmt.Fprintf(w, `{"status":%d, "message": "method doesn't exist"}`, status)
}

// Transform message to json
func MsgResponse(w http.ResponseWriter, status int, message string) {
	w.WriteHeader(status)
	fmt.Fprintf(w, `{"status":%d, "message": "%s"}`, status, message)
}

// Transform data to json
func DataResponse(w http.ResponseWriter, status int, data interface{}) {
	value, err := json.Marshal(data) //entity to json
	if err != nil {
		MsgResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	w.WriteHeader(status)
	fmt.Fprintf(w, `{"status":%d, "data": %s}`, status, value) //%s for json
}
