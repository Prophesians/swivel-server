package repositories

import (
	"database/sql"
	"gitlab.com/Prophesians/swivel-server/config"
	"fmt"
)

type Repository struct {
	DB *sql.DB
}

func generateDBString(config *config.AppConfig) string{
	return fmt.Sprintf("host=%s dbname=%s user=%s password=%s search_path=%s sslmode=disable")
}

func GetRepository(config *config.AppConfig) (Repository, error) {
	connStr := generateDBString(config)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("Error while creating database connection:", err)
		return Repository{}, err
	}
	return Repository{db}, nil
}

func (r *Repository) Ping() error{
	return r.DB.Ping()
}