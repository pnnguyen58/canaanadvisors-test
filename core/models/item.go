package models

type Item struct {
	Id          int64   ` json:"id,omitempty"`
	Name        string  `json:"name,omitempty"`
	Description string  `json:"description,omitempty"`
	Price       float32 `json:"price,omitempty"`
}

type Category struct {
	Id    int64   `json:"id,omitempty"`
	Name  string  `json:"name,omitempty"`
	Items []*Item `json:"items,omitempty"`
}