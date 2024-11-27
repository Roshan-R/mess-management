package main

import (
	//"encoding/json"
	"log"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/joho/godotenv"
	"mestos.backend/database"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database.CreateDatabaseConnection()

	app := fiber.New()
	app.Use(cors.New())

	app.Get("/getTodaysMenu", func(c fiber.Ctx) error {
		menu, err := database.GetTodaysMenu()
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, err.Error())
		}
		return c.JSON(menu)
	})

	app.Post("/takeMessout", func(c fiber.Ctx) error {
		payload := struct {
			Date time.Time `json:"date"`
		}{}
		if err := c.Bind().Body(&payload); err != nil {
			return err
		}
		database.TakeMessout(payload.Date)
		return nil
	})

	log.Fatal(app.Listen(":3000"))
}
