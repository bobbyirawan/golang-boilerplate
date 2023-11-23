package service

import (
	"encoding/json"
	"go-chat/internal/dto"
	"go-chat/internal/model"
	"go-chat/internal/repository"
	"go-chat/protocols/socket"
	"log"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

type (
	ChatService interface {
		HandleSocketConnection(ctx echo.Context, req *dto.ChatReq) error
		GetListChatUser(req *dto.GetListChatUserReq, res *dto.GetListChatUserRes) error
	}

	chatDependency struct {
		socket     socket.Holder
		repository repository.Holder
	}
)

func NewChatService(socket socket.Holder, repository repository.Holder) ChatService {
	return &chatDependency{
		socket:     socket,
		repository: repository,
	}
}

func (impl *chatDependency) HandleSocketConnection(ctx echo.Context, req *dto.ChatReq) error {
	newSocket, err := impl.socket.Upgrader.Upgrade(ctx.Response(), ctx.Request(), nil)
	if err != nil {
		log.Println("Error while upgrading connection:", err)
		return err
	}
	defer newSocket.Close()

	var (
		reqMessage = &dto.Message{
			ChatID:      nil,
			SenderID:    "server",
			RecipientID: req.UserID,
			Content:     "connected",
		}
		client           = &socket.Client{}
		message          = &model.Message{}
		chat             = &model.Chat{}
		messageRecipient = &model.MessageRecipient{}
	)

	// send message welcome from server
	welcome, err := json.Marshal(reqMessage)
	if err != nil {
		log.Println("ERROR CONVERSION WELCOME MESSAGE FROM SERVER")
		return nil
	} else {
		newSocket.WriteMessage(websocket.TextMessage, welcome)
	}

	user, err := impl.getUserByID(req.UserID)
	if err != nil {
		log.Println("error get user email by id : ", err)
		return nil
	}

	client.Conn = newSocket
	client.UserID = user.ID
	client.Email = user.Email

	impl.socket.SocketConnectionRegistry.RegisterClient(client) // Register the client

	defer impl.socket.SocketConnectionRegistry.UnregisterClient(client) // unregister client connection

	for {

		if err := newSocket.ReadJSON(reqMessage); err != nil { // incoming message
			log.Printf("ERROR READ MESSAGE FROM %s : %s", client.UserID, err)
			break
		}

		if reqMessage.ChatID == nil { // if Chat ID is UnAvailable
			impl.createChatRoom(chat, message, reqMessage)

			recipientConn, ok := impl.socket.SocketConnectionRegistry.Connections[reqMessage.SenderID]
			reqMessage.ChatID = &chat.ID
			if ok {
				if err := recipientConn.WriteJSON(reqMessage); err != nil {
					log.Println("ERROR SEND MESSAGE TO CLIENT")
					break
				}
			}
		} else { // if Chat ID is Available
			message.ChatID = *reqMessage.ChatID
			message.SenderID = reqMessage.SenderID
			message.Content = reqMessage.Content
			impl.createMessage(message)
		}

		recipientConn, ok := impl.socket.SocketConnectionRegistry.Connections[reqMessage.RecipientID]
		if ok {
			if err := impl.repository.MessageRepository.UpdateStatusSend(message.ID); err != nil {

				log.Println("error update status send message")
			}

			messageRecipient.MessageID = message.ID
			messageRecipient.RecipientID = reqMessage.RecipientID
			messageRecipient.Status = model.MESSAGE_RECIPIENT_SEND
			if err := impl.createMessageRecipient(messageRecipient); err != nil {
				log.Println("error create message recipient")
			}

			if err := recipientConn.WriteJSON(reqMessage); err != nil {
				log.Println("ERROR SEND MESSAGE TO CLIENT")
				break
			}

			sender, ok := impl.socket.SocketConnectionRegistry.Connections[reqMessage.SenderID]
			if ok {
				if err := sender.WriteJSON(reqMessage); err != nil {
					log.Println("ERROR SEND mark message sended to sender")
					break
				}
			}

		} else {
			messageRecipient.MessageID = message.ID
			messageRecipient.RecipientID = reqMessage.RecipientID
			messageRecipient.Status = model.MESSAGE_RECIPIENT_NOT_SEND
			if err := impl.createMessageRecipient(messageRecipient); err != nil {
				log.Println("error create message recipient")
			}
		}
	}

	return nil
}

func (impl *chatDependency) GetListChatUser(req *dto.GetListChatUserReq, res *dto.GetListChatUserRes) error {

	listChatUser := new([]model.ListDetailChatUser)
	if err := impl.repository.ChatRepository.ListChatUser(req.ChatID, listChatUser); err != nil {
		return err
	}

	res.Messages = listChatUser

	return nil
}

func (impl *chatDependency) createChatRoom(chat *model.Chat, message *model.Message, req *dto.Message) error {
	if err := impl.repository.ChatRepository.Create(chat); err != nil {
		log.Println("error create chat room 1 on 1")
		return err
	}

	message.ChatID = chat.ID
	message.SenderID = req.SenderID
	message.Content = req.Content

	return impl.repository.MessageRepository.Create(message)
}

func (impl *chatDependency) createMessage(message *model.Message) error {
	return impl.repository.MessageRepository.Create(message)
}

func (impl *chatDependency) createMessageRecipient(messageRecipient *model.MessageRecipient) error {
	return impl.repository.MessageRecipientRepository.Create(messageRecipient)
}

func (impl *chatDependency) getUserByID(id string) (*model.User, error) {
	user := new(model.User)
	user.ID = id

	err := impl.repository.UserRepository.GetUserByID(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
