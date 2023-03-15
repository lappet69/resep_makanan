package foodrecipe

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lappet69/resep_makanan/config"
	"github.com/lappet69/resep_makanan/models"
	"gorm.io/gorm"
)

func ListFRecipe(c *gin.Context) {
	var FRecipe []models.FoodRecipe

	config.DB.Find(&FRecipe)
	c.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "success",
		"data":    &FRecipe})
}
func CreateFRecipe(c *gin.Context) {
	var request *models.FoodRecipe
	var category *models.Category

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"response_code":    "400",
			"response_message": "all field is required",
		})
		return
	}

	id := request.Id_Category
	checkIdCategory := config.DB.First(&category, id).Error

	if checkIdCategory != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"response_code":    "400",
			"response_message": "id category not found",
		})
		return
	}
	result := config.DB.Create(&request)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"status":  200,
		"message": "success",
	})
}

func GetById(c *gin.Context) {
	var FRecipe models.FoodRecipe
	id := c.Param("id")

	if err := config.DB.First(&FRecipe, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"status":  "400",
				"message": "Data tidak ditemukan",
			})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"status":  "400",
				"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "success",
		"data":    FRecipe,
	})
}

func Update(c *gin.Context) {
	var FRecipe models.FoodRecipe
	id := c.Param("id")

	if err := c.ShouldBindJSON(&FRecipe); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "400",
			"message": err.Error(),
		})
		return
	}

	if config.DB.Model(&FRecipe).Where("id = ?", id).Updates(&FRecipe).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "400",
			"message": "tidak dapat mengupdate FRecipe"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "Data berhasil diperbarui",
	})

}

func Delete(c *gin.Context) {

	var FRecipe models.FoodRecipe

	var input struct {
		Id json.Number
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "400",
			"message": err.Error()})
		return
	}

	id, _ := input.Id.Int64()
	if config.DB.Delete(&FRecipe, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "400",
			"message": "Tidak dapat menghapus FRecipe"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "200",
		"message": "Data berhasil dihapus"})
}
