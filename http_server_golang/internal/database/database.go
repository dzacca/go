package database

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
)

// Client

type Client struct {
	path string
}

func NewClient(path string) Client {
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
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UserEmail string    `json:"userEmail"`
	Text      string    `json:"text"`
}

// EnsureDB creates the db file if it doesn't exist
func (c Client) EnsureDB() error {
	_, err := os.ReadFile(c.path)
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
	err = os.WriteFile(c.path, dat, 0644)
	if err != nil {
		return err
	}
	return nil
}

// updateDB
func (c Client) updateDB(db databaseSchema) error {
	dat, err := json.Marshal(db)
	if err != nil {
		return err
	}
	err = os.WriteFile(c.path, dat, 0644)
	if err != nil {
		return err
	}

	return nil
}

// readDB
func (c Client) readDB() (databaseSchema, error) {

	data, err := os.ReadFile(c.path)
	if err != nil {
		log.Fatal(err)
	}

	db := databaseSchema{}
	err = json.Unmarshal(data, &db)
	if err != nil {
		log.Fatal(err)
	}
	return db, nil
}

// CreateUser
func (c Client) CreateUser(email, password, name string, age int) (User, error) {
	db, err := c.readDB()
	if err != nil {
		log.Fatal(err)
	}

	user := User{
		CreatedAt: time.Now().UTC(),
		Email:     email,
		Password:  password,
		Name:      name,
		Age:       age,
	}

	db.Users[email] = user

	err = c.updateDB(db)
	if err != nil {
		return user, err
	}
	return user, nil
}

// UpdateUser
func (c Client) UpdateUser(email, password, name string, age int) (User, error) {
	db, err := c.readDB()
	if err != nil {
		log.Fatal(err)
	}

	user := User{
		Email:    email,
		Password: password,
		Name:     name,
		Age:      age,
	}

	_, ok := db.Users[email]
	if !ok {
		return user, errors.New("user doesn't exist")
	}

	db.Users[email] = user

	err = c.updateDB(db)
	if err != nil {
		u := User{}
		return u, err
	}
	return user, nil

}

// GetUser
func (c Client) GetUser(email string) (User, error) {
	db, err := c.readDB()
	if err != nil {
		u := User{}
		return u, errors.New("User not found")
	}
	if _, ok := db.Users[email]; !ok {
		u := User{}
		return u, errors.New("User not found")
	}
	return db.Users[email], nil
}

// DeleteUser
func (c Client) DeleteUser(email string) error {
	db, err := c.readDB()
	if err != nil {
		return err
	}
	delete(db.Users, email)
	err = c.updateDB(db)
	if err != nil {
		return err
	}

	return nil
}

// Posts
//CreatePost
func (c Client) CreatePost(userEmail, text string) (Post, error) {
	db, err := c.readDB()

	post := Post{}
	if err != nil {
		return post, errors.New("can't open the DB")
	}
	if _, ok := db.Users[userEmail]; !ok {
		return post, errors.New("User not found")
	}

	id := uuid.New().String()
	post = Post{
		CreatedAt: time.Now().UTC(),
		ID:        id,
		UserEmail: userEmail,
		Text:      text}

	db.Posts[id] = post

	err = c.updateDB(db)
	if err != nil {
		return post, errors.New("Can't update the DB")
	}
	return post, nil
}

// GetPosts
func (c Client) GetPosts(userEmail string) ([]Post, error) {
	db, err := c.readDB()
	if err != nil {
		p := []Post{}
		return p, errors.New("Can't open the DB")
	}

	posts := make([]Post, 0)
	for _, post := range db.Posts {
		if post.UserEmail == userEmail {
			posts = append(posts, post)
		}
	}

	return posts, nil
}

// DeletePost
func (c Client) DeletePost(id string) error {
	db, err := c.readDB()
	if err != nil {
		return errors.New("Can't open the DB")
	}
	if _, ok := db.Posts[id]; !ok {
		return errors.New("Post not found")
	}
	delete(db.Posts, id)

	err = c.updateDB(db)
	if err != nil {
		return err
	}
	return nil
}