package web

import (
	"net/http"
	"os"
	"github.com/gorilla/mux"
	"fmt"
	"goworkshop/persistence"
	"time"
)

const API_PORT_NAME = "API_PORT"
const API_PORT_VALUE = "8000"

type RestServer struct {
	Port   int
	routes Routes
	router *mux.Router
	Store persistence.DataStore
}

func getPort() string {
	port := os.Getenv(API_PORT_NAME)
	if port != "" {
		return port
	} else {
		return API_PORT_VALUE
	}
}

func (server *RestServer) StartServer() {
	server.initRoutes()
	server.router = mux.NewRouter()

	//setting the endpoints
	for _, route := range server.routes {
		server.router.Handle(route.Pattern, log(route.HandlerFunc)).Methods(route.Method)
	}
	fmt.Println("+-------------------------------+")
	fmt.Printf("| Starting sever on port: %d\t|\n", server.Port)
	fmt.Println("+-------------------------------+")
	if err := http.ListenAndServe(fmt.Sprintf(":%d", server.Port), server.router); err != nil {
		panic(err)
	}
}

func log(routeFunc RouteFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Request to: " + r.URL.Path)
		start := time.Now().UnixNano()
		err := routeFunc(w, r) // call original
		//handle the errors
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			fmt.Errorf("error occurred while processing the request:%v", err)
		}
		end := time.Now().UnixNano()
		fmt.Printf("Request took: %d nano\n", end-start)
	})
}
