package server

import (
	"github.com/gorilla/mux"
	// recipe "github.com/mikeletux/home-recipes/pkg"
	"net/http"
)

type api struct {
	router http.Handler
}

type Server interface {
	Router() http.Handler
}

func New() Server {
	a := &api{}

	r := mux.NewRouter()
	r.HandleFunc("/recipes", a.fetchRecipes).Methods(http.MethodGet)
	r.HandleFunc("/recipes/{ID:[a-zA-Z0-9_]+}", a.fetchRecipe).Methods(http.MethodGet)

	a.router = r
	return a
}

func (a *api) Router() http.Handler {
	return a.router
}

func (a *api) fetchRecipes(w http.ResponseWriter, r *http.Request) {}
func (a *api) fetchRecipe(w http.ResponseWriter, r *http.Request)  {}