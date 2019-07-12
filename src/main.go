package main

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"

	"github.com/olebedev/config"
	"github/kdswto/webserver_example/src/routes"
	"go.mongodb.org/mongo-driver/mongo"
)

func getConfig() *config.Config {
	cfg, err := config.ParseYamlFile("config.yaml")
	if err != nil {
		log.Fatal("getConfig: ", err)
	}

	return cfg
}

type Kernel struct {
	config *config.Config
}

func (k *Kernel) getParam(param string) string {
	value, err := k.config.String(param)
	if err != nil {
		log.Fatal(err)
	}
	return value
}

func main() {
	kernel := &Kernel{config: getConfig()}

	appPort := kernel.getParam("port")
	dbHost := kernel.getParam("db.host")
	dbPort := kernel.getParam("db.port")
	dbName := kernel.getParam("db.name")
	dbUri := fmt.Sprintf("mongodb://%s:%s", dbHost, dbPort)
	dbClient, _ := mongo.Connect(context.Background(), options.Client().ApplyURI(dbUri))
	database := dbClient.Database(dbName)

	r := mux.NewRouter()
	r.HandleFunc("/api/users", routes.GetUsers(database)).Methods("GET")
	r.HandleFunc("/api/users/{id}", routes.GetUserById(database)).Methods("GET")

	err := http.ListenAndServe(appPort, r)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
