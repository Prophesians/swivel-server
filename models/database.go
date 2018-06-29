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
	queryString := "INSERT INTO users (user_id, tags) values ($1,$2)"
	_, err := db.Exec(queryString, body.UserID, pq.Array(body.Tags))
	if err != nil {
		fmt.Println(err)
	}
	return nil
}
