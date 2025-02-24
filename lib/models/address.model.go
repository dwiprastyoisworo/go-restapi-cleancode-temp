package models

type Address struct {
	ID        int    `json:"id" gorm:"primaryKey"`
	City      string `json:"city" gorm:"not null"`
	Address   string `json:"address" gorm:"not null"`
	UserID    int    `json:"user_id" gorm:"not null"`
	CreatedAt string `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt string `json:"updated_at" gorm:"autoUpdateTime"`
}
