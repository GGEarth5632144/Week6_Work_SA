package librarianmembers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"example.com/sa-68-example/entity"
)

// Struct ที่ใช้เก็บข้อมูล LibrarianMember โดยใช้ GORM
var db *gorm.DB // กำหนด db เพื่อใช้งาน GORM

func InitDatabase(database *gorm.DB) {
	db = database
	// อัปเดตฐานข้อมูล
	db.AutoMigrate(&entity.LibrarianMember{})
}

// Create function สำหรับเพิ่มข้อมูล LibrarianMember
func Create(c *gin.Context) {
	var librarianMember entity.LibrarianMember

	// รับข้อมูลจาก request body
	if err := c.ShouldBindJSON(&librarianMember); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// บันทึกข้อมูลลงในฐานข้อมูล
	if result := db.Create(&librarianMember); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "LibrarianMember created successfully", "data": librarianMember})
}

// GetAll function สำหรับดึงข้อมูลของ librarianmembers ทั้งหมด
func GetAll(c *gin.Context) {
	var librarianMembers []entity.LibrarianMember

	// ดึงข้อมูลทั้งหมดจากฐานข้อมูล
	if result := db.Find(&librarianMembers); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, librarianMembers)
}

// Get function สำหรับดึงข้อมูลของ librarianmember ที่มี MemberID
func Get(c *gin.Context) {
	id := c.Param("id")
	var librarianMember entity.LibrarianMember

	// หาข้อมูล librarianmember ที่ตรงกับ ID ที่ส่งมา
	if result := db.First(&librarianMember, id); result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "LibrarianMember not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, librarianMember)
}

// Update function สำหรับอัปเดตข้อมูล librarianmember
func Update(c *gin.Context) {
	id := c.Param("id")
	var librarianMember entity.LibrarianMember

	// หาข้อมูล librarianmember ที่ตรงกับ ID ที่ส่งมา
	if result := db.First(&librarianMember, id); result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "LibrarianMember not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	// รับข้อมูลใหม่ที่ต้องการอัปเดต
	if err := c.ShouldBindJSON(&librarianMember); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// อัปเดตข้อมูลในฐานข้อมูล
	if result := db.Save(&librarianMember); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "LibrarianMember updated successfully", "data": librarianMember})
}

// Delete function สำหรับลบ librarianmember ที่มี MemberID
func Delete(c *gin.Context) {
	id := c.Param("id")
	var librarianMember entity.LibrarianMember

	// หาข้อมูล librarianmember ที่ตรงกับ ID ที่ส่งมา
	if result := db.First(&librarianMember, id); result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "LibrarianMember not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	// ลบข้อมูลออกจากฐานข้อมูล
	if result := db.Delete(&librarianMember); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "LibrarianMember deleted successfully"})
}
