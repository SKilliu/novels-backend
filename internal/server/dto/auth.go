package dto

type SignUpRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	DeviceID string `json:"deviceId"`
}

type SignInRequest struct {
	Login    string `json:"login" example:"test_login"`
	Password string `json:"password" example:"supersecretpassword"`
}

type GuestSignInRequest struct {
	DeviceID string `json:"deviceId" example:"thisIsMyDeviceId"`
}

type ResetPasswordRequest struct {
	Email string `json:"email" example:"myemail@mail.com"`
}

type AuthResponse struct {
	ID          string `json:"id" example:"some_id"`
	Username    string `json:"username" example:"awesome_user"`
	Email       string `json:"email" example:"my@testmail.com"`
	Token       string `json:"token" example:"someSuperseCretToken.ForuseRAuthoriZATIon"`
	DateOfBirth int64  `json:"dateOfBirth" example:"12345672"`
	Gender      string `json:"gender" example:"male"`
	Membership  string `json:"membership" example:"some_info"`
	AvatarData  string `json:"avatarData" example:"avatar_data"`
	Rate        int    `json:"rate" example:"0"`
}

type SocialsSignInRequest struct {
	ID       string `json:"id"`
	Social   string `json:"social"`
	DeviceID string `json:"deviceId"`
}
