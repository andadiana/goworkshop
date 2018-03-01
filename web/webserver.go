package web

import (
	"net/http"
	"os"
	"fmt"
	"goworkshop/model"
	"encoding/json"
	"github.com/gorilla/mux"
	"strings"
	"io/ioutil"
)

const API_PORT_NAME = "API_PORT"
const API_PORT_VALUE = "8000"

func StartServer() {
	http.HandleFunc("/books", getAllBooks)
	http.HandleFunc("/authors",getAllAuthors)
	//http.HandleFunc("/authors/",getAuthorsByUuid)
	var port = getPort()
	fmt.Println("+-------------------------------+")
	fmt.Printf("| Starting sever on port: %s\t|\n", port)
	fmt.Println("+-------------------------------+")
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		panic(err)
	}
}

func getAllBooks(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Request method:", r.Method)
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case "GET":
		serializedData, err := json.Marshal(model.Books)
		if err != nil {
			panic(err)
		}
		fmt.Fprintln(w, string(serializedData))
		break
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintln(w, "Method not supported!")
		break
	}

}

func getAuthorsByUuid(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path //authors/<uuid>
	uuid := strings.TrimPrefix(path, "/authors/")
	for _, author := range model.Authors{
		if author.UUID == uuid {
			if serializedData, err := json.Marshal(model.Authors); err != nil {
				fmt.Println(w, "{\"errorMessage\":%s\"}", err.Error())
				return
			}else {
				fmt.Println(w, string(serializedData))
				return
			}

		}
	}
	w.WriteHeader(http.StatusNotFound)
}

func getAllAuthors(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	serializedData, err := json.Marshal(model.Authors)
	if err != nil {
		panic(err)
	}
	fmt.Fprintln(w, string(serializedData))
}

func getPort() string {
	port := os.Getenv(API_PORT_NAME)
	if port != "" {
		return port
	} else {
		return API_PORT_VALUE
	}
}

func createAuthor(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		fmt.Println(w, "{\"errorMessage\":\"%s\"}", err.Error())
		return
	} else {
		var author model.AuthorDto
		if err := json.Unmarshal(body, &author); err != nil {
			fmt.Println(w, "{\"errorMessage\":\"%s\"}", err.Error())
			return
		} else {
			model.Authors = append(model.Authors, author)
			if serializedData, err := json.Marshal(model.Authors); err != nil {
				fmt.Println(w, "{\"errorMessage\":%s\"}", err.Error())
				return
			}else {
				fmt.Println(w, string(serializedData))
				return
			}
		}
	}
}

// using Gorilla mux

func StartWebServer() {
	gorillaMux := mux.NewRouter()

	gorillaMux.HandleFunc("/authors", getAllAuthors).Methods("GET") //endpoint
	gorillaMux.HandleFunc("/authors/{uuid}", getAuthorsByUuid).Methods("GET")

	if err := http.ListenAndServe(":8000", gorillaMux); err != nil {
		panic(err)
	}
}

//TODO: Methods for author: GET, POST, DELETE, etc.
