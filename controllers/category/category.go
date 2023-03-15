package category

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lappet69/resep_makanan/config"
	"github.com/lappet69/resep_makanan/models"
	"gorm.io/gorm"
)

func ListCategory(c *gin.Context) {
	var category []models.Category

	config.DB.Find(&category)
	c.JSON(http.StatusOK, gin.H{"categories": category})
}
func CreateCategory(c *gin.Context) {
	var request *models.Category

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
		"category": result,
	})
}

func GetById(c *gin.Context) {
	var category models.Category
	id := c.Param("id")

	if err := config.DB.First(&category, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data tidak ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"category": category})
}

func Update(c *gin.Context) {
	var category models.Category
	id := c.Param("id")

	if err := c.ShouldBindJSON(&category); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if config.DB.Model(&category).Where("id = ?", id).Updates(&category).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "tidak dapat mengupdate category"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil diperbarui"})

}

func Delete(c *gin.Context) {

	var category models.Category

	var input struct {
		Id json.Number
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	id, _ := input.Id.Int64()
	if config.DB.Delete(&category, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat menghapus category"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Data berhasil dihapus"})
}
