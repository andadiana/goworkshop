package web

import (
	"net/http"
	"os"
	"github.com/gorilla/mux"
	"fmt"
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
		handlerFunc := log(route.HandlerFunc)
		router.HandleFunc(route.Pattern, handlerFunc).Methods(route.Method)
	}

	port := getPort()
	fmt.Println("+-------------------------------+")
	fmt.Printf("| Starting sever on port: %s\t|\n", port)
	fmt.Println("+-------------------------------+")
	if err := http.ListenAndServe(":" + port, router); err != nil {
		panic(err)
	}
}

func log(funcHandler http.HandlerFunc) http.HandlerFunc{
	return func (rw http.ResponseWriter, r *http.Request) {
		fmt.Println("New REST request to URL: " + r.URL.Path)
		funcHandler(rw, r)
		fmt.Println("REST request ended")
	}
}
