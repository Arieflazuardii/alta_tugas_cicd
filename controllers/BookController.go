package controllers

import (
	"net/http"
	"praktikum/config"
	"praktikum/models"
	"strconv"

	"github.com/labstack/echo/v4"
)

// get all books
func GetBooksController(e echo.Context) error {
	var bookDB []models.Book
	if err := config.DB.Find(&bookDB).Error; err != nil {
	  return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	var books []models.BookResponse
	for _, u := range bookDB {
		bookRes := models.BookResponse{
			Id:         u.ID,
			Title:      u.Title,
			Author:     u.Author,
			Publisher:  u.Publisher,
			CreatedAt:  u.CreatedAt,
			UpdatedAt:  u.UpdatedAt,
		}
		books = append(books, bookRes)
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
	  "message": "success get all books",
	  "books":   books,
	})
}
  
  
  // get book by id
func GetBookController(e echo.Context) error {
	// your solution here
	searchID := e.Param("id")
	var bookDB []models.Book
	if err := config.DB.First(&bookDB, searchID).Error; err != nil {
	  // book not found
	  return e.JSON(http.StatusNotFound, map[string]interface{}{
		"message": "book not found",
	  })
	}
	
	var books []models.BookResponse
	for _, u := range bookDB {
		bookRes := models.BookResponse{
			Id:          u.ID,
			Title:       u.Title,
			Author:      u.Author,
			Publisher:   u.Publisher,
			CreatedAt:   u.CreatedAt,
			UpdatedAt:   u.UpdatedAt,
		}
		books = append(books, bookRes)
	}
	// book found
	return e.JSON(http.StatusOK, map[string]interface{}{
	  "message": "Successfully get book by ID",
	  "book": books,
	})
}
  
  
  // create new book
func CreateBookController(e echo.Context) error {
	book := models.Book{}
	e.Bind(&book)
  
  
	if err := config.DB.Save(&book).Error; err != nil {
	  return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return e.JSON(http.StatusOK, map[string]interface{}{
	  "message": "success create new book",
	  "book":    book,
	})
}
  
	
	// delete book by id
func DeleteBookController(e echo.Context) error {
	searchID := e.Param("id")
	bookID, err := strconv.Atoi(searchID)
	if err != nil {
	// Invalid ID format
		return e.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid ID format",
		})
	}
	
	var bookDB []models.Book
	if err := config.DB.First(&bookDB, bookID).Error; err != nil {
		// book not found
		return e.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "book not found",
		})
	}
	
	if err := config.DB.Delete(&bookDB).Error; err != nil {
		// Error deleting book
		return e.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Error deleting book",
		})
	}
	// book deleted
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "Successfully delete book",
		"data":    bookDB,
	})
}
	
	
// update book by id
func UpdateBookController(c echo.Context) error {
	searchID := c.Param("id")
	bookID, err := strconv.Atoi(searchID)
	if err != nil {
		  // Invalid ID format
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid ID format",
		})
	}
  
	var book models.Book
	if err := config.DB.First(&book, bookID).Error; err != nil {
		// User not found
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"message": "User not found",
		})
	}
  
	// Bind the updated data from the request body
	if err := c.Bind(&book); err != nil {
		// Invalid request body
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "Invalid request body",
		})
	}

	if err := config.DB.Save(&book).Error; err != nil {
		 // Error updating user
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Error updating user",
		})
	}
	// User updated
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Successfully update user",
		"data":    book,
	})
}