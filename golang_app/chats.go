package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func createChat(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("|%s| Request - Create new chat.\n", r.Method)
	vars := mux.Vars(r)
	appToken := vars["application_token"]

	for {
		tx, err := _db.Begin()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		defer tx.Rollback()

		var appID int
		err = tx.QueryRow("SELECT id FROM applications WHERE token = ?", appToken).Scan(&appID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		var chatNumber int
		err = tx.QueryRow("SELECT COALESCE(MAX(number), 0) + 1 FROM chats WHERE application_id = ? FOR UPDATE", appID).Scan(&chatNumber)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		_, err = tx.Exec("INSERT INTO chats (created_at, updated_at, application_id, number) VALUES (NOW(), NOW(), ?, ?)", appID, chatNumber)
		if err != nil {
			// Check if the error is a unique constraint violation
			if mysqlErr, ok := err.(*mysql.MySQLError); ok && mysqlErr.Number == 1062 {
				// Retry the transaction
				continue
			}
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = tx.Commit()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]int{"number": chatNumber})
		return
	}
}
