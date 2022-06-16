package models

import (
	"time"
)

type Answer struct {
	Id               int    `gorm:"primary_key" json:"id"`
	Body             string `sql:"type:text;" json:"body"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	Views            int
	Likes            int
	Dislikes         int
	IsAcceptedAnswer bool
	UserID           int `gorm:"size:10" sql:"type:integer REFERENCES user(id)"`
	User             User
	QuestionID       int `gorm:"size:10" sql:"type:integer REFERENCES questions(id)"`
	Question         Question
	//Question_id      int `sql:"type:integer REFERENCES questions(id)"`
	//User_id          int `sql:"type:integer REFERENCES user(id)"`
}
