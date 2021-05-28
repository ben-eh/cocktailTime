package ingredient

import "github.com/ben-eh/cocktailTime/database"

type Ingredient struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func GetAllIngredients() []Ingredient {
	db := database.DBConnection()
	defer db.Close()

	results, err2 := db.Query("SELECT ingredient_id, name FROM ingredients")
	if err2 != nil {
		panic(err2.Error())
	}

	var ingredients []Ingredient

	for results.Next() {
		var ingredient Ingredient
		err := results.Scan(&ingredient.ID, &ingredient.Name)
		if err != nil {
			panic(err.Error())
		}

		ingredients = append(ingredients, ingredient)
	}

	return ingredients
}
