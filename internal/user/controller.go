package user

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type (
	Controller func(http.ResponseWriter, *http.Request)
	Endpoint   struct {
		Create Controller
		GetAll Controller
	}

	CreateReq struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Email     string `json:"email"`
	}
)

func MakeEndpoint(ctx context.Context, s Service) Controller {
	return func(w http.ResponseWriter, r *http.Request) {
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

func GetAllUsers(ctx context.Context, s Service, w http.ResponseWriter) {
	//Search users in service
	users, err := s.GetAll(ctx)
	if err != nil {
		MsgResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	DataResponse(w, http.StatusOK, users)
}

func PostUser(ctx context.Context, s Service, w http.ResponseWriter, data interface{}) {
	req := data.(CreateReq)

	//Validation
	if req.FirstName == "" {
		MsgResponse(w, http.StatusBadRequest, "first name is required")
		return
	}

	if req.LastName == "" {
		MsgResponse(w, http.StatusBadRequest, "last name is required")
		return
	}

	if req.Email == "" {
		MsgResponse(w, http.StatusBadRequest, "email is required")
		return
	}
	user, err := s.Create(ctx, req.FirstName, req.LastName, req.Email)
	if err != nil {
		MsgResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	DataResponse(w, http.StatusCreated, user)
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
