package models

import "time"

type Tag struct {
	Id        int `gorm:"primary_key" json:"id"`
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	Questions []Question `gorm:"many2many:taggings;"`
}
