package service

import (
	"go-chat/internal/dto"
	"go-chat/internal/model"
	"go-chat/internal/repository"
)

type (
	MessageRecipientService interface {
		GetAllMessages(req *dto.GetListChatReq, res *dto.GetListChatRes) error
	}

	messageRecipientDependency struct {
		repository repository.Holder
	}
)

func NewMessageRecipientService(repository repository.Holder) MessageRecipientService {
	return &messageRecipientDependency{
		repository: repository,
	}
}

func (impl *messageRecipientDependency) GetAllMessages(req *dto.GetListChatReq, res *dto.GetListChatRes) error {
	listChat := new([]model.ListChat)

	if err := impl.repository.MessageRecipientRepository.GetAllMessages(req.UserID, listChat); err != nil {
		return err
	}

	res.Messages = listChat

	return nil

}
