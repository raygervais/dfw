package db

import (
	"errors"

	"github.com/google/uuid"
	"github.com/raygervais/dfw/pkg/entities"
)

func (db *Database) AddComment(id string, comment *entities.Comment) error {
	room := db.GetRoom(id)
	if room == nil {
		return errors.New("Room not found")
	}

	//Check if comment exists
	if comment.Id != "" {
		for _, c := range room.Comments {
			if c.Id == comment.Id {
				return errors.New("Comment already exists")
			}
		}
	} else {
		comment.Id = uuid.New().String()
	}

	room.Comments = append(room.Comments, *comment)

	db.UpdateRoom(room)

	return nil
}

// Get comments for a room
func (db *Database) GetComments(id string) entities.CommentList {
	room := db.GetRoom(id)
	if room != nil {
		return room.Comments
	}
	return nil
}
