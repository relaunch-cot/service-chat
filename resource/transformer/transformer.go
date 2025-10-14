package transformer

import (
	"encoding/json"
	"errors"

	libModels "github.com/relaunch-cot/lib-relaunch-cot/models"
	pbBaseModels "github.com/relaunch-cot/lib-relaunch-cot/proto/base_models"
)

func GetAllChatsFromUserToBaseModels(in []*libModels.Chat) ([]*pbBaseModels.Chat, error) {
	var chat []*pbBaseModels.Chat

	b, err := json.Marshal(in)
	if err != nil {
		return nil, errors.New("Error marshalling chat model: " + err.Error())
	}

	err = json.Unmarshal(b, &chat)
	if err != nil {
		return nil, errors.New("Error unmarshalling chat model: " + err.Error())
	}

	return chat, nil
}

func GetAllMessagesFromChatToBaseModels(in []*libModels.Message) ([]*pbBaseModels.Message, error) {
	var message []*pbBaseModels.Message

	b, err := json.Marshal(in)
	if err != nil {
		return nil, errors.New("Error marshalling chat model: " + err.Error())
	}

	err = json.Unmarshal(b, &message)
	if err != nil {
		return nil, errors.New("Error unmarshalling chat model: " + err.Error())
	}

	return message, nil
}
