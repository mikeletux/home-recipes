package server

import (
	"github.com/gorilla/mux"
	recipe "github.com/mikeletux/home-recipes/pkg"
	"net/http"
	//"fmt"
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
	recipeRouter := r.PathPrefix("/api/v1").Subrouter()
	//Retrieve all recipes
	recipeRouter.HandleFunc("/recipes", a.fetchAllRecipes).Methods(http.MethodGet)
	//Add a recipe
	recipeRouter.HandleFunc("/recipes", a.addRecipe).Methods(http.MethodPost)
	//Retrieve specific recipe
	recipeRouter.HandleFunc("/recipes/{ID:[a-zA-Z0-9_]+}", a.fetchSpecificRecipe).Methods(http.MethodGet)
	//Delete a specific recipe
	recipeRouter.HandleFunc("/recipes/{ID:[a-zA-Z0-9_]+}", a.removeSpecificRecipe).Methods(http.MethodDelete)

	a.router = recipeRouter
	a.storage = repo

	return a
}
