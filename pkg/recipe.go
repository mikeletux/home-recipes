package recipe

import (
	// "encoding/json"
	"time"
)

type Recipe struct {
	ID           string    `json:"id,omitempty"`
	Name         string    `json:"name,omitempty"`
	Image        string    `json:"image,omitempty"`
	Ingredients  []string  `json:"ingredients,omitempty"`
	Text         string    `json:"text,omitempty"`
	CreationTime time.Time `json:"creationTime,omitempty"`
	UpdatedTime  time.Time `json:"updatedTime,omitempty"`
}

type RecipeRepository interface {
	FetchRecipeByID(ID string) (*Recipe, error)
	FetchAllRecipes() ([]*Recipe, error)
	CreateRecipe(recipe *Recipe) (string, error)
	DeleteRecipe(ID string) error
	UpdateRecipe(ID string, recipe *Recipe) error
}
