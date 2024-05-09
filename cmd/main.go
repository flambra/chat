package main

import (
	"fmt"
	"log"
	"os"

	"github.com/flambra/chat/database"
	"github.com/flambra/chat/internal"
	"github.com/flambra/chat/internal/watcher"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	err := database.New()
	if err != nil {
		log.Fatal(err)
	}
	defer database.Disconnect()

	go watcher.Init()

	app := fiber.New()

	fiber.SetParserDecoder(fiber.ParserConfig{
		IgnoreUnknownKeys: true,
		ZeroEmpty:         true,
	})

	internal.InitializeRoutes(app)

	port := os.Getenv("SERVER_PORT")
	if len(port) == 0 {
		port = "8080"
	}

	/* Start Server */
	err = app.Listen(fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatal(err)
	}
}
