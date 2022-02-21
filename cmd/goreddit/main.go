package main

import (
	"log"
	"net/http"

	"github.com/aerovize/goreddit/postgres"
	"github.com/aerovize/goreddit/web"
)

func main() {
	store, err := postgres.NewStore("postgres://postgres:secret@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	h := web.NewHandler(store)
	http.ListenAndServe(":3000", h)
}
