package dto

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

type CreateUserRequest struct {
	Username       string `json:"username"`
	Password       string `json:"password"`
	PasswordRepeat string `json:"passwordRepeat"`
}

type CreateUserResponse struct {
	ID        uuid.UUID `json:"id"`
	Username  uuid.UUID `json:"username"`
	CreatedAt time.Time `json:"createdAt"`
}
