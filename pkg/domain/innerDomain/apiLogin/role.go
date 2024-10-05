package apiLogin

type Role struct {
	Id          int          `json:"id,omitempty"`
	Role        string       `json:"role"`
	AppId       int          `json:"appId"`
	Permissions []Permission `json:"permissions,omitempty"`
}
