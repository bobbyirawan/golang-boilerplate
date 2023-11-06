package service

import "go.uber.org/dig"

type Holder struct {
	dig.In

	UserService UserService
	AuthService AuthService
}

func Register(container *dig.Container) error {
	if err := container.Provide(NewUserService); err != nil {
		return err
	}

	if err := container.Provide(NewAuthService); err != nil {
		return err
	}

	return nil
}
