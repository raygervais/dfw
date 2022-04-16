package entities

type Room struct {
	Id         string      `json:"id"`
	HostEmail  string      `json:"hostEmail"`
	TimeToLive int         `json:"timeToLive"` // In minutes
	Prompts    PromptList  `json:"prompt"`
	Comments   CommentList `json:"comment"`
	CreatedOn  string      `json:"createdOn"`
}

type RoomList []Room
