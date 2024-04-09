package models

import (
	"time"

	"github.com/google/uuid"
)

type SignUpInput struct {
	Name            string `json:"name" validate:"required"`
	Email           string `json:"email" validate:"required"`
	Password        string `json:"password" validate:"required,min=4"`
	PasswordConfirm string `json:"passwordConfirm" validate:"required,min=4"`
	Role            string `json:"role"`
}

type SignInInput struct {
	Email    string `json:"email"  validate:"required"`
	Password string `json:"password"  validate:"required"`
}

type UserResponse struct {
	ID        uuid.UUID `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Email     string    `json:"email,omitempty"`
	Role      string    `json:"role,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type CreateBannerDTO struct {
	TagIDs    []int64           `json:"tag_ids"`
	FeatureID int             `json:"feature_id"`
	Content   map[string]string `json:"content"`
	IsActive  bool              `json:"is_active"`
}
