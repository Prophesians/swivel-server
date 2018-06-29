package repositories

import (
	"testing"
	config2 "gitlab.com/Prophesians/swivel-server/config"
	"github.com/stretchr/testify/assert"
)

func TestGenerateDBString(t *testing.T) {
	config := config2.AppConfig{
		Port: 1234,
		Host: "localhost",
		DBName: "dbname",
		User: "user",
		Password: "pass",
		Search_Path: "schema",
		DBPort: 5432,
	}

	dbString := GenerateDBString(&config)
	expected := "host=localhost dbname=dbname user=user password=pass search_path=schema port=5432 sslmode=disable"
	assert.Equal(t, expected, dbString, "should match the string")
}

