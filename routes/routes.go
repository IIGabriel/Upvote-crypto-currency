package routes

import "github.com/gofiber/fiber/v2"

func AllRoutes() *fiber.App {
	app := fiber.New()

	// Votes
	app.Post("/upvote/:coin", CreateUpVote)
	app.Post("/downvote/:coin", CreateDownVote)

	// Currency
	app.Post("/currency", CreateCurrency)
	app.Get("/currency/:coin", GetCurrency)
	app.Put("/currency/:coin", EditCurrency)
	app.Delete("/currency/:coin", DeleteCurrency)

	return app
}
