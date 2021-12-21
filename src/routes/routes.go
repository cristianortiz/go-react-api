package routes

import (
	"go-react-api/src/controllers"

	"github.com/gofiber/fiber/v2"
)

//Setup function defines groups por every module in app "api" or 'admin'
func Setup(app *fiber.App) {
	//url prefix for api module routes
	api := app.Group("api")
	//url prefix for admin module inside api prefix
	admin := api.Group("admin")
	//this complete route is /api/admin/register
	admin.Post("/register", controllers.Register)
	admin.Post("/login", controllers.Login)
	admin.Post("/logout", controllers.Logout)
	admin.Get("/auth-user", controllers.AuthenticatedUser)

}
