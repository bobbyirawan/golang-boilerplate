package service

import (
	"go-chat/config"
	"go-chat/internal/dto"
	"go-chat/internal/model"
	"go-chat/internal/repository"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type (
	AuthService interface {
		SignUp(signUp *dto.SignUp) error
		LogIn(login *dto.LoginReq) (*dto.LoginRes, error)
	}

	service struct {
		repository repository.Holder
		config     *config.Environment
	}
)

func NewAuthService(repo repository.Holder, config *config.Environment) AuthService {
	return &service{
		repository: repo,
		config:     config,
	}
}

func (impl *service) SignUp(signUp *dto.SignUp) error {

	if err := signUp.HashPassword(); err != nil {
		return err
	}

	user := &model.User{
		Email:    signUp.Email,
		Username: signUp.Username,
		Password: signUp.Password,
	}

	return impl.repository.UserRepository.CreateUser(user)
}

type MyJWT struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func (impl *service) LogIn(login *dto.LoginReq) (*dto.LoginRes, error) {

	user := &model.User{
		Email: login.Email,
	}

	if err := impl.repository.UserRepository.GetUserByEmail(user); err != nil {
		return nil, err
	}

	if err := user.CheckPassword(login.Password); err != nil {
		return nil, err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, MyJWT{
		ID:       user.ID,
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    strconv.Itoa(int(user.ID)),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	})

	ss, err := token.SignedString([]byte(impl.config.SecretKey))
	if err != nil {
		return nil, err
	}

	return &dto.LoginRes{
		ID:          user.ID,
		Username:    user.Username,
		Email:       user.Email,
		AccessToken: ss,
	}, nil
}
