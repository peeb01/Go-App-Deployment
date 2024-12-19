package api_db

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Fname  string `json:"firstname"`
	Lname  string `json:"lastname"`
}

func GetStudents(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var students []Student
		if result := db.Find(&students); result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to fetch students",
			})
			return
		}
		c.JSON(http.StatusOK, students)
	}
}

func GetStudent(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var student Student
		if result := db.First(&student, id); result.Error != nil {
			if result.Error == gorm.ErrRecordNotFound {
				c.JSON(http.StatusNotFound, gin.H{
					"error": "Student not found",
				})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to fetch student",
			})
			return
		}
		c.JSON(http.StatusOK, student)
	}
}

func NewStudent(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var students []Student

		if err := c.ShouldBindJSON(&students); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Failed to parse request body",
			})
			return
		}

		for _, student := range students {
			if result := db.Create(&student); result.Error != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "Failed to create student",
				})
				return
			}
		}

		c.JSON(http.StatusCreated, gin.H{
			"message":  "Students added successfully",
			"students": students,
		})
	}
}

func GraduateStudent(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")

		if result := db.Delete(&Student{}, id); result.Error != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to graduate student",
			})
			return
		}

		c.String(http.StatusOK, "Student Graduated.")
	}
}


// var students []Student = []Student{
// 	{ID: 1, Fname: "Mika", Lname: "Misono"},
// 	{ID: 2, Fname: "Miyako", Lname: "Tsukiyuki"},
// }

// func SeedStudents(db *gorm.DB) {
// 	var count int64
// 	db.Model(&Student{}).Count(&count)
// 	if count == 0 {
// 		if err := db.Create(&students).Error; err != nil {
// 			fmt.Println("Failed to seed students:", err)
// 		} else {
// 			fmt.Println("Successfully seeded students.")
// 		}
// 	} else {
// 		fmt.Println("Students are already seeded.")
// 	}
// }