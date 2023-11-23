package controller

import "go.uber.org/dig"

type Holder struct {
	dig.In

	UserController             *UserController
	AuthController             *AuthController
	ContactController          *ContactController
	ChatController             *ChatController
	MessageRecipientController *MessageRecipientController
}

func Register(container *dig.Container) error {
	if err := container.Provide(NewUserController); err != nil {
		return err
	}

	if err := container.Provide(NewAuthController); err != nil {
		return err
	}

	if err := container.Provide(NewContactController); err != nil {
		return err
	}

	if err := container.Provide(NewChatController); err != nil {
		return err
	}

	if err := container.Provide(NewMessageRecipientController); err != nil {
		return err
	}

	return nil
}
