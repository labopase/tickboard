package v1

type GetUserRequestDto struct {
	ID string `json:"id"`
}

type GetUserResponseDto struct {
	ID          string `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	PhoneArea   string `json:"phone_area"`
	PhoneNumber string `json:"phone_number"`
}
