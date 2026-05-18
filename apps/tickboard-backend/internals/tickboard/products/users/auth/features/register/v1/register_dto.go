package v1

// RegisterRequestDto represents the request data for the register endpoint.
type RegisterRequestDto struct {
	Email           string `json:"email" validate:"required"`
	FirstName       string `json:"first_name" validate:"required"`
	LastName        string `json:"last_name" validate:"required"`
	Password        string `json:"password" validate:"required"`
	ConfirmPassword string `json:"confirm_password" validate:"required"`
}

// RegisterResponseDto represents the response data for the register endpoint.
type RegisterResponseDto struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
