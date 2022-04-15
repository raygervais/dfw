package db

import (
	"github.com/raygervais/dfw/pkg/entities"
)

func (db *Database) AddRoom(room *entities.Room) {
	db.Rooms = append(db.Rooms, *room)
}

func (db *Database) GetRoom(id string) *entities.Room {
	for _, room := range db.Rooms {
		if room.Id == id {
			return &room
		}
	}
	return nil
}

func (db *Database) RemoveRoom(id string) {
	for i, room := range db.Rooms {
		if room.Id == id {
			db.Rooms = append(db.Rooms[:i], db.Rooms[i+1:]...)
			return
		}
	}
}
