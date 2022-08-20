package request

type Email struct {
	To      string `json:"To"`
	Subject string `json:"Subject"`
	Text    string `json:"Text"`
	Html    string `json:"Html"`
}
