package requests

type UserRoleCreateRequest struct {
	UserId int64 `json:"userId"`
	RoleId int64 `json:"roleId"`
}
