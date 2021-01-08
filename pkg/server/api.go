package server

import (
	"github.com/gorilla/mux"
	recipe "github.com/mikeletux/home-recipes/pkg"
	"net/http"
	"fmt"
	"github.com/rs/cors"
	"time"
	"encoding/json"
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
	//Creation time is handled by the storage
	ID, err := a.storage.CreateRecipe(&newRecipe)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}
	//Write location header
	w.Header().Set("Location", fmt.Sprintf("%s/%s", r.URL.RequestURI(), ID))
	w.WriteHeader(http.StatusCreated)
}

type recipeSumary struct {
	ID           string    `json:"ID,omitempty"`
	Name         string    `json:"name,omitempty"`
	CreationTime time.Time `json:"creationTime,omitempty"`
	UpdatedTime  time.Time `json:"updatedTime,omitempty"`
	Location     string    `json:"location,omitempty"`
}

func (a *api) fetchAllRecipes(w http.ResponseWriter, r *http.Request) {
	sumary := []*recipeSumary{}
	recipes, _ := a.storage.FetchAllRecipes()
	for _, v := range recipes {
		sumary = append(sumary, &recipeSumary{
												ID: v.ID, 
												Name: v.Name, 
												CreationTime: 
												v.CreationTime, 
												UpdatedTime: v.UpdatedTime, 
												Location: fmt.Sprintf("%s/%s", r.URL.RequestURI(), v.ID)}) //Figure out how to get scheme and hostname with or without reverse proxy
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sumary)
}

func (a *api) fetchSpecificRecipe(w http.ResponseWriter, r *http.Request) {
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

func (a *api) removeSpecificRecipe(w http.ResponseWriter, r *http.Request) {
	//Get recipe ID
	vars := mux.Vars(r)
	err := a.storage.DeleteRecipe(vars["ID"])
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Recipe not found")
		return
	}
	w.WriteHeader(http.StatusAccepted)
}
