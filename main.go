package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/lambda-platform/ebarimt-rest-api/router"
	"log"
)

// @title_en            EBarimt PosAPI REST API using Go and Fiber
// @title            EBarimt PosAPI REST API Go ба Fiber ашигласан
// @version             1.0
// @description_en      This is a REST API implementation for EBarimt PosAPI, built using the Fiber framework in Go, showcasing the creation of a fast and efficient RESTful API.
// @description      Энэ жинээ нь Go хэл дээр Fiber framework-ийг ашиглаж EBarimt PosAPI-ийн REST API-г хурдан ба хэмнэлттэй бүтээсэн жишээ.
// @termsOfService      http://swagger.io/terms/
// @contact.name        Munkh-Altai Chuluunbaatar
// @contact.email       munkaltai@gmail.com
// @license.name        MIT License
// @license.url         https://opensource.org/licenses/MIT
// @github.repo         https://github.com/lambda-platform/ebarimt-rest-api
// @BasePath            /
func main() {
	app := fiber.New()

	// Create the Swagger JSON file

	app.Static("/", "docs")

	app.Get("/en/*", swagger.New(swagger.Config{ // custom
		URL:          "/en/swagger.json",
		DeepLinking:  false,
		ValidatorUrl: "127.0.0.1", // Disable validator
	}))
	app.Get("/mn/*", swagger.New(swagger.Config{ // custom
		URL:          "/mn/swagger.json",
		DeepLinking:  false,
		ValidatorUrl: "127.0.0.1", // Disable validator
	}))

	router.Routes(app)

	log.Fatal(app.Listen(":3388"))
}
