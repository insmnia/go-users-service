package models

type User struct {
	ID       string `json:"id" gorm:"primary_key"`
	Username string `json:"username" gorm:"unique_index"`
	Password string `json:"password" gorm:"not null"`
}

type CreateUserResponse struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
