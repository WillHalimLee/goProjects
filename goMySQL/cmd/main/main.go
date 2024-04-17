package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/goProjects/goMySQL/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	fmt.Println("Server is starting on localhost:8010")
	log.Fatal(http.ListenAndServe("localhost:8010", r))
}
