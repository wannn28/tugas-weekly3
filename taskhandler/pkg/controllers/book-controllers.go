package controllers

import (
	"net/http"

	"TaskHandler/pkg/models"
	"TaskHandler/pkg/utils"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

var Database *gorm.DB // Variabel global untuk database

func FetchBooks(w http.ResponseWriter, r *http.Request) {
	var books []models.Book
	if err := Database.Find(&books).Error; err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Failed to retrieve books")
		return
	}
	utils.SendResponse(w, http.StatusOK, books)
}

func FetchBookByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var book models.Book
	if err := Database.First(&book, params["id"]).Error; err != nil {
		utils.SendError(w, http.StatusNotFound, "Book not found")
		return
	}
	utils.SendResponse(w, http.StatusOK, book)
}

func AddBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	if err := utils.DecodeBody(r, &book); err != nil {
		utils.SendError(w, http.StatusBadRequest, "Invalid input")
		return
	}
	if err := Database.Create(&book).Error; err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Failed to create book")
		return
	}
	utils.SendResponse(w, http.StatusCreated, book)
}

func ModifyBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var book models.Book
	if err := Database.First(&book, params["id"]).Error; err != nil {
		utils.SendError(w, http.StatusNotFound, "Book not found")
		return
	}

	var updatedBook models.Book
	if err := utils.DecodeBody(r, &updatedBook); err != nil {
		utils.SendError(w, http.StatusBadRequest, "Invalid input")
		return
	}

	// Update field buku
	book.Title = updatedBook.Title
	book.Author = updatedBook.Author
	book.Price = updatedBook.Price

	if err := Database.Save(&book).Error; err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Failed to update book")
		return
	}
	utils.SendResponse(w, http.StatusOK, book)
}

func RemoveBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var book models.Book
	if err := Database.First(&book, params["id"]).Error; err != nil {
		utils.SendError(w, http.StatusNotFound, "Book not found")
		return
	}
	if err := Database.Delete(&book).Error; err != nil {
		utils.SendError(w, http.StatusInternalServerError, "Failed to delete book")
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
