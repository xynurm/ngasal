package models

import "time"

type User struct {
	ID        int       `json:"id" gorm:"primary_key:auto_increment"`
	Name      string    `json:"name" gorm:"type: varchar (50)"`
	Email     string    `json:"email" gorm:"type: varchar (120)"`
	Password  string    `json:"password" gorm:"type: varchar (255)"`
	Profile   string    `json:"profile" gorm:"type: varchar (255)"`
	Phone     string    `json:"phone" gorm:"type: varchar (20)"`
	Role      string    `json:"role" gorm:"type: varchar (1)"`
	CreatedAt time.Time `json:"createdAt"`
	LastLogin time.Time `json:"lastLogin"`
}
