package models

import "time"

type UserModel struct {
	ID          string
	FirstName   string
	LastName    string
	Email       string
	Password    string
	PhoneArea   string
	PhoneNumber string
	VerifiedAt  *time.Time
	CreatedAt   time.Time
	CreatedBy   string
	ModifiedAt  *time.Time
	ModifiedBy  *string
	DeletedAt   *time.Time
	DeletedBy   *string
}
