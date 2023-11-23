package repository

import "go.uber.org/dig"

type Holder struct {
	dig.In

	UserRepository             UserRepository
	ContactRepository          ContactRepository
	ChatRepository             ChatRepository
	MessageRepository          MessageRepository
	MessageRecipientRepository MessageRecipientRepository
}

func Register(container *dig.Container) error {
	if err := container.Provide(NewUserRepository); err != nil {
		return err
	}

	if err := container.Provide(NewContactRepository); err != nil {
		return err
	}

	if err := container.Provide(NewChatRepository); err != nil {
		return err
	}

	if err := container.Provide(NewMessageRepository); err != nil {
		return err
	}

	if err := container.Provide(NewMessageRecipientRepository); err != nil {
		return err
	}

	return nil
}
