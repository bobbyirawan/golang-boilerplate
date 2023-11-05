package dto

type UserDTO struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	// ... tambahkan lebih banyak bidang sesuai kebutuhan
}
