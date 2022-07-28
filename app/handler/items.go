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
	db.Where("book_id = ?", bookId).Find(&items).Scan(&items)
	if items == nil {
		return
	}
	respondJSON(w, http.StatusOK, items)
}
