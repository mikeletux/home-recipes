package server

import (
	"github.com/mikeletux/home-recipes/cmd/sample-data"
	"github.com/mikeletux/home-recipes/pkg/localstorage"
	"testing"
	"net/http"
	"encoding/json"
	"fmt"
	recipe "github.com/mikeletux/home-recipes/pkg"
	"bytes"
	"net/http/httptest"
	"strings"
)

func TestAddRecipe(t *testing.T) {
	//Create recipe
	recipe := recipe.Recipe{         
		Name: "Beef with boiled potatoes",         
		Image: "https://t2.rg.ltmcdn.com/es/images/5/2/9/estofado_de_carne_con_patatas_73925_orig.jpg",       
		Ingredients: []string{"Beef dices", "Potatoes", "Red wine", "garlic", "onion", "paprika", "beef soup dice"},
		Text: "We put everything in a pressure pot and enjoy!!!",         
	}
	//Create mocked body
	b, err := json.Marshal(recipe)
	if err != nil {
		t.Fatalf("could not create mocked body %v", err)
	}
	req, err := http.NewRequest("POST", fmt.Sprintf("/api/%s/%s", apiVersion,recipesEndpoint), bytes.NewBuffer(b))
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
	recipeID := s[len(s) - 1]
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