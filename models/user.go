package models

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// DbConnection *sql.DB

// User struct
type User struct {
	UUID       int       `json:"uuid"`
	Name       string    `json:"name"`
	DeleteFlag bool      `json:"delete_flag"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// NewUser *User
func NewUser(uuid int, name string, deleteFlag bool, createdAt time.Time, updatedAt time.Time) *User {
	return &User{
		uuid,
		name,
		deleteFlag,
		createdAt,
		updatedAt,
	}
}

// TableName users
func (c *User) TableName() string {
	return "users"
}

// Create User.Create
func (c *User) Create() error {
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3336)/my_database")
	ins, err := db.Prepare("INSERT INTO users (uuid, name, delete_flag, created_at, updated_at) VALUES (?, ?, ?, ?, ?)")
	ins.Exec(c.UUID, c.Name, c.DeleteFlag, time.Now(), time.Now())
	if err != nil {
		log.Fatal(err)
		return err
	}
	return err
}

// Update User.Update
func (c *User) Update() error {
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3336)/my_database")
	upd, err := db.Prepare("UPDATE users SET uuid = ?, name = ?, delete_flag = ?, updated_at = ? WHERE name = ?")
	if err != nil {
		log.Fatal(err)
		return err
	}
	upd.Exec(c.UUID, c.Name, c.DeleteFlag, time.Now(), c.Name)
	return err
}
