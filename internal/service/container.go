package service

import "go.uber.org/dig"

type Holder struct {
	dig.In

	UserService UserService
}

func Register(container *dig.Container) error {
	if err := container.Provide(NewUserService); err != nil {
		return err
	}

	return nil
}
