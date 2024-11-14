package server

import "github.com/google/uuid"

type Client struct {
	Address string
	UUID uuid.UUID
}