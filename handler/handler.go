package handler

import (
	"context"

	"github.com/google/uuid"
	pb "github.com/relaunch-cot/lib-relaunch-cot/proto/chat"
	"github.com/relaunch-cot/service-chat/repositories"
	"github.com/relaunch-cot/service-chat/resource/transformer"
)

type IChatHandler interface {
	CreateNewChat(ctx *context.Context, createdBy string, userIds []string) error
	SendMessage(ctx *context.Context, chatId, senderId, messageContent string) error
	GetAllMessagesFromChat(ctx *context.Context, chatId string) (*pb.GetAllMessagesFromChatResponse, error)
	GetAllChatsFromUser(ctx *context.Context, userId string) (*pb.GetAllChatsFromUserResponse, error)
	GetChatFromUsers(ctx *context.Context, userIds []string) (*pb.GetChatFromUsersResponse, error)
	GetChatById(ctx *context.Context, chatId string) (*pb.GetChatByIdResponse, error)
}

type resource struct {
	repositories *repositories.Repositories
}

func (r *resource) CreateNewChat(ctx *context.Context, createdBy string, userIds []string) error {
	chatId := uuid.New()
	err := r.repositories.Mysql.CreateNewChat(ctx, chatId.String(), createdBy, userIds)
	if err != nil {
		return err
	}

	return nil
}

func (r *resource) SendMessage(ctx *context.Context, chatId, senderId, messageContent string) error {
	messageId := uuid.New()
	err := r.repositories.Mysql.SendMessage(ctx, messageId.String(), chatId, senderId, messageContent)
	if err != nil {
		return err
	}

	return nil
}

func (r *resource) GetAllMessagesFromChat(ctx *context.Context, chatId string) (*pb.GetAllMessagesFromChatResponse, error) {
	mysqlResponse, err := r.repositories.Mysql.GetAllMessagesFromChat(ctx, chatId)
	if err != nil {
		return nil, err
	}

	baseModelsMessage, err := transformer.GetAllMessagesFromChatToBaseModels(mysqlResponse)
	if err != nil {
		return nil, err
	}

	getAllMessagesFromChatResponse := &pb.GetAllMessagesFromChatResponse{
		Messages: baseModelsMessage,
	}

	return getAllMessagesFromChatResponse, nil
}

func (r *resource) GetAllChatsFromUser(ctx *context.Context, userId string) (*pb.GetAllChatsFromUserResponse, error) {
	mysqlResponse, err := r.repositories.Mysql.GetAllChatsFromUser(ctx, userId)
	if err != nil {
		return nil, err
	}

	baseModelsChats, err := transformer.GetAllChatsFromUserToBaseModels(mysqlResponse)
	if err != nil {
		return nil, err
	}

	getAllChatsFromUserResponse := &pb.GetAllChatsFromUserResponse{
		Chats: baseModelsChats,
	}

	return getAllChatsFromUserResponse, nil
}

func (r *resource) GetChatFromUsers(ctx *context.Context, userIds []string) (*pb.GetChatFromUsersResponse, error) {
	mysqlResponse, err := r.repositories.Mysql.GetChatFromUsers(ctx, userIds)
	if err != nil {
		return nil, err
	}

	baseModelsChat, err := transformer.GetChatFromUsersToBaseModels(mysqlResponse)
	if err != nil {
		return nil, err
	}

	getChatFromUsersResponse := &pb.GetChatFromUsersResponse{
		Chat: baseModelsChat,
	}

	return getChatFromUsersResponse, nil
}

func (r *resource) GetChatById(ctx *context.Context, chatId string) (*pb.GetChatByIdResponse, error) {
	mysqlResponse, err := r.repositories.Mysql.GetChatById(ctx, chatId)
	if err != nil {
		return nil, err
	}

	baseModelsChat, err := transformer.GetChatByIdToBaseModels(mysqlResponse)
	if err != nil {
		return nil, err
	}

	getChatByIdResponse := &pb.GetChatByIdResponse{
		Chat: baseModelsChat,
	}

	return getChatByIdResponse, nil
}

func NewChatHandler(repositories *repositories.Repositories) IChatHandler {
	return &resource{
		repositories: repositories,
	}
}
