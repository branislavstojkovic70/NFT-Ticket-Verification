package route

import (
	"net/http"

	"github.com/branislavstojkovic70/nft-ticket-verification/api/controller"
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(router *gin.Engine, userController *controller.UserController) {
	userRoutes := router.Group("/user")
	{
		userRoutes.GET("/", userController.GetAllUsers)
		userRoutes.GET("/:id", userController.GetUserByID)
		userRoutes.POST("/", userController.CreateUser)
		userRoutes.PUT("/:id", userController.UpdateUser)
		userRoutes.DELETE("/:id", userController.DeleteUser)
	}
}

func InitRoutes(server *gin.Engine) {

	// Group user related routes together
	userRoutes := server.Group("/user")
	{
		userRoutes.GET("/", func(c *gin.Context) {
			c.String(http.StatusOK, "Hello from HTTP server")
		})
		// // Handle the GET requests at /u/login
		// // Show the login page
		// // Ensure that the user is not logged in by using the middleware
		// userRoutes.GET("/login", ensureNotLoggedIn(), showLoginPage)

		// // Handle POST requests at /u/login
		// // Ensure that the user is not logged in by using the middleware
		// userRoutes.POST("/login", ensureNotLoggedIn(), performLogin)

		// // Handle GET requests at /u/logout
		// // Ensure that the user is logged in by using the middleware
		// userRoutes.GET("/logout", ensureLoggedIn(), logout)

		// // Handle the GET requests at /u/register
		// // Show the registration page
		// // Ensure that the user is not logged in by using the middleware
		// userRoutes.GET("/register", ensureNotLoggedIn(), showRegistrationPage)

		// // Handle POST requests at /u/register
		// // Ensure that the user is not logged in by using the middleware
		// userRoutes.POST("/register", ensureNotLoggedIn(), register)
	}
}
