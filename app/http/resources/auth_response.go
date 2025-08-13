package resources

type AuthResponse struct {
	Token JWTResponse  `json:"token"`
	User  UserResource `json:"user"`
}

func NewAuthResponse(token JWTResponse, user UserResource) *AuthResponse {
	return &AuthResponse{
		Token: token,
		User:  user,
	}
}
