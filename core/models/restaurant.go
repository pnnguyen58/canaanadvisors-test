package models

type Restaurant struct {
	Id         int64       `json:"id,omitempty"`
	Name       string      `json:"name,omitempty"`
	Categories []*Category `json:"categories,omitempty"`
}
