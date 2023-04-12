package main

import (
	"techTestK-Style/database"
	"techTestK-Style/route"

	"github.com/gofiber/fiber/v2"
)

func main() {
    database.DatabaseInit()


    app := fiber.New()

	route.RouteInit(app)

	app.Listen(":8080")
}
