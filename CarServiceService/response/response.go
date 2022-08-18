package response

type Error struct {
	Message string `json:"Message"`
	Status  int    `json:"Status"`
}

func (e *Error) Error() string {
	return e.Message
}
