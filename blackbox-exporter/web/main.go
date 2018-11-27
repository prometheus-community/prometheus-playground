package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/unrolled/render"
)

var (
	renderer = render.New()
)

func health(w http.ResponseWriter, r *http.Request) {
	renderer.Text(w, http.StatusOK, "")
}

func hello(w http.ResponseWriter, r *http.Request) {
	renderer.JSON(w, http.StatusOK, map[string]string{"hello": "world"})
}

func main() {
	router := chi.NewRouter()

	router.Get("/health", health)

	router.Get("/hello", hello)

	log.Fatal(http.ListenAndServe(":2112", router))
}
