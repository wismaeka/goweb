package models

import "time"

type Item struct {
	Id          int    `json:"item_id"`
	Item_code   string `json:"item_code"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
	Order_id    int    `json:"order_id"`
}

type Order struct {
	Id            int       `json:"order_id"`
	Customer_name string    `json:"customer_name"`
	Order_at      time.Time `json:"order_at"`
	Items         []Item    `json:"items" gorm:"foreignkey:order_id"`
}
