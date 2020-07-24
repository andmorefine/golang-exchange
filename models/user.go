package models

import (
	"database/sql"
	"fmt"
	"time"
)

// DbConnection *sql.DB
var DbConnection *sql.DB

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
	cmd := fmt.Sprintf("INSERT INTO %s (uuid, name, delete_flag, created_at, updated_at) VALUES (?, ?, ?, ?, ?)", c.TableName())
	_, err := DbConnection.Exec(cmd, c.UUID, c.Name, c.DeleteFlag, c.CreatedAt.Format(time.RFC3339), c.UpdatedAt.Format(time.RFC3339))
	if err != nil {
		return err
	}
	return err
}

// Save User.Save
func (c *User) Save() error {
	cmd := fmt.Sprintf("UPDATE %s SET uuid = ?, name = ?, delete_flag = ?, created_at = ?, updated_at = ? WHERE name = ?", c.TableName())
	_, err := DbConnection.Exec(cmd, c.UUID, c.Name, c.DeleteFlag, c.CreatedAt.Format(time.RFC3339), c.UpdatedAt.Format(time.RFC3339), c.Name)
	if err != nil {
		return err
	}
	return err
}

// GetUser User.GetUser
func (c User) GetUser(productCode string, duration time.Duration, dateTime time.Time) *User {
	tableName := c.TableName()
	cmd := fmt.Sprintf("SELECT uuid, name, delete_flag, created_at, updated_at FROM %s WHERE name = ?", tableName)
	row := DbConnection.QueryRow(cmd, dateTime.Format(time.RFC3339))
	var user User
	err := row.Scan(&user.UUID, &user.Name, &user.DeleteFlag, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil
	}
	return NewUser(user.UUID, user.Name, user.DeleteFlag, user.CreatedAt, user.UpdatedAt)
}
