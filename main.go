package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	m "github.com/fatihgokce/gomodule/math"
	"github.com/gorilla/mux"
	// "./math"
)

func Spath(path ...string) {
	for _, value := range path {
		fmt.Println(value)
	}
}
func main() {
	// fmt.Println(m.Add(12, 21))
	fmt.Println(m.Add(21, 2))
	Spath("12", "32", "tg")
	router := mux.NewRouter()
	router.HandleFunc("/latest-messages", LatestMessage)
	server := &http.Server{
		Handler:      router,
		Addr:         ":8081",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(server.ListenAndServe())
}
func LatestMessage(writer http.ResponseWriter, request *http.Request) {
	messages := []Message{
		{"fatih", "gökçe"},
		{"bayram", "gökçe"},
	}
	result, _ := json.Marshal(messages)
	writer.Header().Add("Access-Control-Allow-Origin", "*")
	writer.Write(result)
}

type Message struct {
	Name    string `json:"name"`
	Surname string `json:"surname2"`
}
