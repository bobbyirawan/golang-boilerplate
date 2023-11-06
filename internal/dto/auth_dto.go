package dto

import "golang.org/x/crypto/bcrypt"

type SignUp struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (sign *SignUp) HashPassword() error {
	passwd, err := bcrypt.GenerateFromPassword([]byte(sign.Password), bcrypt.DefaultCost)
	sign.Password = string(passwd)

	if err != nil {
		return err
	}

	return nil
}

type LoginReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRes struct {
	ID          uint64 `json:"id"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	AccessToken string `json:"access_token"`
}
