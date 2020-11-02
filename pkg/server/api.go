package server

import (
	"encoding/json"
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
	r.HandleFunc("/recipes", a.fetchRecipes).Methods(http.MethodGet)
	r.HandleFunc("/recipes/{ID:[a-zA-Z0-9_]+}", a.fetchRecipe).Methods(http.MethodGet)

	a.router = r
	a.storage = repo

	return a
}

func (a *api) Router() http.Handler {
	return a.router
}

func (a *api) fetchRecipes(w http.ResponseWriter, r *http.Request) {
	recipes, _ := a.storage.FetchAllRecipes()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(recipes)
}

func (a *api) fetchRecipe(w http.ResponseWriter, r *http.Request) {
	//Get vars
	vars := mux.Vars(r)
	recipe, err := a.storage.FetchRecipeByID(vars["ID"])
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Recipe not found")
		return
	}
	json.NewEncoder(w).Encode(recipe)
}
