package model

type AuthenticatedUser struct {
	Id   uint   `json:"Id"`
	Role string `json:"Role"`
}

type Roles struct {
	Roles []string `json:"Roles"`
}
