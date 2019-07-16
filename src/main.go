package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/olebedev/config"
	"github/kdswto/webserver_example/src/routes"
	"log"
	"net/http"
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
		log.Fatal(fmt.Sprintf("Kernel.getParam: param %s not exist", param), err)
	}
	return value
}

func main() {
	kernel := &Kernel{config: getConfig()}

	appPort := kernel.getParam("port")
	dbHost := kernel.getParam("db.host")
	dbPort := kernel.getParam("db.port")
	dbName := kernel.getParam("db.name")
	dbUser := kernel.getParam("db.user")
	dbPassword := kernel.getParam("db.password")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Can't initialize database", err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal("Database not respond", err.Error())
	}

	r := mux.NewRouter()
	r.HandleFunc("/api/users", routes.GetUsers(db)).Methods("GET")
	r.HandleFunc("/api/users/{id}", routes.GetUserById(db)).Methods("GET")
	//r.HandleFunc("/api/users/{id}/deposit", routes.DepositByUserId(database)).Methods("POST")
	// "/api/user/{id}/deposit" POST
	// "/api/user/{id}/withdrawal" POST
	// "/api/user/{id}/block" PUT
	// "/api/user/{id}/transfer" POST

	err = http.ListenAndServe(appPort, r)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
