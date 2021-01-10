package localstorage

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sync"
	"time"

	recipe "github.com/mikeletux/home-recipes/pkg"
	"github.com/mikeletux/home-recipes/pkg/guid"
)

type LocalStorage struct {
	guidGenerator guid.Guid
	recipes       map[string]*recipe.Recipe
	mux           sync.RWMutex
	filepath      string
}

/*
NewLocalStorage creates a local storage with in memory dict
Parameters:
	recipes -> map of recipes that can be use for sample input data
	filename -> name of the file that holds the data to read from/write to (to persist data in disk)
	guidGenerator -> struct that implement the guid.Guid interface

IMPORTANT: if a sample of recipes is passed, filename MUST be set to an empty string ("") and viceversa, if a filename is passed, a nil should be pass as recipes.
			BOTH CANNOT BE SET. IT WILL PANIC.
			A nil, "" scenario is possible if a only in-memory scenario is desired with no sample dataset.
*/
func NewLocalStorage(recipes map[string]*recipe.Recipe, filepath string, guidGenerator guid.Guid) recipe.RecipeRepository {
	//Use file scenario
	if recipes == nil && len(filepath) > 0 {
		recipes, err := readFromFile(filepath)
		if err != nil {
			panic(err)
		}
		return &LocalStorage{
			filepath:      filepath,
			guidGenerator: guidGenerator,
			recipes:       recipes,
		}
	}
	//Use sample dataset without persistance scenario
	if recipes != nil && len(filepath) == 0 {
		return &LocalStorage{
			filepath:      "",
			guidGenerator: guidGenerator,
			recipes:       recipes,
		}
	}
	//Use empty dataset without persistance scenario
	if recipes == nil && len(filepath) == 0 {
		recipes = make(map[string]*recipe.Recipe)
		return &LocalStorage{
			filepath:      "",
			guidGenerator: guidGenerator,
			recipes:       recipes,
		}
	}
	//Imposible scenario (it panics)
	panic("You cannot instanciate this object with both recipes and filepath set")
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
	if len(l.filepath) > 0 {
		err := writeToFile(l.filepath, l.recipes)
		if err != nil {
			return "", err
		}
		//TO-DO REMOVE RECIPE FROM MAP
	}
	return recipe.ID, nil
}

func (l *LocalStorage) DeleteRecipe(ID string) error {
	l.mux.Lock()
	defer l.mux.Unlock()
	if _, ok := l.recipes[ID]; !ok {
		return fmt.Errorf("There's no recipe with such ID")
	}
	delete(l.recipes, ID)
	if len(l.filepath) > 0 {
		err := writeToFile(l.filepath, l.recipes)
		if err != nil {
			return err
		}
		//TO-DO RECOVER RECIPE TO MAP
	}
	return nil
}

func (l *LocalStorage) UpdateRecipe(ID string, recipe *recipe.Recipe) error {
	//TODO
	return nil
}

func readFromFile(filepath string) (map[string]*recipe.Recipe, error) {
	recipes := make(map[string]*recipe.Recipe)
	//Check if file exists first. If not write an empty map of recipes
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		return recipes, nil
	}
	b, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("The was an error when reading from the file: %s", err)
	}

	err = json.Unmarshal(b, &recipes)
	if err != nil {
		return nil, fmt.Errorf("The was an error when unmarshalling json: %s", err)
	}
	return recipes, nil

}

func writeToFile(filepath string, recipes map[string]*recipe.Recipe) error {
	//Marshall map into bytes
	b, err := json.Marshal(recipes)
	if err != nil {
		return fmt.Errorf("The was an error when marshalling json: %s", err)
	}
	err = ioutil.WriteFile(filepath, b, 0644)
	if err != nil {
		return fmt.Errorf("The was an error when writing to file: %s", err)
	}
	return nil
}
