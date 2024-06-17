package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Message struct {
	Body          string `json:"body"`
	ChatNumber    int    `json:"chat_number"`
	MessageNumber int    `json:"message_number"`
}

func createMessage(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("|%s| Request - Create new message.\n", r.Method)
	ctx := context.Background()
	vars := mux.Vars(r)
	appToken := vars["application_token"]
	chatNumber := vars["chat_number"]

	var msg Message

	// Decode the JSON request body
	if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if msg.Body == "" {
		http.Error(w, "Message body is required", http.StatusBadRequest)
		return
	}

	chatNum, err := strconv.Atoi(chatNumber)
	if err != nil {
		http.Error(w, "Invalid chat number", http.StatusBadRequest)
		return
	}

	messageNumber, err := createMessageInDB(ctx, appToken, chatNum, msg.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	msg.ChatNumber = chatNum
	msg.MessageNumber = messageNumber

	if err := indexMessageInES(ctx, msg); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(msg)
}

func createMessageInDB(ctx context.Context, appToken string, chatNumber int, body string) (int, error) {
	for {
		tx, err := _db.Begin()
		if err != nil {
			return 0, fmt.Errorf("failed to begin transaction: %v", err)
		}
		defer tx.Rollback()

		var chatID int
		if err := tx.QueryRowContext(ctx, "SELECT c.id FROM chats c JOIN applications a ON c.application_id = a.id WHERE a.token = ? AND c.number = ? FOR UPDATE", appToken, chatNumber).Scan(&chatID); err != nil {
			return 0, fmt.Errorf("failed to get chat ID: %v", err)
		}

		var messageNumber int
		if err := tx.QueryRowContext(ctx, "SELECT COALESCE(MAX(number), 0) + 1 FROM messages WHERE chat_id = ? FOR UPDATE", chatID).Scan(&messageNumber); err != nil {
			return 0, fmt.Errorf("failed to get message number: %v", err)
		}

		if _, err := tx.ExecContext(ctx, "INSERT INTO messages (created_at, updated_at, chat_id, number, body) VALUES (NOW(), NOW(), ?, ?, ?)", chatID, messageNumber, body); err != nil {
			if mysqlErr, ok := err.(*mysql.MySQLError); ok && mysqlErr.Number == 1062 {
				continue // Retry transaction in case of unique constraint violation
			}
			return 0, fmt.Errorf("failed to insert message: %v", err)
		}

		if err := tx.Commit(); err != nil {
			return 0, fmt.Errorf("failed to commit transaction: %v", err)
		}

		return messageNumber, nil
	}
}

func indexMessageInES(ctx context.Context, msg Message) error {
	esClient, err := GetESClient()
	if err != nil {
		return fmt.Errorf("failed to get Elasticsearch client: %v", err)
	}

	dataJSON, err := json.Marshal(msg)
	if err != nil {
		return fmt.Errorf("failed to marshal message: %v", err)
	}

	_, err = esClient.Index().
		Index("messages").
		BodyJson(string(dataJSON)).
		Do(ctx)
	if err != nil {
		return fmt.Errorf("failed to index message in Elasticsearch: %v", err)
	}

	return nil
}
