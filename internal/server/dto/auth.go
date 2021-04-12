package dto

type SignUpRequest struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type AuthResponse struct {
	Token string `json:"token"`
}