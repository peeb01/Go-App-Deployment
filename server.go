package main

import (
	"fmt"
	"log"
	"project/api"
	"project/api_db"

	"github.com/gofiber/fiber/v2"
	// "github.com/gofiber/fiber/v2/middleware/cors"
	// "gorm.io/gorm"

	"os"
	// "time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	// "gorm.io/gorm/logger"
	"github.com/joho/godotenv"
)

func main() {
	app := fiber.New()
	
	Connect()

	db.AutoMigrate(&api_db.Student{})

	db.AutoMigrate(&api_db.Customer{})
	db.AutoMigrate(&api_db.Product{})
	db.AutoMigrate(&api_db.Order{})
	db.AutoMigrate(&api_db.OrderDetail{})
	// db.AutoMigrate(&api_db.CreateOrderRequest{})
	// db.AutoMigrate(&api_db.OrderProduct{})
	

	app.Get("/", api.Helloworld)
	app.Get("/student", api.GetStudents)
	app.Get("/student/:id", api.GetStudent)
	app.Post("/new-student", api.NewStudent)
	app.Delete("/student-graduated/:id", api.GraduateStudent)

	app.Get("/charactor", func(c *fiber.Ctx) error {
		return api_db.GetStudents(db, c)
	})
	app.Get("/charactor/:id", func(c *fiber.Ctx) error {
		return api_db.GetStudent(db, c)
	})
	app.Post("/new-char", func(c *fiber.Ctx) error {
		return api_db.NewStudent(db, c)
	})
	app.Delete("/rm-charactors/:id", func(c *fiber.Ctx) error {
		return api_db.GraduateStudent(db, c)
	})

	app.Post("/register", func(c *fiber.Ctx) error {
		return api_db.RegisterCustomer(db, c)
	})
	app.Get("/cnsmr", func(c *fiber.Ctx) error {
		return api_db.GetCustomer(db, c)
	})
	app.Get("/product", func(c *fiber.Ctx) error {
		return api_db.GetProduct(db, c)
	})
	app.Get("/order", func(c *fiber.Ctx) error {
		return api_db.GetOrder(db, c)
	})
	app.Get("/order-detail", func(c *fiber.Ctx) error{
		return api_db.GetOrderDetail(db, c)
	})

	app.Post("/create-order", func(c *fiber.Ctx) error{
		return api_db.NewOrder(db, c)
	})

	fmt.Println("Starting server")

	if err := app.Listen(":8080"); err != nil {
		log.Fatal(err)
	}
}

var db *gorm.DB

func Connect() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Asia%%2FBangkok", dbUser, dbPassword, dbHost, dbPort, dbName)

	// // Set a custom logger for GORM
	// newLogger := logger.New(
	// 	log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
	// 	logger.Config{
	// 		SlowThreshold: time.Second, // Slow SQL threshold
	// 		LogLevel:      logger.Info, // Log level
	// 		Colorful:      true,        // Enable color
	// 	},
	// )

	// Connect to mysql
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		// Logger: newLogger,
	})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	fmt.Println("Database connection established")
}