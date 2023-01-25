package dto

import (
	"github.com/google/uuid"
	"time"
)

type CreateUserRequest struct {
	Username       string `json:"username"`
	Password       string `json:"password"`
	PasswordRepeat string `json:"passwordRepeat"`
}
type UpdateUserRequest struct {
	Username string `json:"username"`
}

type UserResponse struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"createdAt"`
}
