package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lappet69/resep_makanan/config"
	"github.com/lappet69/resep_makanan/controllers/category"
	fi "github.com/lappet69/resep_makanan/controllers/foodingredients"
	fRecipe "github.com/lappet69/resep_makanan/controllers/foodrecipe"
	recipeDetail "github.com/lappet69/resep_makanan/controllers/recipedetail"
)

func init() {
	config.LoadEnvVar()
	config.ConnectDB()
}

func main() {
	router := gin.Default()

	// category
	router.GET("/list-category", category.ListCategory)
	router.POST("/create-category", category.CreateCategory)
	router.GET("/category/:id", category.GetById)
	router.POST("/update-category/:id", category.Update)
	router.POST("/delete-category", category.Delete)

	// ingredients
	router.GET("/list-ingredients", fi.ListFIngredients)
	router.POST("/create-ingredients", fi.CreateFIngredients)
	router.GET("/ingredients/:id", fi.GetById)
	router.POST("/update-ingredients/:id", fi.Update)
	router.POST("/delete-ingredients", fi.Delete)

	// food recipe
	router.GET("/list-food-recipe", fRecipe.ListFRecipe)
	router.POST("/create-food-recipe", fRecipe.CreateFRecipe)
	router.GET("/food-recipe/:id", fRecipe.GetById)
	router.POST("/update-food-recipe/:id", fRecipe.Update)
	router.POST("/delete-food-recipe", fRecipe.Delete)

	//  recipe detail
	router.GET("/list-recipe-detail", recipeDetail.ListRecipeDetail)
	router.POST("/create-recipe-detail", recipeDetail.CreateRecipeDetail)
	router.GET("/recipe-detail/:id", recipeDetail.GetById)
	router.POST("/update-recipe-detail/:id", recipeDetail.Update)
	router.POST("/delete-recipe-detail", recipeDetail.Delete)

	router.Run()

}
