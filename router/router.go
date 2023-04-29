package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lambda-platform/ebarimt-rest-api/controller"
)

func Routes(app *fiber.App) {

	app.Get("/info", controller.Info)

	app.Get("/check", controller.Check)

	app.Get("/call", controller.Call)

	app.Post("/put", controller.Put)

	app.Post("/put-batch", controller.PutBatch)

	app.Post("/return", controller.Return)

	app.Get("/send", controller.Send)

	app.Get("/organization/:register", controller.OrganizationInfo)
}
