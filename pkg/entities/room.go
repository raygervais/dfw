package entities

import "time"

type Room struct {
	Id          string      `json:"id"`
	HostEmail   string      `json:"hostEmail"`
	TimeNotLive time.Time   `json:"timeToLive"`
	Prompt      PromptList  `json:"prompt"`
	Comment     CommentList `json:"comment"`
}

type RoomList []Room
