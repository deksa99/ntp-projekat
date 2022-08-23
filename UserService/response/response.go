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
	AccountId uint   `json:"AccountId"`
	Email     string `json:"Email"`
	Username  string `json:"Username"`
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
}

type WorkerInfo struct {
	WorkerID     uint   `json:"ID"`
	FirstName    string `json:"FirstName"`
	LastName     string `json:"LastName"`
	CarServiceID uint   `json:"CarServiceID"`
}

type Jwt struct {
	Token string `json:"Token"`
}

type Authentication struct {
	AccId  uint   `json:"AccId"`
	UserId uint   `json:"UserId"`
	Role   string `json:"Role"`
}
