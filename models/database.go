package models

import (
	"database/sql"
	"fmt"
	"github.com/lib/pq"
)

type RequestBody struct {
	UserID string
	Tags   []string
}

func SaveTags(db *sql.DB, body RequestBody) error {
	queryString := "INSERT INTO users (user_id, tags) values ($1,$2) ON CONFLICT (user_id) DO UPDATE SET tags = $3"
	_, err := db.Exec(queryString, body.UserID, pq.Array(body.Tags), pq.Array(body.Tags))
	if err != nil {
		fmt.Println("Error while saving tags to DB", err)
		return err
	}
	return nil
}
