package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/maksimartemev/golang-db-pg-example/model"
	"github.com/maksimartemev/golang-db-pg-example/store"
)

func middleware_logging(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)
		f(w, r)
	}
}

func middleware_set_headers(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		f(w, r)
	}
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	userId := r.URL.Query().Get("user_id")

	if userId != "" {
		userIdInt, err := strconv.Atoi(userId)

		if err != nil {
			http.Error(w, "user_id must be an integer!", http.StatusUnprocessableEntity)
			return
		}

		user, _ := store.GetById(userIdInt)
		if user.ID < 1 {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(user)
		return
	}

	http.Error(w, "Specify user_id field", http.StatusUnprocessableEntity)
}
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	decoder := json.NewDecoder(r.Body)

	var regUser model.RegisterUser
	err := decoder.Decode(&regUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if len(regUser.Email) == 0 {
		http.Error(w, "email is required", http.StatusUnprocessableEntity)
		return
	}
	if len(regUser.Password) == 0 {
		http.Error(w, "password is required", http.StatusUnprocessableEntity)
		return
	}

	createdUser, err := store.Create(regUser.Email, regUser.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(createdUser)
}

func ListUsersHandler(w http.ResponseWriter, r *http.Request) {
	users, err := store.List()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	json.NewEncoder(w).Encode(users)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", middleware_logging(middleware_set_headers(HomeHandler)))
	router.HandleFunc("/register", middleware_logging(middleware_set_headers(RegisterHandler)))
	router.HandleFunc("/list", middleware_logging(middleware_set_headers(ListUsersHandler)))

	log.Fatal(http.ListenAndServe(":8081", router))
}
