package db

import (
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/raygervais/dfw/pkg/entities"
)

func (db *Database) AddRoom(room *entities.Room) error {
	if room.TimeToLive <= 0 {
		return errors.New("invalid room timeToLive provided") // Do not add room
	}
	if room.Id == "" {
		room.Id = strings.Split(uuid.New().String(), "-")[0]
	}

	if room.CreatedOn == "" {
		room.CreatedOn = time.Now().Format(time.RFC3339)
	}

	for _, r := range db.Rooms {
		if r.Id == room.Id {
			return errors.New("room already exists")
		}
	}

	db.Rooms = append(db.Rooms, *room)
	return nil
}

// Update Room
func (db *Database) UpdateRoom(room *entities.Room) {
	for i, r := range db.Rooms {
		if r.Id == room.Id {
			db.Rooms[i] = *room
			return
		}
	}
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

// Remove rooms which TimeToLive has expired
func (db *Database) RemoveExpiredRooms() {
	for i, room := range db.Rooms {
		// compare the current time to the createdOn time
		// if the difference is greater than the TimeToLive, remove the room
		createdOn, _ := time.Parse(time.RFC3339, room.CreatedOn)
		if time.Now().Sub(createdOn) > time.Duration(room.TimeToLive)*time.Minute {
			db.Rooms = append(db.Rooms[:i], db.Rooms[i+1:]...)
		}
	}
}
