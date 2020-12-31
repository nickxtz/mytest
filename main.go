package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println(time.Now(),"server start...")
	defer fmt.Println("%s server end")

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
		log.Fatalf("server: start faile : %v\n", err)
	}
}

func Pong(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("receive a request from %v\n", r.RemoteAddr)
	json.NewEncoder(w).Encode("hello")
}
