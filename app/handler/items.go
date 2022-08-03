package handler

import (
	"github.com/da-maltsev/UDV_Go_test/model"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"

	"github.com/jinzhu/gorm"
)

func GetItems(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	bookId, _ := strconv.Atoi(vars["id"])
	items := []model.Item{}
	books := []model.Book{}
	responseItems := []model.Item{}
	responseBooks := []model.Book{}
	db.Where("book_id = ?", bookId).Find(&items)
	db.Preload("Publisher").Find(&books)
	for i, _ := range items {
		db.Model(items[i]).Related(&items[i].Book).Related(&items[i].Book).Scan(&responseItems)
		db.Model(items[i].Book).Related(&items[i].Book.Publisher).Scan(&responseBooks)
	}
	if items == nil {
		return
	}
	respondJSON(w, http.StatusOK, items)
}
