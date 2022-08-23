package request

type CreateUser struct {
	Email     string `json:"Email"`
	Username  string `json:"Username"`
	FirstName string `json:"FirstName"`
	LastName  string `json:"LastName"`
	Password  string `json:"Password"`
}

type Login struct {
	Username string `json:"Username"`
	Password string `json:"Password"`
}

type ChangePassword struct {
	Password    string `json:"Password"`
	NewPassword string `json:"NewPassword"`
}

type Authenticate struct {
	Roles []string `json:"Roles"`
}
