package server

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

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
