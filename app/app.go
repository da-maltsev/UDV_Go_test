package app

import (
	"github.com/da-maltsev/UDV_Go_test/app/handler"
	"github.com/da-maltsev/UDV_Go_test/config"
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

func (a *App) Initialize(config *config.Config) {
	//dbURI := fmt.Sprintf("%s:%s@/%s?charset=%s&parseTime=True",
	//	config.DB.Username,
	//	config.DB.Password,
	//	config.DB.Name,
	//	config.DB.Charset)

	db, err := gorm.Open("postgres", "host=localhost port=5432  user=postgres password=postgres dbname=booksdb ")
	//db, err := gorm.Open(config.DB.Dialect, dbURI)
	if err != nil {
		log.Fatalf("Could not connect database: %s", err)
	}

	a.DB = model.DBMigrate(db)
	a.Router = mux.NewRouter()
	a.setRouters()
}

func (a *App) setRouters() {
	a.Get("/api/book", a.GetBook)
	//a.Get("/api/book/xxx", a.GetEmployee)
}

func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

func (a *App) GetBook(w http.ResponseWriter, r *http.Request) {
	handler.GetBook(a.DB, w, r)
}

func (a *App) Run(host string) {
	log.Fatal(http.ListenAndServe(host, a.Router))
}
