package service

import (
	"go-chat/internal/dto"
	"go-chat/internal/model"
	"go-chat/internal/repository"
)

type (
	ContactService interface {
		Create(req *dto.CreateContactReq) error
		GetAllByID(req *dto.GetAllContactByIDReq, res *[]dto.GetAllContactByIDRes) error
	}

	contactService struct {
		repo repository.Holder
	}
)

func NewContactService(repo repository.Holder) ContactService {
	return &contactService{
		repo: repo,
	}
}

func (impl *contactService) Create(req *dto.CreateContactReq) error {

	user := new(model.User)
	user.Email = req.Email

	if err := impl.repo.UserRepository.GetUserByEmail(user); err != nil {
		return err
	}

	contact := new(model.Contact)
	contact.UserID = req.UserID
	contact.RecipientID = user.ID
	contact.Username = req.Username
	contact.Email = req.Email

	return impl.repo.ContactRepository.Create(contact)
}

func (impl *contactService) GetAllByID(req *dto.GetAllContactByIDReq, res *[]dto.GetAllContactByIDRes) error {
	contacts := new([]model.Contact)
	if err := impl.repo.ContactRepository.GetAllByID(req.UserID, contacts); err != nil {
		return err
	}

	contact := &dto.GetAllContactByIDRes{}

	for _, value := range *contacts {
		contact.Username = value.Username
		contact.Email = value.Email
		contact.RecipientID = value.RecipientID

		(*res) = append((*res), *contact)
	}

	return nil
}
