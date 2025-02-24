package models

type Users struct {
	ID        int    `json:"id" gorm:"primaryKey"`
	Username  string `json:"username" gorm:"unique"`
	Password  string `json:"password" gorm:"not null"`
	Email     string `json:"email" gorm:"unique"`
	FullName  string `json:"fullName" gorm:"not null"`
	CreatedAt string `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt string `json:"updated_at" gorm:"autoUpdateTime"`
}

type RegisterPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	FullName string `json:"fullName"`
}
