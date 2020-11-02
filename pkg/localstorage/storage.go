package localstorage

import (
	"fmt"
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
	l.mux.RLock()
	defer l.mux.RUnlock()
	v, ok := l.recipes[ID]
	if !ok {
		return nil, fmt.Errorf("[INFO] There's no recipe with that ID")
	}
	return v, nil
}

func (l *LocalStorage) FetchAllRecipes() ([]*recipe.Recipe, error) {
	l.mux.RLock()
	defer l.mux.RUnlock()
	values := make([]*recipe.Recipe, 0, len(l.recipes))
	for _, v := range l.recipes {
		values = append(values, v)
	}
	return values, nil
}
