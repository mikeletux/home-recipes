package localstorage

import (
	//"fmt"
	recipe "github.com/mikeletux/home-recipes/pkg"
	"sync"
)

type LocalStorage struct {
	recipes map[string]*recipe.Recipe
	mux     sync.RWMutex
}

// NewLocalStorage creates a local storage
func NewLocalStorage(recipes map[string]*recipe.Recipe) *LocalStorage {
	if recipes == nil {
		recipes = make(map[string]*recipe.Recipe)
	}

	return &LocalStorage{
		recipes: recipes,
	}
}

func (l *LocalStorage) FetchRecipeByID(ID string) (*recipe.Recipe, error) {
	return nil, nil
}

func (l *LocalStorage) FetchAllRecipes() ([]*recipe.Recipe, error) {
	return nil, nil
}
