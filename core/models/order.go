package models

import "time"

type Order struct {
	Id int64 `json:"id,omitempty"`
	Description string `json:"description,omitempty"`
	RestaurantId int64  `json:"restaurantId,omitempty"`
	Items []Item `json:"items,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}

type OrderCreate struct {
	Name    string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

func (oc *OrderCreate) CheckValid() error {
	// To check flow input here, return true if valid
	return nil
}