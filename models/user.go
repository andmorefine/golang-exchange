package models

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var layout = "2006-01-02 15:04:05"

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
func (c *User) Create() (result sql.Result, err error) {
	db, err := sql.Open("mysql", "root:password@tcp(db_mysql5.7:3306)/my_database?parseTime=true&loc=Asia%2FTokyo")
	ins, err := db.Prepare("INSERT INTO users (uuid, name, delete_flag, created_at, updated_at) VALUES (?, ?, ?, ?, ?)")
	result, err = ins.Exec(c.UUID, c.Name, c.DeleteFlag, time.Now(), time.Now())
	if err != nil {
		log.Fatal(err)
	}
	return result, err
}

// Update User.Update
func (c *User) Update() error {
	db, err := sql.Open("mysql", "root:password@tcp(db_mysql5.7:3306)/my_database?parseTime=true&loc=Asia%2FTokyo")
	upd, err := db.Prepare("UPDATE users SET uuid = ?, name = ?, delete_flag = ?, updated_at = ? WHERE name = ?")
	if err != nil {
		log.Fatal(err)
		return err
	}
	upd.Exec(c.UUID, c.Name, c.DeleteFlag, time.Now(), c.Name)
	return err
}
