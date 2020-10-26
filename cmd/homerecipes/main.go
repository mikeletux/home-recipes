package main

import (
	"github.com/mikeletux/home-recipes/pkg/localStorage"
	"github.com/mikeletux/home-recipes/pkg/server"
	"log"
	"net/http"
)

func main() {
	s := server.New()
	storage := localStorage.NewLocalStorage()
	log.Fatal(http.ListenAndServe(":8080", s.Router()))
}
