package data

import (
	recipe "github.com/mikeletux/home-recipes/pkg"
)

// SampleRecipes is a sample dataset for testing purposes
// Recipes are not well explained to be tasty, so do not try at home :P
var SampleRecipes = map[string]*recipe.Recipe{
	"01D3XZ3ZHCP3KG9VT4FGAD8KDR": &recipe.Recipe{
		ID:          "01D3XZ3ZHCP3KG9VT4FGAD8KDR",
		Name:        "Lentils with spanish sausage",
		Image:       "https://www.hogarmania.com/archivos/201702/receta-lentejas-chorizo-1280x720x80xX.jpg",
		Ingredients: []string{"lentils", "spanish sausage", "onion", "garlic", "olive oil", "paprika", "vegetables soup dice"},
		Text:        "We put everything in a pressure pot",
	},
	"01D3XZ7CN92AKS9HAPSZ4D5DP9": &recipe.Recipe{
		ID:          "01D3XZ7CN92AKS9HAPSZ4D5DP9",
		Name:        "Sausages with white wine",
		Image:       "https://recetasdecocina.elmundo.es/wp-content/uploads/2013/04/receta-salchichas-al-vino-blanco.jpg",
		Ingredients: []string{"fresh sausages", "white wine", "onion", "garlic", "meat soup dice"},
		Text:        "We put everything in a pan and reduce the wine. Uhmm so tasty!",
	},
	"01D3XZ89NFJZ9QT2DHVD462AC2": &recipe.Recipe{
		ID:          "01D3XZ89NFJZ9QT2DHVD462AC2",
		Name:        "Homemade pancakes",
		Image:       "https://estaticos.miarevista.es/media/cache/760x570_thumb/uploads/images/recipe/554207ab0a73fed41f01dd26/tortitas-int.jpg",
		Ingredients: []string{"milk", "sugar", "flour", "eggs", "vanilla extract", "yeast"},
		Text:        "Mix everything and put it in a pan. Get fat!",
	},
}
