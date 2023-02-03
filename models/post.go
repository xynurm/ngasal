package models

import "time"

type Post struct {
	ID          int       `json:"id" gorm:"primary_key:auto_increment"`
	Title       string    `json:"title" gorm:"type: varchar(50)"`
	Thumbnail   string    `json:"thumbnail" gorm:"type: varchar(255)"`
	Description string    `json:"description" gorm:"type: varchar(255)"`
	Status      int       `json:"status" gorm:"type: int(1)"`
	StartDate   time.Time `json:"startDate"`
	DueDate     time.Time `json:"dueDate"`
	CampaignID  int       `json:"campaignId"  gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Campaign    Campaign  `json:"campaign"`
	CategoryID  int       `json:"categoryId"  gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Category    Category  `json:"category"`
	Slug        string    `json:"slug" gorm:"type: varchar(50)"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
