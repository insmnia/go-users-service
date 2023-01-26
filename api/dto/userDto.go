package dto

import (
	"github.com/google/uuid"
	"time"
)

type CreateUserRequest struct {
	Username       string `json:"username" binding:"required"`
	Password       string `json:"password" binding:"required"`
	PasswordRepeat string `json:"passwordRepeat" binding:"required"`
}
type UpdateUserRequest struct {
	Username string `json:"username" binding:"required"`
}

type UserResponse struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"createdAt"`
}
