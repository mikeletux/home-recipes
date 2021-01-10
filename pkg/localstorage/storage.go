package localstorage

import (
	"fmt"
	"sync"
	"time"

	recipe "github.com/mikeletux/home-recipes/pkg"
	"github.com/mikeletux/home-recipes/pkg/guid"
)

type LocalStorage struct {
	guidGenerator guid.Guid
	recipes       map[string]*recipe.Recipe
	mux           sync.RWMutex
}

// NewLocalStorage creates a local storage
func NewLocalStorage(recipes map[string]*recipe.Recipe, guidGenerator guid.Guid) *LocalStorage {
	if recipes == nil {
		recipes = make(map[string]*recipe.Recipe)
	}

	return &LocalStorage{
		guidGenerator: guidGenerator,
		recipes:       recipes,
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

func (l *LocalStorage) CreateRecipe(recipe *recipe.Recipe) (string, error) {
	l.mux.Lock()
	defer l.mux.Unlock()
	guid := l.guidGenerator.GetGUID()
	recipe.ID = guid
	//Set creation and update time to recipe
	time := time.Now()
	recipe.CreationTime = time
	recipe.UpdatedTime = time
	l.recipes[guid] = recipe
	return recipe.ID, nil
}

func (l *LocalStorage) DeleteRecipe(ID string) error {
	l.mux.Lock()
	defer l.mux.Unlock()
	if _, ok := l.recipes[ID]; !ok {
		return fmt.Errorf("There's no recipe with such ID")
	}
	delete(l.recipes, ID)
	return nil
}

func (l *LocalStorage) UpdateRecipe(ID string, recipe *recipe.Recipe) error {
	//TODO
	return nil
}
