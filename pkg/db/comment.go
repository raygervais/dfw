package db

import "github.com/raygervais/dfw/pkg/entities"

func (db *Database) AddComment(id string, comment *entities.Comment) {
	room := db.GetRoom(id)
	if room != nil {
		room.Comment = append(room.Comment, *comment)
	}
}
