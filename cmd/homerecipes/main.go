package main

import (
	"github.com/mikeletux/home-recipes/pkg/localstorage"
	"github.com/mikeletux/home-recipes/pkg/server"
	"log"
	"net/http"
)

func main() {
	storage := localstorage.NewLocalStorage(nil)
	s := server.New(storage)
	log.Fatal(http.ListenAndServe(":8080", s.Router()))
}
