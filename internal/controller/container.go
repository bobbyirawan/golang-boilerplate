package controller

import "go.uber.org/dig"

type Holder struct {
	dig.In

	UserController UserController
}

func Register(container *dig.Container) error {
	if err := container.Provide(NewUserController); err != nil {
		return err
	}

	return nil
}
