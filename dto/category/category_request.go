package artisdto

type CategoryRequest struct {
	Name string `gorm:"type: varchar" json:"name" validate:"required"`
}
