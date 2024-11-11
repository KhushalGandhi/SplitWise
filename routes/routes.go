package routes

import (
	"github.com/gofiber/fiber/v2"
	"splitwise/handlers"
	"splitwise/middleware"
)

func SetupRoutes(app *fiber.App) {
	// Account Routes

	app.Post("/api/accounts/register", handlers.RegisterAccount)
	app.Post("/api/accounts/login", handlers.Login)

	// Group Routes (Protected)

	group := app.Group("/api/groups", middleware.JWTAuth)
	group.Post("/", handlers.CreateGroup)
	//group.Delete("/:id", handlers.)
	//group.Post("/:id/add-user", handlers)

	// Spend Routes (Protected)

	spend := app.Group("/api/spends", middleware.JWTAuth)
	spend.Post("/", handlers.CreateSpend)

	// Balance Routes (Protected)

	balance := app.Group("/api/balance", middleware.JWTAuth)
	balance.Get("/:group_id", handlers.GetRemainingBalance)
}
