package request


type UserCreateRequest struct{
	FullName string `json:"full_name" binding:"required"`
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserUpdateRequest struct{
	SchoolID *uint64    `json:"school_id"`
	FullName string    `json:"full_name"`
	Email    string    `json:"email"`
	Password *string    `json:"password"`
	Profile  string    `json:"profile"`
	IsActive bool      `json:"is_active"`
}

type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserOTP struct {
	Email string `json:"email"`
	Code string `json:"code"`
}