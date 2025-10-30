package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Struct for user
type User struct {
	ID        uint64 `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}

// Generate users
var users []User

// Variable to store users
var maxUserID uint64

// Func init to add some users
func init() {
	users = append(users, User{ID: 1, FirstName: "John", LastName: "Doe", Email: "john.doe@example.com"})
	users = append(users, User{ID: 2, FirstName: "Jane", LastName: "Smith", Email: "jane.smith@example.com"})
	users = append(users, User{ID: 3, FirstName: "Alice", LastName: "Johnson", Email: "alice.johnson@example.com"})
	maxUserID = 3
}

func main() {

	http.HandleFunc("/users", UserServer)

	fmt.Println("Server started in port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func UserServer(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		GetAllUsers(w)
	case http.MethodPost:
		decode := json.NewDecoder(r.Body)
		var newUser User
		if err := decode.Decode(&newUser); err != nil {
			MsgResponse(w, http.StatusBadRequest, err.Error())
			return
		}
		PostUser(w, newUser)
	default:
		InvalidMethod(w)
	}
}

func GetAllUsers(w http.ResponseWriter) {
	DataResponse(w, http.StatusOK, users)
}

func PostUser(w http.ResponseWriter, data interface{}) {
	newUser := data.(User) //Cast to User struct
	//Validation
	if newUser.FirstName == "" {
		MsgResponse(w, http.StatusBadRequest, "first name is required")
		return
	}

	if newUser.LastName == "" {
		MsgResponse(w, http.StatusBadRequest, "last name is required")
		return
	}

	if newUser.Email == "" {
		MsgResponse(w, http.StatusBadRequest, "email is required")
		return
	}

	maxUserID++            //Increment ID
	newUser.ID = maxUserID //Assign new ID
	users = append(users, newUser)
	DataResponse(w, http.StatusCreated, newUser)
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
