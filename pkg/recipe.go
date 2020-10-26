package recipe

type Recipe struct {
	ID    string `json:"ID"`
	Name  string `json:"name,omitempty"`
	Image string `json:"image,omitempty"`
	Text  string `json:"text,omitempty"`
}

type RecipeRepository interface {
	FetchRecipeByID(ID string) (*Recipe, error)
	FetchAllRecipes() ([]*Recipe, error)
}
