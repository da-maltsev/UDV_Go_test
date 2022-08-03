package handler

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	"net/http"
	"strconv"
)

func respondJSON(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write([]byte(response))
}

func respondError(w http.ResponseWriter, code int, message string) {
	respondJSON(w, code, map[string]string{"error": message})
}

func paginate(r *http.Request) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		q := r.URL.Query()
		page, _ := strconv.Atoi(q.Get("page"))
		switch {
		case page <= 0:
			page = 1

			// Just for test purpose
		case page > 5:
			page = 5
		}
		pageSize, _ := strconv.Atoi(q.Get("size"))
		switch {
		case pageSize > 5:
			pageSize = 5
		case pageSize <= 0:
			pageSize = 5
		}
		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
