package recipedetail

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lappet69/resep_makanan/config"
	"github.com/lappet69/resep_makanan/models"
	"gorm.io/gorm"
)

func ListRecipeDetail(c *gin.Context) {
	var rcpDetail []models.RecipeDetail

	config.DB.Find(&rcpDetail)
	c.JSON(http.StatusOK, gin.H{"categories": rcpDetail})
}
func CreateRecipeDetail(c *gin.Context) {
	var request *models.RecipeDetail

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
		"rcpDetail": result,
	})
}

func GetById(c *gin.Context) {
	var rcpDetail models.RecipeDetail
	id := c.Param("id")

	if err := config.DB.First(&rcpDetail, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"rcpDetail": rcpDetail})
}

func Update(c *gin.Context) {
	var rcpDetail models.RecipeDetail
	id := c.Param("id")

	if err := c.ShouldBindJSON(&rcpDetail); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if config.DB.Model(&rcpDetail).Where("id = ?", id).Updates(&rcpDetail).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "tidak dapat mengupdate rcpDetail"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil diperbarui"})

}

func Delete(c *gin.Context) {

	var rcpDetail models.RecipeDetail

	var input struct {
		Id json.Number
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	id, _ := input.Id.Int64()
	if config.DB.Delete(&rcpDetail, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat menghapus rcpDetail"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil dihapus"})
}
