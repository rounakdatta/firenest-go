package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rounakdatta/firenest/firefly"
)

func main() {
	serviceName := "firenest"
	servicePort := "6996"

	router := mux.NewRouter()
	firenestRouter := router.PathPrefix("/" + serviceName).Subrouter()

	firenestRouter.HandleFunc("/", GetRoot).Methods("GET")
	firenestRouter.HandleFunc("/api/parse/sms", ParseSMS).Methods("POST")

	http.Handle("/", router)

	log.Printf("Firenest started on port %s", servicePort)
	log.Fatal(http.ListenAndServe(":"+servicePort, nil))
}

func GetRoot(w http.ResponseWriter, r *http.Request) {
	payload := []byte("OK")
	w.Write(payload)
}

func ParseSMS(w http.ResponseWriter, r *http.Request) {
	message := r.FormValue("message")
	sender := r.FormValue("sender")

	processor := firefly.ParseMessage(message, sender)
	transaction := firefly.CreateTransaction(processor)

	headers := map[string]string{
		"Cache-Control":   "no-cache",
		"Accept":          "application/json, text/plain, /",
		"Content-Type":    "application/json;charset=UTF-8",
		"Accept-Language": "en-US,en;q=0.9",
		"Authorization":   fmt.Sprintf("Bearer %s", firefly.GetPersonalAccessToken()),
	}

	payload := []byte("OK")

	err := firefly.SendFireflyRequest("POST", firefly.GetEndpoint(), transaction, headers)
	if err != nil {
		payload = []byte("ERROR")
	}

	w.Write(payload)
}
