package controller

import (
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"github.com/joon0000/sa-65/entity"
)

// POST /users
func CreateUser(c *gin.Context) {
	var user entity.USER
	var employee entity.EMPLOYEE
	var memberclass entity.MemberClass
	var province entity.PROVINCE
	var role entity.ROLE

	// ผลลัพธ์ที่ได้จากขั้นตอนที่ 8 จะถูก bind เข้าตัวแปร user
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 9: ค้นหา employee ด้วย id
	if tx := entity.DB().Where("id = ?", user.EmpID).First(&employee); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "employee not found"})
		return
	}

	// 10: ค้นหา member ด้วย id
	if tx := entity.DB().Where("id = ?", user.MemberClassID).First(&memberclass); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "memberclass not found"})
		return
	}

	// 11: ค้นหา province ด้วย id
	if tx := entity.DB().Where("id = ?", user.ProvinceID).First(&province); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "province not found"})
		return
	}

	// 11: ค้นหา role ด้วย id
	if tx := entity.DB().Where("id = ?", user.RoleID).First(&role); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "role not found"})
		return
	}

	// 12: สร้าง user
	us := entity.USER{
		EMPLOYEE:    employee,
		MemberClass: memberclass,
		PROVINCE:    province,
		ROLE:        role,
	}

	// ขั้นตอนการ validate ที่นำมาจาก unit test
	if _, err := govalidator.ValidateStruct(us); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 13: บันทึก
	if err := entity.DB().Create(&us).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": us})
}

// GET /user/:id
func GetUser(c *gin.Context) {
	var user entity.USER
	id := c.Param("id")
	if err := entity.DB().Raw("SELECT * FROM user WHERE id = ?", id).Scan(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

// GET /users
func ListUser(c *gin.Context) {
	var users []entity.USER
	if err := entity.DB().Raw("SELECT * FROM user").Scan(&users).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": users})
}

// DELETE /users/:id
func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	if tx := entity.DB().Exec("DELETE FROM user WHERE id = ?", id); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": id})
}

// PATCH /users
func UpdateUser(c *gin.Context) {
	var user entity.USER
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if tx := entity.DB().Where("id = ?", user.ID).First(&user); tx.RowsAffected == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
		return
	}

	if err := entity.DB().Save(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}
