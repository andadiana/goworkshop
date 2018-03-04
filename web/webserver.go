package web

import (
	"net/http"
	"os"
	"github.com/gorilla/mux"
)

const API_PORT_NAME = "API_PORT"
const API_PORT_VALUE = "8000"

func getPort() string {
	port := os.Getenv(API_PORT_NAME)
	if port != "" {
		return port
	} else {
		return API_PORT_VALUE
	}
}

// using Gorilla mux

func StartServer() {
	router := mux.NewRouter()

	//setting the endpoints
	for _, route := range routes {
		router.HandleFunc(route.Pattern, route.HandlerFunc).Methods(route.Method)
	}

	port := getPort()
	if err := http.ListenAndServe(":" + port, router); err != nil {
		panic(err)
	}
}

//TODO: Methods for author: GET, POST, DELETE, etc.
