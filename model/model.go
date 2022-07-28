package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"time"
)

type Publisher struct {
	gorm.Model
	Title    string `json:"title"`
	Address  string `json:"address"`
	Contacts string `json:"contacts"`
}

type Book struct {
	gorm.Model
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Authors     string    `json:"authors"`
	PublishYear time.Time `json:"publishYear"`
	PublisherID int       `json:"publisher_id"`
	Edition     string    `json:"edition"`
}

type Item struct {
	gorm.Model
	Quantity int `json:"quantity"`
	BookID   int `json:"book_id"`
}

type Position struct {
	gorm.Model
	Room   string `json:"room"`
	Shelf  string `json:"shelf"`
	ItemID int    `json:"item_id"`
}

type Customer struct {
	gorm.Model
	Name     string `json:"name"`
	Contacts string `json:"contacts"`
}

type Status struct {
	gorm.Model
	ItemID     int       `json:"item_id"`
	Available  bool      `json:"available"`
	CustomerID int       `json:"customer_id"`
	GivenAt    time.Time `json:"given_at"`
	ReceivedAt time.Time `json:"received_at"`
}

func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Publisher{}, &Book{}, &Item{}, &Position{}, &Customer{}, &Status{})
	return db
}
