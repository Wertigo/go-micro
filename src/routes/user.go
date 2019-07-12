package routes

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github/kdswto/webserver_example/src/repositories"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
	"strconv"
)

func GetUsers(db *mongo.Database) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		userRepository := &repositories.UserRepository{Database: db}
		users := userRepository.GetAll()

		usersJson, err := json.Marshal(users)
		if err != nil {
			log.Fatal(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(usersJson)
	}
}

func GetUserById(db *mongo.Database) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		reguestVars := mux.Vars(r)
		idVar := reguestVars["id"]
		id, _ := strconv.Atoi(idVar)
		userRepository := &repositories.UserRepository{Database: db}
		user := userRepository.GetById(id)

		if user == nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
		} else {
			userJson, err := json.Marshal(user)
			if err != nil {
				log.Fatal(err)
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(userJson)
		}
	}
}
