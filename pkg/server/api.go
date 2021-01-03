package server

import (
	"github.com/gorilla/mux"
	recipe "github.com/mikeletux/home-recipes/pkg"
	"net/http"
	"fmt"
	"github.com/rs/cors"
)

const (
	apiVersion = "v1"
	recipesEndpoint = "recipes"
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
	addRecipe(w http.ResponseWriter, r *http.Request)
	fetchAllRecipes(w http.ResponseWriter, r *http.Request)
	fetchSpecificRecipe(w http.ResponseWriter, r *http.Request)
	removeSpecificRecipe(w http.ResponseWriter, r *http.Request)
}

func New(repo recipe.RecipeRepository) Server {
	a := &api{}
	
	r := mux.NewRouter()
	recipeRouter := r.PathPrefix(fmt.Sprintf("/api/%s", apiVersion)).Subrouter()
	//Retrieve all recipes (GET /recipes)
	recipeRouter.HandleFunc(fmt.Sprintf("/%s", recipesEndpoint), a.fetchAllRecipes).Methods(http.MethodGet)
	//Add a recipe (POST /recipes)
	recipeRouter.HandleFunc(fmt.Sprintf("/%s", recipesEndpoint), a.addRecipe).Methods(http.MethodPost)
	//Fetch specific recipe (GET /recipes/{ID:[a-zA-Z0-9_]+})
	recipeRouter.HandleFunc(fmt.Sprintf("/%s/{ID:[a-zA-Z0-9_]+}", recipesEndpoint), a.fetchSpecificRecipe).Methods(http.MethodGet)
	//Delete a specific recipe (DELETE /recipes/{ID:[a-zA-Z0-9_]+})
	recipeRouter.HandleFunc(fmt.Sprintf("/%s/{ID:[a-zA-Z0-9_]+}", recipesEndpoint), a.removeSpecificRecipe).Methods(http.MethodDelete)
	//Update a specific recipe
	//TBD

	//Set default CORS
	a.router = cors.Default().Handler(recipeRouter)
	a.storage = repo

	return a
}
