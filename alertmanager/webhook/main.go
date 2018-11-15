package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/unrolled/render"
)

const (
	port = 5001
)

var (
	renderer = render.New()

	address = fmt.Sprintf(":%d", port)
)

func getRequestBody(r *http.Request) string {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalf("Could not read webhook body: %v", err)
	}
	defer r.Body.Close()

	return string(body)
}

func webhookHandler(w http.ResponseWriter, r *http.Request) {
	requestBody := getRequestBody(r)

	msg := fmt.Sprintf("Webhook received: %s", requestBody)

	log.Print(msg)

	renderer.Text(w, http.StatusAccepted, "")
}

func main() {
	router := chi.NewRouter()

	router.Post("/alert", webhookHandler)

	log.Fatal(http.ListenAndServe(address, router))
}
