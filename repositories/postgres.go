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
	DBString := "host=%s dbname=%s user=%s password=%s search_path=%s sslmode=disable"
	return fmt.Sprintf(DBString, config.Host, config.DBName, config.User, config.Password, config.Search_Path)
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