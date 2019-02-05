package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"user/conf"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("server start...")
	defer fmt.Println("server end")

	r := mux.NewRouter()
	r.HandleFunc("/pong", Pong).Methods(http.MethodGet)
	http.Handle("/", r)

	server := http.Server{
		Addr:           ":12346",
		ReadTimeout:    6 * time.Second,
		WriteTimeout:   6 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	err := server.ListenAndServe()
	if err != nil {
		conf.Logger().Fatalf("server: start faile : %v\n", err)
	}
}

func Pong(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("receive a request from %v\n", r.RemoteAddr)
	json.NewEncoder(w).Encode("hello")
}
