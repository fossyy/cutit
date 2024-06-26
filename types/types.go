package types

import "github.com/google/uuid"

type Message struct {
	Code    int
	Message string
}

type User struct {
	UserID        uuid.UUID
	Email         string
	Username      string
	Authenticated bool
}
