package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"time"
)

type Publisher struct {
	ID       int    `gorm:"primary_key, AUTO_INCREMENT" json:"id"`
	Title    string `json:"title"`
	Address  string `json:"address"`
	Contacts string `json:"contacts"`
}

type Book struct {
	ID          int       `gorm:"primary_key, AUTO_INCREMENT" json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Authors     string    `json:"authors"`
	PublishYear time.Time `json:"publish_year"`
	PublisherID int       `json:"publisher_id"`
	Publisher   Publisher `json:"publisher"`
	Edition     string    `json:"edition"`
}

type Item struct {
	ID       int  `gorm:"primary_key, AUTO_INCREMENT" json:"id"`
	Quantity int  `json:"quantity"`
	BookID   int  `json:"book_id"`
	Book     Book `json:"book"`
}

type Position struct {
	ID     int    `gorm:"primary_key, AUTO_INCREMENT" json:"id"`
	Room   string `json:"room"`
	Shelf  string `json:"shelf"`
	ItemID int    `json:"item_id"`
	Item   Item   `json:"item"`
}

type Customer struct {
	ID       int    `gorm:"primary_key, AUTO_INCREMENT" json:"id"`
	Name     string `json:"name"`
	Contacts string `json:"contacts"`
}

type Status struct {
	ID         int       `gorm:"primary_key, AUTO_INCREMENT" json:"id"`
	ItemID     int       `json:"item_id"`
	Item       Item      `json:"item"`
	Available  bool      `json:"available"`
	CustomerID int       `json:"customer_id"`
	Customer   Customer  `json:"customer"`
	GivenAt    time.Time `json:"given_at"`
	ReceivedAt time.Time `json:"received_at"`
}

func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Publisher{}, &Book{}, &Item{}, &Position{}, &Customer{}, &Status{})
	return db
}
