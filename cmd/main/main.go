package main

import (
	"log"
	"net/http"

	"github.com/Hosseinhgz/Go-bookstore-MYSQL-RestAPI/pkg/routes"
	"github.com/github.com/Hosseinhgz/Go-bookstore-MYSQL-RestAPI/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
