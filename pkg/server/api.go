package server

import (
	"encoding/json"
	"fmt"
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
	r.HandleFunc("/recipes", a.fetchRecipes).Methods(http.MethodGet)
	//Add a recipe
	r.HandleFunc("/recipes", a.addRecipe).Methods(http.MethodPost)
	//Retrieve specific recipe
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

func (a *api) addRecipe(w http.ResponseWriter, r *http.Request) {
	newRecipe := recipe.Recipe{}
	w.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(r.Body).Decode(&newRecipe)
	defer r.Body.Close()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(fmt.Sprintf("Recipe format not correct. %v", err))
		return
	}
	err = a.storage.CreateRecipe(&newRecipe)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}
	w.WriteHeader(http.StatusAccepted)
}
