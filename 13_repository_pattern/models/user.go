package models

import "time"

// User represents a customer in DB with a backing table of users
type User struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	Email          string    `json:"email"`
	HashedPassword string    `json:"-"`
	CreatedAt      time.Time `json:"created_at"`
	Profile        Profile   `json:"profile"`
}

// Profile belongs to a User
type Profile struct {
	UserId  int       `json:"user_id"`
	Avatar  string    `json:"avatar"`
	Created time.Time `json:"created"`
}