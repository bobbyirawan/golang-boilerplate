package repository

import "go.uber.org/dig"

type Holder struct {
	dig.In

	UserRepository UserRepository
}

func Register(container *dig.Container) error {
	if err := container.Provide(NewUserRepository); err != nil {
		return err
	}

	return nil
}
