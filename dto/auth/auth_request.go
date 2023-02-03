package authdto

type RegisterRequest struct {
	Name     string `gorm:"type: varchar" json:"fullName" validate:"required"`
	Email    string `gorm:"type: varchar" json:"email" validate:"required"`
	Password string `gorm:"type: varchar" json:"password" validate:"required"`
	Profile  string `gorm:"type: varchar" json:"profile"`
	Role     string `gorm:"type: varchar" json:"role"`
	Phone    string `json:"phone"`
}

type LoginRequest struct {
	Email    string `gorm:"type: varchar(255)" json:"email" validate:"required"`
	Password string `gorm:"type: varchar(255)" json:"password" validate:"required"`
}
