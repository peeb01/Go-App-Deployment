package main

import (
	"fmt"
	"log"
	"project/api"
	"project/api_db"

	"github.com/gofiber/fiber/v2"
	"github.com/gin-gonic/gin"
	// "github.com/gofiber/fiber/v2/middleware/cors"
	fiberlog "github.com/gofiber/fiber/v2/middleware/logger"
	"gorm.io/gorm"

	"os"
	"time"

	"gorm.io/driver/mysql"
	// "gorm.io/gorm"
	"gorm.io/gorm/logger"
	"github.com/joho/godotenv"
)

func main() {
	fiberApp := fiber.New()
	ginApp := gin.Default()
	
	Connect()

	fiberApp.Use(fiberlog.New(fiberlog.Config{
		Format:     "[${time}] ${status} - ${method} ${path} ${latency}\n",
		TimeFormat: "2006-01-02 15:04:05",
		TimeZone:   "Local",
	}))

	db.AutoMigrate(&api_db.Student{})

	db.AutoMigrate(&api_db.Customer{})
	db.AutoMigrate(&api_db.Product{})
	db.AutoMigrate(&api_db.Order{})
	db.AutoMigrate(&api_db.OrderDetail{})
	// db.AutoMigrate(&api_db.CreateOrderRequest{})
	// db.AutoMigrate(&api_db.OrderProduct{})
	
	fiberApp.Get("/", api.Helloworld)
	fiberApp.Get("/student", api.GetStudents)
	fiberApp.Get("/student/:id", api.GetStudent)
	fiberApp.Post("/new-student", api.NewStudent)
	fiberApp.Delete("/student-graduated/:id", api.GraduateStudent)

	fiberApp.Post("/register", func(c *fiber.Ctx) error {
		return api_db.RegisterCustomer(db, c)
	})
	fiberApp.Get("/cnsmr", func(c *fiber.Ctx) error {
		return api_db.GetCustomer(db, c)
	})
	fiberApp.Get("/product", func(c *fiber.Ctx) error {
		return api_db.GetProduct(db, c)
	})
	fiberApp.Get("/order", func(c *fiber.Ctx) error {
		return api_db.GetOrder(db, c)
	})
	fiberApp.Get("/order-detail", func(c *fiber.Ctx) error{
		return api_db.GetOrderDetail(db, c)
	})

	fiberApp.Post("/create-order", func(c *fiber.Ctx) error{
		return api_db.NewOrder(db, c)
	})

	// Set up Gin routes
	ginApp.GET("/gin/student", api_db.GetStudents(db))
	ginApp.GET("/gin/student/:id", api_db.GetStudent(db))
	ginApp.POST("/gin/new-student", api_db.NewStudent(db))
	ginApp.DELETE("/gin/student-graduated/:id", api_db.GraduateStudent(db))


	fmt.Println("Starting server")

	go func() {
		if err := fiberApp.Listen(":8080"); err != nil {
			log.Fatalf("Fiber server failed to start: %v", err)
		}
	}()


	go func() {
		if err := ginApp.Run(":8081"); err != nil {
			log.Fatalf("Gin server failed to start: %v", err)
		}
	}()

	select {}
}
var db *gorm.DB

func Connect() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	dbUser := os.Getenv("dev_db_user")	// get from variable, var name in gitlab
	dbPassword := os.Getenv("dev_db_password")
	dbHost := os.Getenv("dev_db_host")
	dbPort := os.Getenv("dev_db_port")
	dbName := os.Getenv("dev_db_name")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Asia%%2FBangkok", dbUser, dbPassword, dbHost, dbPort, dbName)

	// Set a custom logger for GORM
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second, 
			LogLevel:      logger.Info, 
			Colorful:      true,        
		},
	)

	// Connect to mysql
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	fmt.Println("Database connection established")
}