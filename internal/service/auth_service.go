package service

import (
	"go-chat/internal/dto"
	"go-chat/internal/model"
	"go-chat/internal/repository"
	"go-chat/pkg/config"
	"go-chat/pkg/utils"

	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
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

	if err := impl.repository.UserRepository.GetUserByEmail(user); err == nil {
		log.Println("ERROR GET USER BY EMAIL: ", err)
		return echo.NewHTTPError(http.StatusBadRequest, "Email has already registered")
	}

	err := impl.repository.UserRepository.CreateUser(user)
	if err != nil {
		log.Println("ERROR CREATE USER: ", err)
		return err
	}

	return nil
}

type MyJWT struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func (impl *service) LogIn(login *dto.LoginReq) (*dto.LoginRes, error) {

	user := &model.User{
		Email: login.Email,
	}

	if err := impl.repository.UserRepository.GetUserByEmail(user); err != nil {
		log.Println("ERROR GET USER BY EMAIL: ", err)
		return nil, echo.NewHTTPError(http.StatusBadRequest, "Email of Password is invalid")
	}

	if err := utils.CheckPassword(user.Password, login.Password); err != nil {
		log.Println("ERROR CHECK PASSWORD: ", err)
		return nil, echo.NewHTTPError(http.StatusBadRequest, "Email of Password is invalid")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, MyJWT{
		ID:       user.ID,
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    user.ID,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	})

	ss, err := token.SignedString([]byte(impl.config.SECRET_KEY))
	if err != nil {
		log.Println("ERROR SIGN TOKEN: ", err)
		return nil, err
	}

	return &dto.LoginRes{
		ID:          user.ID,
		Username:    user.Username,
		Email:       user.Email,
		AccessToken: ss,
	}, nil
}
