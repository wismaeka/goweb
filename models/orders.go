package models

import "time"

type Item struct {
	ItemID      uint   `json:"item_id" gorm:"primary_key"`
	Item_code   string `json:"item_code"`
	Description string `json:"description"`
	Quantity    uint   `json:"quantity"`
	OrderID     uint   `json:"-"`
	// Id          int    `json:"item_id"`
	// Item_code   string `json:"item_code"`
	// Description string `json:"description"`
	// Quantity    int    `json:"quantity"`
	// Order_id    int    `json:"order_id"`
}

type Order struct {
	OrderID       uint      `json:"order_id" gorm:"primary_key"`
	Customer_name string    `json:"customer_name"`
	Order_at      time.Time `json:"order_at"`
	Items         []Item    `json:"items" gorm:"foreignkey:OrderID"`
	// Id            int       `json:"order_id"`
	// Customer_name string    `json:"customer_name"`
	// Order_at      time.Time `json:"order_at"`
	// Items         []Item    `json:"items" gorm:"foreignkey:order_id"`
}
