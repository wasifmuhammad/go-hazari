package controller

import (
	"golang-crud-gin/config"
	"golang-crud-gin/models"
	"net/http"
	"time" // Import the time package

	"github.com/gin-gonic/gin"
)

// CreateUser creates a new user in the database
func CreateUser(name, email, password string) (*models.User, error) {
	// Create a new user instance
	newUser := models.User{Name: name, Email: email, Password: password}

	// Save the new user to the database
	if err := config.DB.Create(&newUser).Error; err != nil {
		return nil, err
	}

	return &newUser, nil
}

// UserController retrieves all users from the database
func UserController(c *gin.Context) {
	users := []models.User{}
	config.DB.Find(&users)
	c.JSON(http.StatusOK, users)
}
func CheckIn(c *gin.Context) {
    // Bind JSON data from request body to input struct
    var inputData struct {
        UserID        uint   `json:"user_id"`
    }
    if err := c.ShouldBindJSON(&inputData); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

	 // Get the last record for the user
    var lastRecord models.InOut
    if err := config.DB.Where("user_id = ?", inputData.UserID).Order("ID desc").First(&lastRecord).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query records"})
        return
    }

    // If the latest command is "Check out", return error
    if lastRecord.LatestCommand == "Check in" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "You are already checked-In"})
        return
    }


	 if lastRecord.LatestCommand == "Sign in" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "You are already signed-In"})
        return
    }

    // Create a new InOut instance
    newCheckIn := models.InOut{
        UserID:        inputData.UserID,
        LatestCommand: "Check in",
  		CreatedAt:     time.Now().Format(time.RFC3339), // Current time
    }

    // Save the new check-in record to the database
    if err := config.DB.Create(&newCheckIn).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create check-in record"})
        return
    }

    c.JSON(http.StatusCreated, newCheckIn) // Return the newly created check-in record
}



func SignIn(c *gin.Context) {
    // Bind JSON data from request body to input struct
    var inputData struct {
        UserID        uint   `json:"user_id"`
    }
    if err := c.ShouldBindJSON(&inputData); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

	 // Get the last record for the user
    var lastRecord models.InOut
    if err := config.DB.Where("user_id = ?", inputData.UserID).Order("ID desc").First(&lastRecord).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query records"})
        return
    }

    // If the latest command is "Check out", return error
    if lastRecord.LatestCommand == "Check in" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "You are already checked-In"})
        return
    }


	  if lastRecord.LatestCommand == "Sign in" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "You are already Signed-In"})
        return
    }

    // Create a new InOut instance
    newCheckIn := models.InOut{
        UserID:        inputData.UserID,
  		LatestCommand: "Sign in",
  		CreatedAt:     time.Now().Format(time.RFC3339), // Current time
    }

    // Save the new check-in record to the database
    if err := config.DB.Create(&newCheckIn).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create check-in record"})
        return
    }

    c.JSON(http.StatusCreated, newCheckIn) // Return the newly created check-in record
}

func Break(c *gin.Context) {
    // Bind JSON data from request body to input struct
    var inputData struct {
        UserID        uint   `json:"user_id"`
    
    }
    if err := c.ShouldBindJSON(&inputData); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

	 // Get the last record for the user
    var lastRecord models.InOut
    if err := config.DB.Where("user_id = ?", inputData.UserID).Order("ID desc").First(&lastRecord).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query records"})
        return
    }

    // If the latest command is "Check out", return error
    if lastRecord.LatestCommand == "Check out" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "You are not checked-In"})
        return
    }


	 if lastRecord.LatestCommand == "Sign out" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "You are not signed-In"})
        return
    }



	 if lastRecord.LatestCommand == "Break" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "You are already on break"})
        return
    }

    // Create a new InOut instance
    newCheckIn := models.InOut{
        UserID:        inputData.UserID,
        LatestCommand: "Break",
  		CreatedAt:     time.Now().Format(time.RFC3339), // Current time
    }

    // Save the new check-in record to the database
    if err := config.DB.Create(&newCheckIn).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create check-in record"})
        return
    }

    c.JSON(http.StatusCreated, newCheckIn) // Return the newly created check-in record
}
func Checkout(c *gin.Context) {
    // Bind JSON data from request body to input struct
    var inputData struct {
        UserID uint `json:"user_id"`
    }
    if err := c.ShouldBindJSON(&inputData); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Get the last record for the user
    var lastRecord models.InOut
    if err := config.DB.Where("user_id = ?", inputData.UserID).Order("ID desc").First(&lastRecord).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query records"})
        return
    }

    // If the latest command is "Check out", return error
    if lastRecord.LatestCommand == "Check out" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "You must check in first"})
        return
    }


	 if lastRecord.LatestCommand == "Sign in" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "you need to sign off"})
        return
    }
    // Create a new checkout record
    newCheckout := models.InOut{
        UserID:        inputData.UserID,
        LatestCommand: "Check out",
        CreatedAt:     time.Now().Format(time.RFC3339), // Current time
    }

    // Save the new checkout record to the database
    if err := config.DB.Create(&newCheckout).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create checkout record"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Checkout successful"})
}



func SignOut(c *gin.Context) {
    // Bind JSON data from request body to input struct
    var inputData struct {
        UserID uint `json:"user_id"`
    }
    if err := c.ShouldBindJSON(&inputData); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Get the last record for the user
    var lastRecord models.InOut
    if err := config.DB.Where("user_id = ?", inputData.UserID).Order("ID desc").First(&lastRecord).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query records"})
        return
    }

    // If the latest command is "Check out", return error
    if lastRecord.LatestCommand == "Check out" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "You must check in first"})
        return
    }


	 if lastRecord.LatestCommand == "Sign out" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "you need to sign in first"})
        return
    }


	 if lastRecord.LatestCommand == "Check in" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "you need to check out"})
        return
    }
    // Create a new checkout record
    newCheckout := models.InOut{
        UserID:        inputData.UserID,
        LatestCommand: "Sign out",
        CreatedAt:     time.Now().Format(time.RFC3339), // Current time
    }

    // Save the new checkout record to the database
    if err := config.DB.Create(&newCheckout).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create checkout record"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Signed Off successful"})
}


func Back(c *gin.Context) {
    // Bind JSON data from request body to input struct
    var inputData struct {
        UserID uint `json:"user_id"`
    }
    if err := c.ShouldBindJSON(&inputData); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Get the last record for the user
    var lastRecord models.InOut
    if err := config.DB.Where("user_id = ?", inputData.UserID).Order("ID desc").First(&lastRecord).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query records"})
        return
    }

    // If the latest command is "Check out", return error
    if lastRecord.LatestCommand == "Check out" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "You must check in first"})
        return
    }


	 if lastRecord.LatestCommand == "Sign out" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "you need to sign in first"})
        return
    }


	 if lastRecord.LatestCommand != "Break" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "You are not on break"})
        return
    }



    // Create a new checkout record
    newCheckout := models.InOut{
        UserID:        inputData.UserID,
        LatestCommand: "Back",
        CreatedAt:     time.Now().Format(time.RFC3339), // Current time
    }

    // Save the new checkout record to the database
    if err := config.DB.Create(&newCheckout).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create checkout record"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Break ended successful"})
}