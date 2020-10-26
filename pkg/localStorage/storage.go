package localStorage

import (
	"fmt"
	recipe "github.com/mikeletux/home-recipes/pkg"
	"sync"
)

type LocalStorage struct {
	recipes map[string]*recipe.Recipe
	mux     sync.RWMutex
}

func NewLocalStorage() *LocalStorage {
	storage := make(map[string]*recipe.Recipe)

	return &LocalStorage{recipes: storage}
}
