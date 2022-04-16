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

	// Check if prompt already exists
	for _, p := range room.Prompts {
		if p.Id == prompt.Id {
			return errors.New("Prompt already exists")
		}
	}

	room.Prompts = append(room.Prompts, *prompt)
	db.UpdateRoom(room)
	return nil

}

func (db *Database) GetPrompts(id string) entities.PromptList {
	room := db.GetRoom(id)
	if room != nil {
		return room.Prompts
	}
	return nil
}

func (db *Database) RemovePrompt(id string, promptId string) {
	room := db.GetRoom(id)
	if room != nil {
		for i, prompt := range room.Prompts {
			if prompt.Id == promptId {
				room.Prompts = append(room.Prompts[:i], room.Prompts[i+1:]...)
				db.UpdateRoom(room)
				return
			}
		}
	}
}
