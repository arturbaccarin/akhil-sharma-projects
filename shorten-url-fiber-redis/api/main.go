package main

import (
	"log"
	"os"

	"github.com/arturbaccarin/akhil-sharma-projects/shorten-url-fiber-redis/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

func setupRoutes(app *fiber.App) {
	app.Get("/:url", routes.ResolveURL)
	app.Post("/api/v1", routes.ShortenURL)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	app := fiber.New()

	app.Use(logger.New())

	setupRoutes(app)

	log.Fatal(app.Listen(os.Getenv("APP_PORT")))
}

// https://youtu.be/3ExDEeSnyvE?list=PL5dTjWUk_cPbXTq9j-3Vaq08rjH1D5cTn
// 47:00
