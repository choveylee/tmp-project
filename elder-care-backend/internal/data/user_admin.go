package data

type UserAdminData struct {
	UserId string `json:"user_id"`

	RoleId   string `json:"role_id"`
	RoleName string `json:"role_name"`

	Name   string `json:"name"`
	Mobile string `json:"mobile"`

	Status int `json:"status"`

	LoginAt string `json:"login_at"`

	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type ListUsersAdminRespData struct {
	List []*UserAdminData `json:"list"`

	Total int64 `json:"total"`
}

type CreateUserAdminRequest struct {
	RoleId string `json:"role_id"`

	Name   string `json:"name"`
	Mobile string `json:"mobile"`

	Password string `json:"password"`

	Status int `json:"status"`
}

type CreateUserAdminRespData struct {
	UserId string `json:"user_id"`
}

type UpdateUserAdminRequest struct {
	RoleId string `json:"role_id"`

	Name   string `json:"name"`
	Mobile string `json:"mobile"`

	Status int `json:"status"`
}

type UpdateUserPasswordAdminRequest struct {
	Password string `json:"password"`
}

type UpdateOwnPasswordAdminRequest struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}
