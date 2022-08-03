package app

import (
	"github.com/da-maltsev/UDV_Go_test/app/handler"
	"github.com/da-maltsev/UDV_Go_test/model"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
)

type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

func (a *App) Initialize() {

	db, err := gorm.Open(
		"postgres",
		"host=localhost port=5432 user=postgres password=postgres dbname=booksdb TimeZone=Asia/Shanghai",
	)
	if err != nil {
		log.Fatalf("Could not connect database: %s", err)
	}

	a.DB = model.DBMigrate(db)
	a.Router = mux.NewRouter()
	a.setRouters()
}

func (a *App) setRouters() {
	a.GetQueries("/api/book/", a.GetBooksPaginator, "page", "{page}", "size", "{size}")
	a.Get("/api/book", a.GetBook)
	a.Get("/api/book/{id}/items", a.GetItems)
}

func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

func (a *App) GetQueries(path string, f func(w http.ResponseWriter, r *http.Request), pairs ...string) {
	a.Router.Path(path).Queries(pairs...).HandlerFunc(f).Methods("GET")
}

func (a *App) GetBook(w http.ResponseWriter, r *http.Request) {
	handler.GetBook(a.DB, w, r)
}

func (a *App) GetBooksPaginator(w http.ResponseWriter, r *http.Request) {
	handler.GetBooksPaginator(a.DB, w, r)
}

func (a *App) GetItems(w http.ResponseWriter, r *http.Request) {
	handler.GetItems(a.DB, w, r)
}

func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}
