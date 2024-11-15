package routes

import (
	"github.com/gofiber/fiber/v2"
	"splitwise/handlers"
	"splitwise/middleware"
)

func SetupRoutes(app *fiber.App) {

	app.Use(middleware.Logging)

	// Account Routes
	app.Post("/api/accounts/register", handlers.RegisterAccount) // done
	app.Post("/api/accounts/login", handlers.Login)              // done

	// Group Routes (Protected)
	group := app.Group("/api/groups", middleware.JWTAuth)
	group.Post("/", handlers.CreateGroup)                // done
	group.Delete("/:id", handlers.DeleteGroup)           // done     // ye rh gya us test case ke sath      // Delete group route
	group.Post("/:id/add-user", handlers.AddUserToGroup) // Add user to group route  // done

	// Spend Routes (Protected)
	spend := app.Group("/api/spends", middleware.JWTAuth)
	spend.Post("/:id", handlers.CreateSpend)
	spend.Post("/payment/complete/:id", handlers.MarkSharePaidHandler)
	//done

	// Balance Routes (Protected)
	balance := app.Group("/api/balance")
	balance.Get("/:group_id", handlers.GetRemainingBalance) // done

	balance.Get(":group_id/user", middleware.JWTAuth, handlers.GetRemainingBalanceforUser) // done

}
