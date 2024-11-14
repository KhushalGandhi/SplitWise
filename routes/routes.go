package routes

import (
	"github.com/gofiber/fiber/v2"
	"splitwise/handlers"
	"splitwise/middleware"
)

func SetupRoutes(app *fiber.App) {
	// Account Routes
	app.Post("/api/accounts/register", handlers.RegisterAccount) // done
	app.Post("/api/accounts/login", handlers.Login)              // done

	// Group Routes (Protected)
	group := app.Group("/api/groups", middleware.JWTAuth)
	group.Post("/", handlers.CreateGroup)                // done
	group.Delete("/:id", handlers.DeleteGroup)           // Delete group route
	group.Post("/:id/add-user", handlers.AddUserToGroup) // Add user to group route  // done

	// Spend Routes (Protected)
	spend := app.Group("/api/spends", middleware.JWTAuth)
	spend.Post("/:id", handlers.CreateSpend) //done

	// Balance Routes (Protected)
	balance := app.Group("/api/balance", middleware.JWTAuth)
	balance.Get("/:group_id", handlers.GetRemainingBalance)
}
