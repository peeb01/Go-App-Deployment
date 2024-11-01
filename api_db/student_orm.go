package api_db

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	// "fmt"
)

type Student struct {
	gorm.Model
	ID     int    `json:"id"`
	Fname  string `json:"firstname"`
	Lname  string `json:"lastname"`
}

func GetStudents(db *gorm.DB, c *fiber.Ctx) error {
	var students []Student
	if result := db.Find(&students); result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch students",
		})
	}
	return c.Status(fiber.StatusOK).JSON(students)
}

func GetStudent(db *gorm.DB, c *fiber.Ctx) error {
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

func NewStudent(db *gorm.DB, c *fiber.Ctx) error {
    var students []Student

    if err := c.BodyParser(&students); err != nil {
        return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
            "error": "Failed to parse request body",
        })
    }

    for _, student := range students {
        if result := db.Create(&student); result.Error != nil {
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
                "error": "Failed to create student",
            })
        }
    }

    return c.Status(fiber.StatusCreated).JSON(fiber.Map{
        "message":  "Students added successfully",
        "students": students,
    })
}

func GraduateStudent(db *gorm.DB, c *fiber.Ctx) error {
	id := c.Params("id")

	if result := db.Delete(&Student{}, id); result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to graduate student",
		})
	}

	return c.Status(fiber.StatusOK).SendString("Student Graduated.")
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