package entities

type Room struct {
	Id         string      `json:"id"`
	HostEmail  string      `json:"hostEmail"`
	TimeToLive int         `json:"timeToLive"` // In minutes
	Prompt     Prompt      `json:"prompt"`
	Comments   CommentList `json:"comment"`
	CreatedOn  string      `json:"createdOn"`
}

// tODO: Gets prompt from create request

type RoomList []Room
