package controllers

import (
	"net/http"
	"praktikum/config"
	"praktikum/helpers"
	"praktikum/middleware"
	"praktikum/models"
	"strconv"

	"github.com/labstack/echo/v4"
)

// get all users
func GetUsersController(e echo.Context) error {
	var userDB []models.User
	if err := config.DB.Find(&userDB).Error; err != nil {
	  return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	var users []models.UserResponse
	for _, u := range userDB {
		userRes := models.UserResponse{
			Id:        u.ID,
			Name:      u.Name,
			Email:     u.Email,
		}
		users = append(users, userRes)
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
	  "message": "success get all users",
	  "users":   users,
	})
}
  
  
  // get user by id
func GetUserController(e echo.Context) error {
	// your solution here
	searchID := e.Param("id")
	var userDB []models.User
	if err := config.DB.First(&userDB, searchID).Error; err != nil {
	  // User not found
	  return e.JSON(http.StatusNotFound, map[string]interface{}{
		"message": "User not found",
	  })
	}
	
	var users []models.UserResponse
	for _, u := range userDB {
		userRes := models.UserResponse{
			Id:        u.ID,
			Name:      u.Name,
			Email:     u.Email,
		}
		users = append(users, userRes)
	}
	// User found
	return e.JSON(http.StatusOK, map[string]interface{}{
	  "message": "Successfully get user by ID",
	  "user": users,
	})
}
  
  
  // create new user
func CreateUserController(e echo.Context) error {
	user := models.User{}

	e.Bind(&user)

	hashedPassword := helpers.HashPassword(user.Password)
	user.Password = hashedPassword
	
	if err := config.DB.Save(&user).Error; err != nil {
	  return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
	  "message": "success create new user",
	  "user":    user,
	})
}
  
	
// delete user by id
func DeleteUserController(e echo.Context) error {
	searchID := e.Param("id")
	userID, err := strconv.Atoi(searchID)
	if err != nil {
		// Invalid ID format
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid ID format",
		})
	}
	
	var userDB []models.User
	if err := config.DB.First(&userDB, userID).Error; err != nil {
		// User not found
		return e.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "User not found",
		})
	}
	
	if err := config.DB.Delete(&userDB).Error; err != nil {
		// Error deleting user
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Error deleting user",
		})
	}
		// User deleted
		return e.JSON(http.StatusOK, map[string]interface{}{
			"message": "Successfully delete user",
			"data":    userDB,
		})
	}
	
	
// update user by id
func UpdateUserController(c echo.Context) error {
	searchID := c.Param("id")
	userID, err := strconv.Atoi(searchID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid ID format",
		})
	}

	var user models.User
	if err := config.DB.First(&user, userID).Error; err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "User not found",
		})
	}

	updatedUser := models.User{}
	if err := c.Bind(&updatedUser); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid request body",
		})
	}

	if updatedUser.Password != "" {
		updatedUser.Password = helpers.HashPassword(updatedUser.Password)
		user.Password = updatedUser.Password
	} else {
		updatedUser.Password = user.Password
	}

	if updatedUser.Name != "" {
		user.Name = updatedUser.Name
	}
	if updatedUser.Email != "" {
		user.Email = updatedUser.Email
	}

	// Save the updated user to the database
	if err := config.DB.Save(&user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Error updating user",
		})
	}

	userResponse := models.UserResponse{
		Id:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}

	// User updated
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Successfully update user",
		"data":    userResponse,
	})
}

 // login new user
 func LoginUserController(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)
	
	err := config.DB.Where("email = ?", user.Email).First(&user).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message" : "failed to login user",
			"error" : err.Error(),
		})
	}

	// Memeriksa kecocokan password
	err = helpers.ComparePassword(user.Password, c.FormValue("password"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message" : "email or password incorrect",
			"error" : err.Error(),
		})
	}

	token, err := middleware.CreateToken(int(user.ID), user.Name)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message" : "failed to login user",
			"error" : err.Error(),
		})
	}
	
	userResponse := models.UserResponse{Id: user.ID, Name: user.Name, Email: user.Email}

	return c.JSON(http.StatusOK, map[string]interface{}{
	  "message": "success Login User",
	  "user":    userResponse,
	  "token": token,
	})
}