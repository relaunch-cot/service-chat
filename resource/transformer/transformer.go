package transformer

import (
	"encoding/json"

	libModels "github.com/relaunch-cot/lib-relaunch-cot/models"
	pbBaseModels "github.com/relaunch-cot/lib-relaunch-cot/proto/base_models"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func GetAllChatsFromUserToBaseModels(in []*libModels.Chat) ([]*pbBaseModels.Chat, error) {
	var chat []*pbBaseModels.Chat

	b, err := json.Marshal(in)
	if err != nil {
		return nil, status.Error(codes.Internal, "Error marshalling chat model: "+err.Error())
	}

	err = json.Unmarshal(b, &chat)
	if err != nil {
		return nil, status.Error(codes.Internal, "Error unmarshalling chat model: "+err.Error())
	}

	return chat, nil
}

func GetAllMessagesFromChatToBaseModels(in []*libModels.Message) ([]*pbBaseModels.Message, error) {
	var message []*pbBaseModels.Message

	b, err := json.Marshal(in)
	if err != nil {
		return nil, status.Error(codes.Internal, "Error marshalling chat model: "+err.Error())
	}

	err = json.Unmarshal(b, &message)
	if err != nil {
		return nil, status.Error(codes.Internal, "Error unmarshalling chat model: "+err.Error())
	}

	return message, nil
}

func GetChatFromUsersToBaseModels(in *libModels.Chat) (*pbBaseModels.Chat, error) {
	var chat pbBaseModels.Chat

	b, err := json.Marshal(in)
	if err != nil {
		return nil, status.Error(codes.Internal, "Error marshalling chat model: "+err.Error())
	}

	err = json.Unmarshal(b, &chat)
	if err != nil {
		return nil, status.Error(codes.Internal, "Error unmarshalling chat model: "+err.Error())
	}

	return &chat, nil
}

func GetChatByIdToBaseModels(in *libModels.Chat) (*pbBaseModels.Chat, error) {
	var chat pbBaseModels.Chat

	b, err := json.Marshal(in)
	if err != nil {
		return nil, status.Error(codes.Internal, "Error marshalling chat model: "+err.Error())
	}

	err = json.Unmarshal(b, &chat)
	if err != nil {
		return nil, status.Error(codes.Internal, "Error unmarshalling chat model: "+err.Error())
	}

	return &chat, nil
}
