package api

import (
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type Student struct {
	ID     	int    `json:"id"`
	Fname  	string `json:"firstname"`
	Lname 	string `json:"lastname"`
}

var students []Student = []Student{
	{ID: 1, Fname: "Mika", Lname: "Misono"},
	{ID: 2, Fname: "Miyako", Lname: "Tsukiyuki"},
}

func GetStudents(c *fiber.Ctx) error {
	return c.JSON(students)
}

func GetStudent(c *fiber.Ctx) error {
    id, err := strconv.Atoi(c.Params("id"))
    if err != nil {
        return c.SendStatus(fiber.StatusBadRequest)
    }
    for _, student := range students {
        if student.ID == id {
            return c.JSON(student)
        }
    }
    return c.SendStatus(fiber.StatusNotFound)
}

func NewStudent(c *fiber.Ctx) error {

	student := new(Student)
	if err := c.BodyParser(student); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	// Checke duplicate studetn
	for _, s := range students {
		if s.Fname == student.Fname && s.Lname == student.Lname {
			return c.Status(fiber.StatusConflict).SendString("Student already exit")
		}
	}
	student.ID = len(students) + 1
	students = append(students, *student)
	return c.JSON(student)
}

func GraduateStudent(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid ID")
	}
	for i, student := range students {
		if student.ID == id {
			students = append(students[:i], students[i+1:]...)
			return c.SendStatus(fiber.StatusOK)
		}
	}
	return c.Status(fiber.StatusNotFound).SendString("Student not found")
}

