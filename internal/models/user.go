package models

import "time"

type User struct {
	ID       string    `json:"id,omitempty"`
	Email    string    `json:"email,omitempty"`
	Password string    `json:"password,omitempty"`
	Roles    []string  `json:"roles"`
	Created  time.Time `json:"created"`
	Updated  time.Time `json:"updated"`
}
