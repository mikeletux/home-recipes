package server

import (
	"github.com/gorilla/mux"
	recipe "github.com/mikeletux/home-recipes/pkg"
	"net/http"
	"fmt"
)

const (
	recipesEndpoint = "/recipes"
)
type api struct {
	router  http.Handler
	storage recipe.RecipeRepository
}

func (a *api) Router() http.Handler {
	return a.router
}

type Server interface {
	Router() http.Handler
}

func New(repo recipe.RecipeRepository) Server {
	a := &api{}

	r := mux.NewRouter()
	//Retrieve all recipes
	r.HandleFunc(recipesEndpoint, a.fetchAllRecipes).Methods(http.MethodGet)
	//Add a recipe
	r.HandleFunc(recipesEndpoint, a.addRecipe).Methods(http.MethodPost)
	//Retrieve specific recipe
	r.HandleFunc(fmt.Sprintf("%s/%s", recipesEndpoint, "{ID:[a-zA-Z0-9_]+}"), a.fetchSpecificRecipe).Methods(http.MethodGet)
	//Delete a specific recipe
	r.HandleFunc(fmt.Sprintf("%s/%s", recipesEndpoint, "{ID:[a-zA-Z0-9_]+}"), a.removeSpecificRecipe).Methods(http.MethodDelete)

	a.router = r
	a.storage = repo

	return a
}
