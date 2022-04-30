package models

type User struct {
	ID        int64  `gorm:"primaryKey" json:"id"`
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Email     string `json:"email" binding:"required"`
	PFPUrl    string `json:"pfp_url" binding:"required"`
	Bio       string `json:"bio" binding:"required"`
}
