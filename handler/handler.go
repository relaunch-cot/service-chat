package handler

import (
	"context"

	pb "github.com/relaunch-cot/lib-relaunch-cot/proto/chat"
	"github.com/relaunch-cot/service-chat/repositories"
)

type IChatHandler interface {
	CreateNewChat(ctx *context.Context, createdBy int64, userIds []int64) error
	SendMessage(ctx *context.Context, chatId, senderId int64, messageContent string) error
	GetAllMessagesFromChat(ctx *context.Context, chatId int64) (*pb.GetAllMessagesFromChatResponse, error)
}

type resource struct {
	repositories *repositories.Repositories
}

func (r *resource) CreateNewChat(ctx *context.Context, createdBy int64, userIds []int64) error {
	err := r.repositories.Mysql.CreateNewChat(ctx, createdBy, userIds)
	if err != nil {
		return err
	}

	return nil
}

func (r *resource) SendMessage(ctx *context.Context, chatId, senderId int64, messageContent string) error {
	err := r.repositories.Mysql.SendMessage(ctx, chatId, senderId, messageContent)
	if err != nil {
		return err
	}

	return nil
}

func (r *resource) GetAllMessagesFromChat(ctx *context.Context, chatId int64) (*pb.GetAllMessagesFromChatResponse, error) {
	mysqlResponse, err := r.repositories.Mysql.GetAllMessagesFromChat(ctx, chatId)
	if err != nil {
		return nil, err
	}

	getAllMessagesFromChatResponse := &pb.GetAllMessagesFromChatResponse{
		Messages: mysqlResponse,
	}

	return getAllMessagesFromChatResponse, nil
}

func NewChatHandler(repositories *repositories.Repositories) IChatHandler {
	return &resource{
		repositories: repositories,
	}
}
