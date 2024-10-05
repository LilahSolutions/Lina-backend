package apiLogin

type Permission struct {
	Id              int    `json:"id"`
	Permission      string `json:"permission"`
	PermissionAlias string `json:"permissionAlias"`
}
