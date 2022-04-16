package entities

type Prompt struct {
	Id     string `json:"id"`
	Active bool   `json:"active"`
	Text   string `json:"text"`
}

type PromptList []Prompt
