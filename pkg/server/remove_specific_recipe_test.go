package server

import (
	"testing"
	"net/http"
	"github.com/mikeletux/home-recipes/cmd/sample-data"
	"github.com/mikeletux/home-recipes/pkg/localstorage"
	"net/http/httptest"
	"fmt"
	"github.com/gorilla/mux"
)

func TestRemoveSpecificRecipe(t *testing.T) {
	//Create request
	req, err := http.NewRequest("DELETE", fmt.Sprintf("/api/%s/%s", apiVersion,recipesEndpoint), nil)
	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}

	//Set URL vars
	sampleID := "01d3xz7cn92aks9hapsz4d5dp9"
	vars := map[string]string{
        "ID": sampleID,
    }
	req = mux.SetURLVars(req, vars)

	//Create fake API
	repo := localstorage.NewLocalStorage(data.SampleRecipes)
	a := New(repo)

	//Check if recipe exists before calling the handler
	_, err = repo.FetchRecipeByID(sampleID)
	if err != nil {
		t.Errorf("recipe does not exist before executing the handler %v", err)
	}

	//Create http recorder
	rec := httptest.NewRecorder()

	a.removeSpecificRecipe(rec, req)

	res := rec.Result()

	//Do checks
	if res.StatusCode != http.StatusAccepted {
		t.Errorf("expected %d, got: %d", http.StatusAccepted, res.StatusCode)
	}
	//Check if the recipe does not exist after executing the handler
	_, err = repo.FetchRecipeByID(sampleID)
	if err == nil {
		t.Errorf("Recipe exist after handler execution")
	}
}