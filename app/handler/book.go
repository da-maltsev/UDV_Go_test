package handler

import (
	"github.com/da-maltsev/UDV_Go_test/model"
	"github.com/jinzhu/gorm"
	"net/http"
)

func GetBooksPaginator(db *gorm.DB, w http.ResponseWriter, r *http.Request) {

	books := []model.Book{}
	responseBooks := []model.Book{}
	response := db.Scopes(paginate(r)).Find(&books)
	for i, _ := range books {
		db.Model(books[i]).Related(&books[i].Publisher).Scan(&responseBooks)
	}
	if response == nil {
		return
	}
	respondJSON(w, http.StatusOK, books)

}

func GetBook(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	book := getBookOr404(db, w, r)
	if book == nil {
		return
	}
	respondJSON(w, http.StatusOK, book)

}

func getBookOr404(db *gorm.DB, w http.ResponseWriter, r *http.Request) *model.Book {
	book := model.Book{}
	if err := db.First(&book, model.Book{}).Related(&book.Publisher).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &book
}
