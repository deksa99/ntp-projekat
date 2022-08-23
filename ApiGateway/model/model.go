package model

type AuthenticatedUser struct {
	AccId  uint   `json:"AccId"`
	UserId uint   `json:"UserId"`
	Role   string `json:"Role"`
}

type Roles struct {
	Roles []string `json:"Roles"`
}
