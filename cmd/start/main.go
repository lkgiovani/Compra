package main

import (
	"compra/internal/app/api/routes"
	"compra/internal/app/infra/config/configEnv"
	"compra/internal/app/infra/config/db"
	"fmt"
	"github.com/joho/godotenv"

	"log"

	"github.com/gofiber/fiber/v2"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// main is the entry point of the application. It initializes the database connection,
// sets up the Fiber web server, configures the routes, and starts listening on port 3000.
func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error em het .env")
	}

	initEnv := configEnv.NewConfig()

	// Initialize the database connection.
	dataBase := db.InitDB(initEnv)

	// Create a new Fiber app instance.
	app := fiber.New()

	// Set up the application routes.
	routes.SetupRoutes(app, dataBase)

	// Start the Fiber app and log any fatal errors.
	log.Fatal(app.Listen(fmt.Sprintf("%s:%s", initEnv.Server.Ip, initEnv.Server.Port)))
}
