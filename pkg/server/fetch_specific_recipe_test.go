package server

import (
	"fmt"
	"github.com/mikeletux/home-recipes/cmd/sample-data"
	"github.com/mikeletux/home-recipes/pkg/localstorage"
	"net/http"
	"testing"
	"net/http/httptest"
	"io/ioutil"
	"encoding/json"
	recipe "github.com/mikeletux/home-recipes/pkg"
	"github.com/gorilla/mux"
)

func TestFetchSpecificRecipe(t *testing.T) {
	//Create fake request
	//Use  mux.SetURLVars(r, vars) to set URL variables when testing
	req, err := http.NewRequest("GET", fmt.Sprintf("/api/%s/%s", apiVersion,recipesEndpoint), nil)
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

	//Create http recorder
	rec := httptest.NewRecorder()

	a.fetchSpecificRecipe(rec, req)

	res := rec.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Errorf("expected %d got: %d", http.StatusOK, res.StatusCode)
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("could not read response body %v", err)
	}
	var recipe recipe.Recipe
	err = json.Unmarshal(b, &recipe)

	if recipe.ID != data.SampleRecipes[sampleID].ID {
		t.Errorf("expected %s got: %s", data.SampleRecipes[sampleID].ID, recipe.ID)
	}
}