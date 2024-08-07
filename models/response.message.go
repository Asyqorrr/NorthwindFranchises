package models

type CreateUserReq struct {
	UserName     string  `json:"user_name,omitempty" binding:"required"`
	UserPassword *string `json:"user_password,omitempty" binding:"required"`
}

type UserResponse struct {
	UserID       int32               `json:"user_id"`
	UserName     *string             `json:"user_name"`
	UserPassword *string             `json:"user_password"`
	UserPhone    *string             `json:"user_phone"`
	UserToken    *string             `json:"user_token"`
	Roles        []*UserRoleResponse `json:"roles"`
}

type UserRoleResponse struct {
	RoleID   int32   `json:"role_id"`
	RoleName *string `json:"role_name"`
}
