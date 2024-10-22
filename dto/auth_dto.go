package dto

type SignupInput struct {
	// binding:"email"はemailの形式であることをバリデーションしている
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}
