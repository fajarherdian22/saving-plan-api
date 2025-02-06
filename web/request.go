package web

import (
	"time"

	"github.com/fajarherdian22/saving-plan-api/repository"
	"github.com/google/uuid"
)

type CreateUserRequest struct {
	Email string `json:"email" binding:"required,email"`
}

func CreateUserPayload(req CreateUserRequest) repository.CreateUserParams {
	return repository.CreateUserParams{
		ID:        uuid.NewString(),
		Email:     req.Email,
		CreatedAt: time.Now(),
	}
}
