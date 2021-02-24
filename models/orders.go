package models

import "time"

type Item struct {
	ItemID      uint   `json:"item_id" gorm:"primary_key"`
	ItemCode    string `json:"item_code"`
	Description string `json:"description"`
	Quantity    uint   `json:"quantity"`
	OrderID     uint   `json:"-"`
}

type Order struct {
	OrderID      uint      `json:"order_id" gorm:"primary_key"`
	CustomerName string    `json:"customer_name"`
	OrderAt      time.Time `json:"order_at"`
	Items        []Item    `json:"items" gorm:"foreignkey:OrderID"`
}
