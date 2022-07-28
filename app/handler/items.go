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

	item := []model.Item{}
	items := db.Where("book_id = ?", bookId).Find(&item)
	if items == nil {
		return
	}
	respondJSON(w, http.StatusOK, items)
}
