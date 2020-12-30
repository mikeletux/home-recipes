package server

import (
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
)

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