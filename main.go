package main

import (
	"changeme/views"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

var PORT = "8080"

func init() {
	if port := os.Getenv("PORT"); port != "" {
		PORT = port
	}
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", views.HandleIndex)
	router.HandleFunc("/health", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
		writer.Write([]byte("OK"))
	})
	router.MatcherFunc(func(r *http.Request, rm *mux.RouteMatch) bool {
		if file, err := os.Stat("assets" + r.URL.Path); !os.IsNotExist(err) && !file.IsDir() {
			return true
		}
		return true
	}).HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "assets"+r.URL.Path)
	})
	fmt.Println("Server is running: http://localhost:" + PORT)
	log.Fatal(http.ListenAndServe(":"+PORT, router))
}
