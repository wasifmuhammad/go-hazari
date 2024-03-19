package routes

import (
	"golang-crud-gin/controller"

	"github.com/gin-gonic/gin"
)

// UserRoute sets up routes related to user functionalities
func UserRoute(router *gin.Engine) {
    // Route to get all users
    router.GET("/", controller.UserController)

    // Route to create a new user
    router.POST("/createUser", CreateUserHandler)


	// Route to handle check-in requests
    router.POST("/checkIn", controller.CheckIn)


	// Route to handle check-in requests
    router.POST("/checkOut", controller.Checkout)


		// Route to handle check-in requests
    router.POST("/signIn", controller.SignIn)


    router.POST("/signOut", controller.SignOut)


	   router.POST("/break", controller.Break)


    router.POST("/back", controller.Back)

}

// CreateUserHandler handles HTTP requests for creating a new user
func CreateUserHandler(c *gin.Context) {
    // Call the CreateUser function from the controller package
    newUser, err := controller.CreateUser("Muhammad Wasif", "wasif@yopmail.com", "123123")
    if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }

    c.JSON(201, newUser) // Return the newly created user
}
