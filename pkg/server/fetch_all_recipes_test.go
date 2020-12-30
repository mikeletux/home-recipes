package server

import (
	"net/http"
	"testing"
	"http"
)

func TestFetchAllRecipes(t *testing.T) {
	req, err := http.NewRequest("GET", recipesEndpoint, nil)
	if err != nil {
		t.Fatalf("couldn't create request: %v", err)
	}
}