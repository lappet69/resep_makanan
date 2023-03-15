package foodingredients

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lappet69/resep_makanan/config"
	"github.com/lappet69/resep_makanan/models"
	"gorm.io/gorm"
)

func ListFIngredients(c *gin.Context) {
	var FIngredients []models.FoodIngredient

	config.DB.Find(&FIngredients)
	c.JSON(http.StatusOK, gin.H{"categories": FIngredients})
}
func CreateFIngredients(c *gin.Context) {
	var request *models.FoodIngredient

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"response_code":    "400",
			"response_message": "payload invalid",
		})
	}
	result := config.DB.Create(&request)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"FIngredients": result,
	})
}

func GetById(c *gin.Context) {
	var FIngredients models.FoodIngredient
	id := c.Param("id")

	if err := config.DB.First(&FIngredients, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"FIngredients": FIngredients})
}

func Update(c *gin.Context) {
	var FIngredients models.FoodIngredient
	id := c.Param("id")

	if err := c.ShouldBindJSON(&FIngredients); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if config.DB.Model(&FIngredients).Where("id = ?", id).Updates(&FIngredients).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "tidak dapat mengupdate FIngredients"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil diperbarui"})

}

func Delete(c *gin.Context) {

	var FIngredients models.FoodIngredient

	var input struct {
		Id json.Number
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	id, _ := input.Id.Int64()
	if config.DB.Delete(&FIngredients, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat menghapus FIngredients"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil dihapus"})
}
