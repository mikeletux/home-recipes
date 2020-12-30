package localstorage

import (
	"fmt"
	recipe "github.com/mikeletux/home-recipes/pkg"
	"github.com/rs/xid"
	"sync"
	"time"
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

func (l *LocalStorage) CreateRecipe(recipe *recipe.Recipe) (string, error) {
	l.mux.Lock()
	defer l.mux.Unlock()
	guid := xid.New()
	recipe.ID = guid.String()
	//Set creation and update time to recipe
	time := time.Now()
	recipe.CreationTime = time
	recipe.UpdatedTime = time
	l.recipes[guid.String()] = recipe
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
