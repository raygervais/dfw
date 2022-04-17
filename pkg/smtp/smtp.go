package smtp

import (
	"fmt"
	"net/smtp"
	"time"

	"github.com/raygervais/dfw/pkg/db"
)

type SMTP struct {
	Host string
	Port int
	Auth smtp.Auth
}

func Thread(db *db.Database) {
	client := SMTP{
		Host: "smtp.gmail.com",
		Port: 587,
		Auth: smtp.PlainAuth("", "", "", ""),
	}

	for {
		expiredRooms := db.ListExpiredRooms()
		for _, room := range expiredRooms {

			// Create a message which contains the contents of all comments in the room
			message := ""
			for _, comment := range room.Comments {
				message += comment.Comment + "\n"
			}

			// Send the message
			if err := client.Send("Room Comments for Prompt: "+room.Prompt.Text, []string{room.HostEmail},
				client.Message("Room Expired", room.HostEmail, "", message)); err != nil {
				fmt.Print(err)
			}

			db.RemoveRoom(room.Id)
		}

		time.Sleep(time.Minute)

	}
}

func (s *SMTP) Message(subject, to, from, msg string) []byte {
	return []byte("From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\n\n" +
		msg)
}

func (s *SMTP) Send(from string, to []string, msg []byte) error {
	return smtp.SendMail(s.Host+":"+string(s.Port), s.Auth, from, to, msg)
}
