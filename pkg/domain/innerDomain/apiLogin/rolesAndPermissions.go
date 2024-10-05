package apiLogin

type RolesAndPermissions struct {
	AppId       int   `json:"appId"`
	Roles       []int `json:"roles"`
	Permissions []int `json:"permissions"`
}
