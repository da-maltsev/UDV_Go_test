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
	Id          int       `gorm:"primary_key, AUTO_INCREMENT" json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Authors     string    `json:"authors"`
	PublishYear time.Time `json:"publishYear"`
	PublisherID int       `json:"publisher_id"`
	Publisher   Publisher
	Edition     string    `json:"edition"`
}

type Item struct {
	Id       int `gorm:"primary_key, AUTO_INCREMENT" json:"id"`
	Quantity int `json:"quantity"`
	BookID   int `json:"book_id"`
	Book     Book
}

type Position struct {
	gorm.Model
	Room   string `json:"room"`
	Shelf  string `json:"shelf"`
	ItemID int    `json:"item_id"`
	Item   Item
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
	Customer   Customer
	GivenAt    time.Time `json:"given_at"`
	ReceivedAt time.Time `json:"received_at"`
}

func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Publisher{}, &Book{}, &Item{}, &Position{}, &Customer{}, &Status{})
	return db
}
