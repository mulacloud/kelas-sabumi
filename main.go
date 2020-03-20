package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/mulacloud/kelas-sabumi/controllers"
)

func handleLanding(w http.ResponseWriter, req *http.Request){
   w.Header().Set("Content-Type", "application/json")
   fmt.Fprintf(w, "{\"version\": \"1.0.10.1\"}")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handleLanding).Methods("GET")
	r.HandleFunc("/document/", controllers.ListDocument).Methods("GET")
	r.HandleFunc("/subject/", controllers.ListSubject).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", r))
}
