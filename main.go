package main

import (
	"log"

	"test/app"
	routego "test/route.go"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func main() {
	// Load env
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("error loading .env file!")
	}

	// Init DB (HANYA SEKALI)
	db := app.Dbconncentio()

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	if err := sqlDB.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Println("Database connected successfully!")

	// Init Echo
	e := echo.New()

	// Middleware (debug & recover)
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// 🔥 PASANG VALIDATOR (INI YANG KEMARIN KURANG)
	e.Validator = &CustomValidator{
		validator: validator.New(),
	}

	// Register routes (inject DB)
	routego.UserRoutes("/api", e)

	log.Println("Server running on :8080")
	if err := e.Start(":8080"); err != nil {
		log.Fatal(err)
	}
}
