package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Category_Name string `json:"category_name"`
}

type FoodRecipe struct {
	gorm.Model
	RecipeName  string   `json:"recipe_name"`
	Id_Category int      ` json:"id_category"` // foreign key to category
	Instruction []string ` json:"instruction"`
}
type FoodIngredient struct {
	gorm.Model
	Name string `json:"name"`
}

type RecipeDetail struct {
	gorm.Model
	Id_FoodRecipe     int    `json:"id_food_recipe"` // foreign key to food_recipe
	Id_Ingredient     int    `json:"id_ingredient"`  // foreign key to food_material
	Ingredient_Amount string `json:"ingredient_amount"`
}
