package mysql

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/relaunch-cot/lib-relaunch-cot/repositories/mysql"

	libModels "github.com/relaunch-cot/lib-relaunch-cot/models"
)

type mysqlResource struct {
	client *mysql.Client
}

type IMySqlChat interface {
	CreateNewChat(ctx *context.Context, createdBy int64, userIds []int64) error
	SendMessage(ctx *context.Context, chatId, senderId int64, messageContent string) error
	GetAllMessagesFromChat(ctx *context.Context, chatId int64) ([]*libModels.Message, error)
	GetAllChatsFromUser(ctx *context.Context, userId int64) ([]*libModels.Chat, error)
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

func (r *mysqlResource) GetAllMessagesFromChat(ctx *context.Context, chatId int64) ([]*libModels.Message, error) {
	baseQuery := fmt.Sprintf(`SELECT * FROM messages WHERE chatId = %d`, chatId)

	rows, err := mysql.DB.QueryContext(*ctx, baseQuery)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	messages := make([]*libModels.Message, 0)
	for rows.Next() {
		message := &libModels.Message{}

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

func (r *mysqlResource) GetAllChatsFromUser(ctx *context.Context, userId int64) ([]*libModels.Chat, error) {
	baseQuery := fmt.Sprintf(
		`SELECT
    c.chatId,
    c.createdAt,
    c.createdBy,
    u1.userId   AS user1_id,
    u1.name AS user1_name,
    u1.email AS user1_email,
    u2.userId   AS user2_id,
    u2.name AS user2_name,
    u2.email AS user2_email
FROM
    chats c
JOIN
    users u1 ON c.user1_id = u1.userId
JOIN
    users u2 ON c.user2_id = u2.userId
WHERE
    c.user1_id = %d OR c.user2_id = %d;`,
		userId, userId)

	rows, err := mysql.DB.QueryContext(*ctx, baseQuery)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	chats := make([]*libModels.Chat, 0)
	for rows.Next() {
		chat := &libModels.Chat{
			User1: libModels.User{},
			User2: libModels.User{},
		}

		err := rows.Scan(
			&chat.ChatId,
			&chat.CreatedAt,
			&chat.CreatedBy,
			&chat.User1.UserId,
			&chat.User1.Name,
			&chat.User1.Email,
			&chat.User2.UserId,
			&chat.User2.Name,
			&chat.User2.Email,
		)
		if err != nil {
			return nil, errors.New("error scanning mysql row: " + err.Error())
		}

		chats = append(chats, chat)
	}

	if err = rows.Err(); err != nil {
		return nil, errors.New("row iteration error: " + err.Error())
	}

	return chats, nil
}

func NewMysqlRepository(client *mysql.Client) IMySqlChat {
	return &mysqlResource{
		client: client,
	}
}
