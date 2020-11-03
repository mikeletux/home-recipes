package main

import (
	"flag"
	"fmt"
	"github.com/mikeletux/home-recipes/cmd/sample-data"
	recipe "github.com/mikeletux/home-recipes/pkg"
	"github.com/mikeletux/home-recipes/pkg/localstorage"
	"github.com/mikeletux/home-recipes/pkg/server"
	"log"
	"net/http"
)

func main() {
	//Flags
	port := flag.Int("port", 8080, "Port where the API will be listening")
	withData := flag.Bool("withData", false, "Initialize with some default data")
	flag.Parse()
	var recipes map[string]*recipe.Recipe
	if *withData {
		recipes = data.SampleRecipes
	}
	storage := localstorage.NewLocalStorage(recipes)
	s := server.New(storage)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), s.Router()))
}
