package server

import (
	"encoding/json"
	"fmt"
	recipe "github.com/mikeletux/home-recipes/pkg"
	"net/http"
)

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
	err = a.storage.CreateRecipe(&newRecipe)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err)
		return
	}
	w.WriteHeader(http.StatusAccepted)
}
