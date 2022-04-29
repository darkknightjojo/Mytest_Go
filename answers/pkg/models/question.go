package models

import "time"

type Question struct {
	Id             int `gorm:"primary_key" json:"id"`
	Title          string
	Body           string `gorm:"type:text;" json:"body"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	Views          int
	Likes          int
	AnswerCount    int
	AcceptedAnswer bool
	UserID         int `gorm:"size:10" sql:"type:integer REFERENCES user(id)"`
	User           User
	Tags           []Tag `gorm:"many2many:taggings;"`
	//User_id        int   `sql:"type:integer REFERENCES user(id)"`
}
