package main

import (
	"fmt"
	"log"
	"project/api"
	"project/api_db"
	"github.com/gofiber/fiber/v2"
	fiberlog "github.com/gofiber/fiber/v2/middleware/logger"
	"gorm.io/gorm"

	"os"
	"gorm.io/driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	fiber_run()
}



// -------------------------------------------------------------------------------------------------------

var db *gorm.DB
func Connect() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	
	// fmt.Println("-----Connecte------")
	// fmt.Println("dbUser: ", dbUser)
	// fmt.Println("dbPassword: ", dbPassword)
	// fmt.Println("dbHost: ", dbHost)
	// fmt.Println("dbPort: ", dbPort)
	// fmt.Println("dbName: ", dbName)

	
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)

	// newLogger := logger.New(
	// 	log.New(os.Stdout, "\r\n", log.LstdFlags),
	// 	logger.Config{
	// 		SlowThreshold: time.Second,
	// 		LogLevel:      logger.Info, 
	// 		Colorful:      true,        
	// 	},
	// )

	// Connect to MySQL using GORM
	db, err = gorm.Open(mysql.Open(dsn))

	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	fmt.Println("Database connection established")
}

func fiber_run(){
	fiberApp := fiber.New()
	fiberApp.Use(fiberlog.New(fiberlog.Config{
		Format:     "[${time}] ${status} - ${method} ${path} ${latency}\n",
		TimeFormat: "2006-01-02 15:04:05",
		TimeZone:   "Local",
	}))

	Connect()
	
	db.AutoMigrate(&api_db.Student{})

	db.AutoMigrate(&api_db.Customer{})
	db.AutoMigrate(&api_db.Product{})
	db.AutoMigrate(&api_db.Order{})
	db.AutoMigrate(&api_db.OrderDetail{})
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

	fiberApp.Get("/charactor", func(c *fiber.Ctx) error{
		return api_db.FiberGetStudents(db, c)
	})
	fiberApp.Get("/charactor/:id", func(c *fiber.Ctx) error{
		return api_db.FiberGetStudent(db, c)
	})
	fiberApp.Post("/new-charactor", func(c *fiber.Ctx) error{
		return api_db.FiberNewStudent(db, c)
	})
	fiberApp.Delete("/finished-charactor", func(c *fiber.Ctx) error{
		return api_db.FiberGraduateStudent(db, c)
	})


	fmt.Println("Starting server...")

	if err := fiberApp.Listen(":8080"); err != nil {
		log.Fatalf("Fiber server failed to start: %v", err)
	}

}
