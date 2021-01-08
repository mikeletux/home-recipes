package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gorilla/mux"
	data "github.com/mikeletux/home-recipes/cmd/sample-data"
	recipe "github.com/mikeletux/home-recipes/pkg"
	"github.com/mikeletux/home-recipes/pkg/localstorage"
)

func TestAddRecipe(t *testing.T) {
	//Create recipe
	recipe := recipe.Recipe{
		ID:           "",
		Name:         "Beef with boiled potatoes",
		Image:        "https://t2.rg.ltmcdn.com/es/images/5/2/9/estofado_de_carne_con_patatas_73925_orig.jpg",
		Ingredients:  []string{"Beef dices", "Potatoes", "Red wine", "garlic", "onion", "paprika", "beef soup dice"},
		Text:         "We put everything in a pressure pot and enjoy!!!",
		CreationTime: time.Time{},
		UpdatedTime:  time.Time{},
	}
	//Create mocked body
	b, err := json.Marshal(recipe)
	if err != nil {
		t.Fatalf("could not create mocked body %v", err)
	}
	req, err := http.NewRequest("POST", fmt.Sprintf("/api/%s/%s", apiVersion, recipesEndpoint), bytes.NewBuffer(b))
	if err != nil {
		t.Fatalf("could not create mocked request %v", err)
	}

	//Create fake API
	repo := localstorage.NewLocalStorage(data.SampleRecipes)
	a := New(repo)

	//Create recorder
	rec := httptest.NewRecorder()

	//Execute the handler
	a.addRecipe(rec, req)
	res := rec.Result()
	res.Body.Close()

	//Do checks
	if res.StatusCode != http.StatusCreated {
		t.Errorf("expected %d got: %d", http.StatusCreated, res.StatusCode)
	}

	location := res.Header.Get("Location")
	if len(location) == 0 {
		t.Errorf("No location header returned")
	}
	//Get recipe ID
	s := strings.Split(location, "/")
	recipeID := s[len(s)-1]
	gotRecipe, err := repo.FetchRecipeByID(recipeID)
	if err != nil {
		t.Errorf("There was some error when returning the recipe from storage %v", err)
	}
	if gotRecipe.Name != recipe.Name {
		t.Errorf("Recipes names do not match")
	}
	if len(recipe.Ingredients) != len(gotRecipe.Ingredients) {
		t.Errorf("Ingredients lenth do not match")
	}
}

func TestFetchAllRecipes(t *testing.T) {
	req, err := http.NewRequest("GET", fmt.Sprintf("/api/%s/%s", apiVersion, recipesEndpoint), nil)
	if err != nil {
		t.Fatalf("couldn't create request: %v", err)
	}

	repo := localstorage.NewLocalStorage(data.SampleRecipes)
	a := New(repo)

	//Set the recorder to check what the function is returning
	rec := httptest.NewRecorder()

	a.fetchAllRecipes(rec, req)

	//Get result
	res := rec.Result()
	defer res.Body.Close()

	//Should return 200 OK
	if res.StatusCode != http.StatusOK {
		t.Errorf("expected %d, got: %d", http.StatusOK, res.StatusCode)
	}

	//Read response
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("could not read response: %v", err)
	}

	//parse recipes
	var got []*recipe.Recipe

	err = json.Unmarshal(b, &got)
	if err != nil {
		t.Fatalf("could not unmarshall response: %v", err)
	}

	expected := len(data.SampleRecipes)

	if len(got) != expected {
		t.Errorf("expected %v recipes, got: %v recipes", data.SampleRecipes, got)
	}

}

func TestFetchSpecificRecipe(t *testing.T) {
	//Create fake request
	//Use  mux.SetURLVars(r, vars) to set URL variables when testing
	req, err := http.NewRequest("GET", fmt.Sprintf("/api/%s/%s", apiVersion, recipesEndpoint), nil)
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

func TestRemoveSpecificRecipe(t *testing.T) {
	//Create request
	req, err := http.NewRequest("DELETE", fmt.Sprintf("/api/%s/%s", apiVersion, recipesEndpoint), nil)
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
