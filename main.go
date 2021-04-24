package main

import (
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

// GetRoot returns OK if server is alive
func GetRoot(w http.ResponseWriter, r *http.Request) {
	payload := []byte("OK")
	w.Write(payload)
}

func ParseSMS(w http.ResponseWriter, r *http.Request) {
	message := r.FormValue("message")
	sender := r.FormValue("sender")

	processor := firefly.ParseMessage(message, sender)

	payload := []byte("OK")
	w.Write(payload)
}
