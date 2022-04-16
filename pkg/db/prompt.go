package db

import (
	"errors"

	"github.com/google/uuid"
	"github.com/raygervais/dfw/pkg/entities"
)

func (db *Database) AddPrompt(id string, prompt *entities.Prompt) error {
	room := db.GetRoom(id)
	if room == nil {
		return errors.New("Room not found")
	}

	if prompt.Id == "" {
		prompt.Id = uuid.New().String()
	}

	room.Prompt = *prompt
	db.UpdateRoom(room)
	return nil

}

func (db *Database) GetPrompts(id string) entities.Prompt {
	room := db.GetRoom(id)
	if room != nil {
		return room.Prompt
	}
	return entities.Prompt{}
}

func (db *Database) RemovePrompt(id string, promptId string) {
	room := db.GetRoom(id)
	if room != nil {
		room.Prompt = entities.Prompt{}
		db.UpdateRoom(room)
	}
}
