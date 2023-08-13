package models

import (
	"gorm.io/gorm/clause"
	"time"
)

type Order struct {
	Id int64 `json:"id,omitempty"  gorm:"<-:create;AUTO_INCREMENT;primaryKey;NOT NULL"`
	Description string `json:"description,omitempty" gorm:"column:description;type:text;NULL"`
	RestaurantId int64  `json:"restaurantId,omitempty" gorm:"column:restaurant_id;type:bigint;NOT NULL"`
	Items []Item `json:"items,omitempty" gorm:"embedded;column:items;type:jsonb;NOT NULL"`
	CreatedAt time.Time `json:"createdAt,omitempty" gorm:"<-:create;column:created_at;NULL"`
}

func (Order) TableName() string {
	return "orders"
}

func (Order) InsertClause() []clause.Expression {
	return []clause.Expression{
		clause.OnConflict{
			Columns:   []clause.Column{{Name: "id"}},
			DoNothing: true,
		},
	}
}