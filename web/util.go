package web

import (
	"net/http"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
)

func WriteJson(w http.ResponseWriter, data interface{}) {
	var jsonData, err =json.Marshal(data)
	if err != nil {
		fmt.Println(w, "Failed to marshal data: %s\n", err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, string(jsonData))
}

func ExtractUuid(r *http.Request) string {
	return mux.Vars(r)["uuid"]
}