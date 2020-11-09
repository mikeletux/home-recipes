package server

import (
	"encoding/json"
	"net/http"
	"time"
)

type recipeSumary struct {
	ID           string    `json:"ID,omitempty"`
	Name         string    `json:"name,omitempty"`
	CreationTime time.Time `json:"creationTime,omitempty"`
	UpdatedTime  time.Time `json:"updatedTime,omitempty"`
}

func (a *api) fetchAllRecipes(w http.ResponseWriter, r *http.Request) {
	sumary := []*recipeSumary{}
	recipes, _ := a.storage.FetchAllRecipes()
	for _, v := range recipes {
		sumary = append(sumary, &recipeSumary{ID: v.ID, Name: v.Name, CreationTime: v.CreationTime, UpdatedTime: v.UpdatedTime})
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(sumary)
}
