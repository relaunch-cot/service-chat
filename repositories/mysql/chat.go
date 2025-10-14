package mysql

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/relaunch-cot/lib-relaunch-cot/repositories/mysql"

	pbBaseModels "github.com/relaunch-cot/lib-relaunch-cot/proto/base_models"
)

type mysqlResource struct {
	client *mysql.Client
}

type IMySqlChat interface {
	CreateNewChat(ctx *context.Context, createdBy int64, userIds []int64) error
	SendMessage(ctx *context.Context, chatId, senderId int64, messageContent string) error
	GetAllMessagesFromChat(ctx *context.Context, chatId int64) ([]*pbBaseModels.Message, error)
}

func (r *mysqlResource) CreateNewChat(ctx *context.Context, createdBy int64, userIds []int64) error {
	currentDate := time.Now()

	queryValidate := fmt.Sprintf(
		`SELECT * 
					FROM chats 
					WHERE user1_id = %d AND user2_id = %d
					OR user1_id = %d AND user2_id = %d`,
		userIds[0], userIds[1], userIds[0], userIds[1],
	)
	rows, err := mysql.DB.QueryContext(*ctx, queryValidate)
	if err != nil {
		return err
	}

	defer rows.Close()
	if rows.Next() {
		return errors.New("already exists an chat with these participants")
	}

	basequery := fmt.Sprintf(
		"INSERT INTO chats (createdBy, user1_id, user2_id, createdAt) VALUES(%d, %d, %d, '%s')",
		createdBy,
		userIds[0],
		userIds[1],
		currentDate.Format("2006-01-02 15:04:05"),
	)

	rows, err = mysql.DB.QueryContext(*ctx, basequery)
	if err != nil {
		return err
	}

	defer rows.Close()

	return nil
}

func (r *mysqlResource) SendMessage(ctx *context.Context, chatId, senderId int64, messageContent string) error {
	currentDate := time.Now()

	queryValidation := fmt.Sprintf(
		`SELECT * 
					FROM chats 
					WHERE chatId = %d 
					    AND user1_id = %d OR user2_id = %d`,
		chatId, senderId, senderId,
	)

	rows, err := mysql.DB.QueryContext(*ctx, queryValidation)
	if err != nil {
		return err
	}

	defer rows.Close()
	if !rows.Next() {
		return errors.New("this user is not part of this chat")
	}

	basequery := fmt.Sprintf(
		"INSERT INTO messages (chatId, senderId, content, createdAt) VALUES(%d, %d, '%s', '%s')",
		chatId,
		senderId,
		messageContent,
		currentDate.Format("2006-01-02 15:04:05"),
	)

	rows, err = mysql.DB.QueryContext(*ctx, basequery)
	if err != nil {
		return err
	}

	defer rows.Close()

	return nil
}

func (r *mysqlResource) GetAllMessagesFromChat(ctx *context.Context, chatId int64) ([]*pbBaseModels.Message, error) {
	baseQuery := fmt.Sprintf(`SELECT * FROM messages WHERE chatId = %d`, chatId)

	rows, err := mysql.DB.QueryContext(*ctx, baseQuery)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	messages := make([]*pbBaseModels.Message, 0)
	for rows.Next() {
		message := &pbBaseModels.Message{}

		err := rows.Scan(
			&message.MessageId,
			&message.ChatId,
			&message.SenderId,
			&message.MessageContent,
			&message.CreatedAt,
		)
		if err != nil {
			return nil, errors.New("error scanning mysql row: " + err.Error())
		}

		messages = append(messages, message)
	}

	if err = rows.Err(); err != nil {
		return nil, errors.New("row iteration error: " + err.Error())
	}

	return messages, nil
}

func NewMysqlRepository(client *mysql.Client) IMySqlChat {
	return &mysqlResource{
		client: client,
	}
}
