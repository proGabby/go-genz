package postgressDatasource

import (
	"database/sql"
	"time"

	_ "github.com/lib/pq"
	"github.com/proGabby/4genz/domain/entity"
)

type PostgresChatDBStore struct {
	DB *sql.DB
}

func NewPostgresChatDBStore(db *sql.DB) *PostgresChatDBStore {
	return &PostgresChatDBStore{
		DB: db,
	}
}

func (chatDb *PostgresChatDBStore) SaveMessage(message *entity.Message) error {

	query := "INSERT INTO Messages(sender_id, receiver_id, conversation_id ,message) VALUES($1, $2, $3)"
	_, err := chatDb.DB.Exec(query, message.SenderId, message.ReceiverId, message.ConversationId, message.Message)
	if err != nil {
		return err
	}

	return nil
}

// check if conversation exist between users on the conversation table
func (chatDb *PostgresChatDBStore) IsConversationExistBtwUsers(senderId int, receiverId int) (bool, error) {

	var conversationId int
	query := "SELECT id FROM Conversations WHERE (sender_id = $1 AND receiver_id = $2) OR (sender_id = $2 AND receiver_id = $1)"
	err := chatDb.DB.QueryRow(query, senderId, receiverId).Scan(&conversationId)
	if err != nil {
		return false, err
	}

	return true, nil

}

func (chatDb *PostgresChatDBStore) InitConversation(senderId, receiverId int) (int, error) {

	var conversationId int
	query := "INSERT INTO Conversations(sender_id, receiver_id, created_at) VALUES($1, $2) RETURNING id"
	err := chatDb.DB.QueryRow(query, senderId, receiverId, time.Now().Format("2006-01-02 15:04:05")).Scan(&conversationId)
	if err != nil {
		return 0, err
	}

	return conversationId, nil
}

func (chatDb *PostgresChatDBStore) GetUserConversations(userId int) ([]entity.Conversation, error) {

	var conversations []entity.Conversation
	query := "SELECT id, sender_id, receiver_id, created_at FROM Conversations WHERE sender_id = $1 OR receiver_id = $1"
	rows, err := chatDb.DB.Query(query, userId)
	if err != nil {
		return conversations, err
	}

	for rows.Next() {
		var conversation entity.Conversation
		err = rows.Scan(&conversation.Id, &conversation.SenderId, &conversation.ReceiverId, &conversation.CreatedAt)
		if err != nil {
			return conversations, err
		}
		conversations = append(conversations, conversation)
	}

	return conversations, nil
}

func (chatDb *PostgresChatDBStore) GetMessagesByConversationId(conversationId int) ([]entity.Message, error) {

	var messages []entity.Message
	query := "SELECT id, sender_id, receiver_id, message, created_at FROM Messages WHERE conversation_id = $1"
	rows, err := chatDb.DB.Query(query, conversationId)
	if err != nil {
		return messages, err
	}

	for rows.Next() {
		var message entity.Message
		err = rows.Scan(&message.Id, &message.SenderId, &message.ReceiverId, &message.Message, &message.CreatedAt)
		if err != nil {
			return messages, err
		}
		messages = append(messages, message)
	}

	return messages, nil
}
