package models

import "time"

type Campaign struct {
	ID           int       `json:"id" gorm:"primary_key:auto_increment"`
	UserID       int       `json:"userId" gorm:"int(20)"`
	User         User      `json:"user" gorm:"int(20)"`
	Location     string    `json:"location" gorm:"type: varchar (255)"`
	RegisteredAt time.Time `json:"registeredAt"`
}
