package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

// @Summary Kursları listele
// @Tags courses
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /courses [get]
func GetCourses(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Kurslar listelendi"})
}

// @Summary Kurs detaylarını getir
// @Tags courses
// @Produce json
// @Param id path string true "Kurs ID"
// @Success 200 {object} map[string]interface{}
// @Router /courses/{id} [get]
func GetCourse(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Kurs detayı"})
}

// @Summary Yeni kurs oluştur
// @Tags courses
// @Accept json
// @Produce json
// @Success 201 {object} map[string]interface{}
// @Router /courses [post]
func CreateCourse(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"message": "Kurs oluşturuldu"})
}

// @Summary Kurs güncelle
// @Tags courses
// @Param id path string true "Kurs ID"
// @Router /courses/{id} [put]
func UpdateCourse(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Kurs güncellendi"})
}

// @Summary Kurs sil
// @Tags courses
// @Param id path string true "Kurs ID"
// @Router /courses/{id} [delete]
func DeleteCourse(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "Kurs silindi"})
}