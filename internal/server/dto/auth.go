package dto

type SignUpRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	DeviceID string `json:"deviceId"`
}

type AuthResponse struct {
	ID          string `json:"id"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	Token       string `json:"token"`
	DateOfBirth int    `json:"dateOfBith"`
	Gender      string `json:"gender"`
	Membership  string `json:"membership"`
	AvatarData  string `json:"avatarData"`
	Rate        int    `json:"rate"`
}
