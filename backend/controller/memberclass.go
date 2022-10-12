package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joon0000/sa-65/entity"
)

// POST /memberclass
func CreateMemberClass(c *gin.Context) {
	var memberclass entity.MemberClass
	if err := c.ShouldBindJSON(&memberclass); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := entity.DB().Create(&memberclass).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": memberclass})
}

// GET /memberclass/:id
func GetMemberClass(c *gin.Context) {
	var memberclass entity.MemberClass
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM memberclass WHERE id = ?", id).Scan(&memberclass).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": memberclass})
}

// GET /memberclass
func ListMemberClass(c *gin.Context) {
	var memberclass []entity.MemberClass
	if err := entity.DB().Raw("SELECT * FROM memberclass").Scan(&memberclass).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": memberclass})
}

// DELETE /memberclass/:id
func DeleteMemberClass(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM memberclass WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "memberclass not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /memberclass
func Updatememberclass(c *gin.Context) {
	var memberclass entity.MemberClass
	if err := c.ShouldBindJSON(&memberclass); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", memberclass.ID).First(&memberclass); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "memberclass not found"})
		return
	}

	if err := entity.DB().Save(&memberclass).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": memberclass})
}
