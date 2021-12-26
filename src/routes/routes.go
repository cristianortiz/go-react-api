package routes

import (
	"go-react-api/src/controllers"
	"go-react-api/src/middlewares"

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

	//middleware to check user credentiales with jwt
	adminAuthenticated := admin.Use(middlewares.IsAuthenticated)
	adminAuthenticated.Get("/get-user", controllers.GetUser)
	adminAuthenticated.Post("/logout", controllers.Logout)
	adminAuthenticated.Put("/update-user", controllers.UpdateUserInfo)
	adminAuthenticated.Put("/update-password", controllers.UpdateUserPassword)
	//ambassadors endpoints
	adminAuthenticated.Get("/get-ambassadors", controllers.GetAmbassadors)
	//products endpints
	adminAuthenticated.Get("/get-products", controllers.GetProducts)
	adminAuthenticated.Post("/create-product", controllers.CreateProduct)
	adminAuthenticated.Get("/product/:id", controllers.GetProductByID)
}
