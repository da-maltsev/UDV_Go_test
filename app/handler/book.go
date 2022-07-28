package handler

import (
	"net/http"

	"github.com/da-maltsev/UDV_Go_test/model"
	"github.com/jinzhu/gorm"
)

func GetBook(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	book := getBookOr404(db, w, r)
	if book == nil {
		return
	}
	respondJSON(w, http.StatusOK, book)
}

func getBookOr404(db *gorm.DB, w http.ResponseWriter, r *http.Request) *model.Book {
	book := model.Book{}
	if err := db.First(&book, model.Book{}).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &book
}
