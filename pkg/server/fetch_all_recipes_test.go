package server

import (
	"github.com/mikeletux/home-recipes/cmd/sample-data"
	"github.com/mikeletux/home-recipes/pkg/localstorage"
	"net/http"
	"testing"
	"net/http/httptest"
	"io/ioutil"
	"encoding/json"
	recipe "github.com/mikeletux/home-recipes/pkg"
)

func TestFetchAllRecipes(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/v1/recipes", nil)
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