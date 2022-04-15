package entities

type Comment struct {
	Id           string `json:"id"`
	Comment      string `json:"comment"`
	OriginCookie string `json:"originCookie"` // IP Address to ensure only one submission per IP
}

type CommentList []Comment
