package response

type UserResponse struct {
	ID uint64 `json:"id"`
	Email string `json:"email"`
	Token string `json:"token"`
}