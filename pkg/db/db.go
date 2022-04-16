package db

import "github.com/raygervais/dfw/pkg/entities"

type Database struct {
	Rooms entities.RoomList
}

type RoomRepository interface {
	AddRoom(room *entities.Room) error
	UpdateRoom(room *entities.Room)
	GetRoom(id string) *entities.Room
	RemoveRoom(id string)
}

type PromptRepository interface {
	AddPrompt(prompt *entities.Prompt) error
	UpdatePrompt(prompt *entities.Prompt)
	GetPrompt(id string) *entities.Prompt
	RemovePrompt(id string)
}

type CommentRepository interface {
	AddComment(id string, comment *entities.Comment) error
	GetComments(id string) entities.CommentList
}
