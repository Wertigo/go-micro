package routes

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"github/kdswto/webserver_example/src/repositories"
	"log"
	"net/http"
	"strconv"
)

func GetUsers(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
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

func GetUserById(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		reguestVars := mux.Vars(r)
		idVar := reguestVars["id"]
		userId, _ := strconv.Atoi(idVar)
		userRepository := &repositories.UserRepository{Database: db}
		user := userRepository.GetById(userId)

		if user == nil {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
		}

		userJson, err := json.Marshal(user)
		if err != nil {
			log.Fatal(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(userJson)
	}
}

//func DepositByUserId(db *mongo.Database) func(w http.ResponseWriter, r *http.Request) {
//	return func(w http.ResponseWriter, r *http.Request) {
//		reguestVars := mux.Vars(r)
//		idVar := reguestVars["id"]
//		userId, _ := strconv.Atoi(idVar)
//		userRepository := &repositories.UserRepository{Database: db}
//		user := userRepository.GetById(userId)
//
//		if user == nil {
//			w.Header().Set("Content-Type", "application/json")
//			w.WriteHeader(http.StatusNotFound)
//		}
//
//		walletRepository := repositories.WalletRepository{Database: db}
//		wallet := walletRepository.GetByUserId(userId)
//		user.Wallet = wallet
//
//		userJson, err := json.Marshal(user)
//		if err != nil {
//			log.Fatal(err)
//		}
//		w.Header().Set("Content-Type", "application/json")
//		w.WriteHeader(http.StatusOK)
//		w.Write(userJson)
//	}
//}
