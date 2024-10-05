package apiLogin

type User struct {
	Id          int    `json:"id,omitempty"`
	Email       string `json:"email,omitempty"`
	Password    string `json:"password"`
	CurrentCode string `json:"currentCode,omitempty"`
	AppId       int    `json:"appId,omitempty"`
	GoogleToken string `json:"googleToken,omitempty"`
	RoleId      int    `json:"roleId,omitempty"`
	AccessToken string `json:"accessToken,omitempty"`
	// This token is used to authenticate an app at Lila login api
	AppToken    string       `json:"appToken,omitempty"`
	Username    string       `json:"userName,omitempty"`
	Permissions []Permission `json:"permissions,omitempty"`
	PhoneNumber string `json:"phoneNumber,omitempty"`
}

type Users struct {
	Users []User `json:"users,omitempty"`
}
