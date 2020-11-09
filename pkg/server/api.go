package server

import (
	"github.com/gorilla/mux"
	recipe "github.com/mikeletux/home-recipes/pkg"
	"net/http"
)

type api struct {
	router  http.Handler
	storage recipe.RecipeRepository
}

type Server interface {
	Router() http.Handler
}

func New(repo recipe.RecipeRepository) Server {
	a := &api{}

	r := mux.NewRouter()
	//Retrieve all recipes
	r.HandleFunc("/recipes", a.fetchAllRecipes).Methods(http.MethodGet)
	//Add a recipe
	r.HandleFunc("/recipes", a.addRecipe).Methods(http.MethodPost)
	//Retrieve specific recipe
	r.HandleFunc("/recipes/{ID:[a-zA-Z0-9_]+}", a.fetchSpecificRecipe).Methods(http.MethodGet)

	a.router = r
	a.storage = repo

	return a
}

func (a *api) Router() http.Handler {
	return a.router
}
