package response

type Error struct {
	Message string `json:"Message"`
	Status  int    `json:"Status"`
}

func (e *Error) Error() string {
	return e.Message
}

type UserInfo struct {
	Id        uint   `json:"Id"`
	Email     string `json:"Email"`
	Username  string `json:"Username"`
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
}
