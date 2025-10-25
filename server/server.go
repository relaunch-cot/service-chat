package server

import (
	"context"

	"github.com/golang/protobuf/ptypes/empty"
	pb "github.com/relaunch-cot/lib-relaunch-cot/proto/chat"
	"github.com/relaunch-cot/service-chat/handler"
)

type chatResource struct {
	handler *handler.Handlers
	pb.UnimplementedChatServiceServer
}

func (r *chatResource) CreateNewChat(ctx context.Context, in *pb.CreateNewChatRequest) (*empty.Empty, error) {
	err := r.handler.Chat.CreateNewChat(&ctx, in.CreatedBy, in.UserIds)
	if err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}

func (r *chatResource) SendMessage(ctx context.Context, in *pb.SendMessageRequest) (*empty.Empty, error) {
	err := r.handler.Chat.SendMessage(&ctx, in.ChatId, in.SenderId, in.MessageContent)
	if err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}

func (r *chatResource) GetAllMessagesFromChat(ctx context.Context, in *pb.GetAllMessagesFromChatRequest) (*pb.GetAllMessagesFromChatResponse, error) {
	response, err := r.handler.Chat.GetAllMessagesFromChat(&ctx, in.ChatId)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (r *chatResource) GetAllChatsFromUser(ctx context.Context, in *pb.GetAllChatsFromUserRequest) (*pb.GetAllChatsFromUserResponse, error) {
	response, err := r.handler.Chat.GetAllChatsFromUser(&ctx, in.UserId)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (r *chatResource) GetChatFromUsers(ctx context.Context, in *pb.GetChatFromUsersRequest) (*pb.GetChatFromUsersResponse, error) {
	response, err := r.handler.Chat.GetChatFromUsers(&ctx, in.UserIds)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func NewChatServer(handler *handler.Handlers) pb.ChatServiceServer {
	return &chatResource{
		handler: handler,
	}
}
