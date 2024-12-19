package api_db

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	// "fmt"
)

// type Student struct {
// 	gorm.Model
// 	ID     int    `json:"id"`
// 	Fname  string `json:"firstname"`
// 	Lname  string `json:"lastname"`
// }

func FiberGetStudents(db *gorm.DB, c *fiber.Ctx) error {
	var students []Student
	if result := db.Find(&students); result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch students",
		})
	}
	return c.Status(fiber.StatusOK).JSON(students)
}

func FiberGetStudent(db *gorm.DB, c *fiber.Ctx) error {
	id := c.Params("id")
	var student Student
	if result := db.First(&student, id); result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Student not found",
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch student",
		})
	}
	return c.Status(fiber.StatusOK).JSON(student)
}

func FiberNewStudent(db *gorm.DB, c *fiber.Ctx) error {

    var student Student

    if err := c.BodyParser(&student); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Failed to parse request body",
        })
    }
	var existingStudent Student
    if err := db.Where("fname = ? AND lname = ?", student.Fname, student.Lname).First(&existingStudent).Error; err == nil {
        return c.Status(fiber.StatusConflict).JSON(fiber.Map{
            "error": "Student already exists",
        })
    }
    if result := db.Create(&student); result.Error != nil {
        return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
            "error": "Failed to create student",
        })
    }

    return c.Status(fiber.StatusCreated).JSON(fiber.Map{
        "message": "Student added successfully",
        "student": student,
    })
}


func FiberGraduateStudent(db *gorm.DB, c *fiber.Ctx) error {
	id := c.Params("id")

	if result := db.Delete(&Student{}, id); result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to graduate student",
		})
	}

	return c.Status(fiber.StatusOK).SendString("Student Graduated.")
}
