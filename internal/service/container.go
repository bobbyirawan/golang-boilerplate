package service

import "go.uber.org/dig"

type Holder struct {
	dig.In

	UserService             UserService
	AuthService             AuthService
	ContactService          ContactService
	ChatService             ChatService
	MessageRecipientService MessageRecipientService
}

func Register(container *dig.Container) error {
	if err := container.Provide(NewUserService); err != nil {
		return err
	}

	if err := container.Provide(NewAuthService); err != nil {
		return err
	}

	if err := container.Provide(NewContactService); err != nil {
		return err
	}

	if err := container.Provide(NewChatService); err != nil {
		return err
	}

	if err := container.Provide(NewMessageRecipientService); err != nil {
		return err
	}

	return nil
}
