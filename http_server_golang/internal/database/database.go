package database

import (
	"encoding/json"
	"os"
	"time"
)

// Client

type Client struct {
	path string
}

func NewClient(path string) {
	return Client{path: path}
}

type databaseSchema struct {
	Users map[string]User `json:"users"`
	Posts map[string]Post `json:"posts"`
}

// User
type User struct {
	CreatedAt time.Time `json:"createdAt"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	Name      string    `json:"name"`
	Age       int       `json:"age"`
}

// Post
type Post struct {
	ID        string `json:"id"`
	CreatedAt string `json:"createdAt"`
	UserEmail string `json:"userEmail"`
	Text      string `json:"text"`
}

// EnsureDB creates the db file if it doesn't exist
func (c Client) EnsureDB() error {
	data, err := os.ReadFile(c.path)
	if err != nil {
		return c.createDB()
	}
	return err
}

// createDB
func (c Client) createDB() error {
	dat, err := json.Marshal(databaseSchema{
		Users: make(map[string]User),
		Posts: make(map[string]Post),
	})
	if err != nil {
		return err
	}
	err := os.WriteFile(c.path, dat, 0644)
	if err != nil {
		return err
	}
	return nil
}