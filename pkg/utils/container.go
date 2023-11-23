package utils

import (
	"go.uber.org/dig"
)

type Holder struct {
	dig.In

	Uuid        *UUID
	Logger      Logger
	TimeService TimeService
}

func Register(container *dig.Container) error {

	if err := container.Provide(NewIdGenertor); err != nil {
		return err
	}

	if err := container.Provide(NewLogger); err != nil {
		return err
	}

	if err := container.Provide(NewTimeService); err != nil {
		return err
	}

	return nil
}
