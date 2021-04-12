package dto

type SignUpRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	DeviceID string `json:"deviceId"`
}

type SignInReq struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type AuthResponse struct {
	ID          string `json:"id"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	Token       string `json:"token"`
	DateOfBirth int64  `json:"dateOfBith"`
	Gender      string `json:"gender"`
	Membership  string `json:"membership"`
	AvatarData  string `json:"avatarData"`
	Rate        int    `json:"rate"`
}
