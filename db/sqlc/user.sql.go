// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: user.sql

package db

import (
	"context"
)

const createUser = `-- name: CreateUser :one

INSERT INTO users(user_name,user_password,user_phone)
VALUES
($1, crypt($2,gen_salt('bf')),$3)
RETURNING user_id, user_name, user_password, user_email, user_phone, user_token
`

type CreateUserParams struct {
	UserName  *string     `json:"user_name"`
	Crypt     interface{} `json:"crypt"`
	UserPhone *string     `json:"user_phone"`
}

// -- name: GetUserRoles :many
// select user_id,user_name,user_password,user_phone,user_token,role_id,role_name
// from users  join user_roles ur on user_id=usro_user_id
// join roles  on role_id = usro_role_id
// WHERE user_name = $1 and user_password = crypt($2, user_password);
func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (*User, error) {
	row := q.db.QueryRow(ctx, createUser, arg.UserName, arg.Crypt, arg.UserPhone)
	var i User
	err := row.Scan(
		&i.UserID,
		&i.UserName,
		&i.UserPassword,
		&i.UserEmail,
		&i.UserPhone,
		&i.UserToken,
	)
	return &i, err
}

const deleteToken = `-- name: DeleteToken :exec
UPDATE users SET user_token = '' WHERE user_token = $1
RETURNING user_id, user_name, user_password, user_email, user_phone, user_token
`

func (q *Queries) DeleteToken(ctx context.Context, userToken *string) error {
	_, err := q.db.Exec(ctx, deleteToken, userToken)
	return err
}

const findUserByPhone = `-- name: FindUserByPhone :one
select user_id,user_name,user_password,user_phone,user_token 
from users  
WHERE user_phone  = $1
`

type FindUserByPhoneRow struct {
	UserID       int32   `json:"user_id"`
	UserName     *string `json:"user_name"`
	UserPassword *string `json:"user_password"`
	UserPhone    *string `json:"user_phone"`
	UserToken    *string `json:"user_token"`
}

func (q *Queries) FindUserByPhone(ctx context.Context, userPhone *string) (*FindUserByPhoneRow, error) {
	row := q.db.QueryRow(ctx, findUserByPhone, userPhone)
	var i FindUserByPhoneRow
	err := row.Scan(
		&i.UserID,
		&i.UserName,
		&i.UserPassword,
		&i.UserPhone,
		&i.UserToken,
	)
	return &i, err
}

const findUserByUserPassword = `-- name: FindUserByUserPassword :one
select user_id,user_name,user_password,user_phone,user_token 
from users  
WHERE user_name = $1 and user_password = crypt($2, user_password)
`

type FindUserByUserPasswordParams struct {
	UserName *string     `json:"user_name"`
	Crypt    interface{} `json:"crypt"`
}

type FindUserByUserPasswordRow struct {
	UserID       int32   `json:"user_id"`
	UserName     *string `json:"user_name"`
	UserPassword *string `json:"user_password"`
	UserPhone    *string `json:"user_phone"`
	UserToken    *string `json:"user_token"`
}

func (q *Queries) FindUserByUserPassword(ctx context.Context, arg FindUserByUserPasswordParams) (*FindUserByUserPasswordRow, error) {
	row := q.db.QueryRow(ctx, findUserByUserPassword, arg.UserName, arg.Crypt)
	var i FindUserByUserPasswordRow
	err := row.Scan(
		&i.UserID,
		&i.UserName,
		&i.UserPassword,
		&i.UserPhone,
		&i.UserToken,
	)
	return &i, err
}

const findUserByUsername = `-- name: FindUserByUsername :one
select user_id,user_name,user_password,user_phone,user_token 
from users  
WHERE user_name = $1
`

type FindUserByUsernameRow struct {
	UserID       int32   `json:"user_id"`
	UserName     *string `json:"user_name"`
	UserPassword *string `json:"user_password"`
	UserPhone    *string `json:"user_phone"`
	UserToken    *string `json:"user_token"`
}

func (q *Queries) FindUserByUsername(ctx context.Context, userName *string) (*FindUserByUsernameRow, error) {
	row := q.db.QueryRow(ctx, findUserByUsername, userName)
	var i FindUserByUsernameRow
	err := row.Scan(
		&i.UserID,
		&i.UserName,
		&i.UserPassword,
		&i.UserPhone,
		&i.UserToken,
	)
	return &i, err
}

const updateToken = `-- name: UpdateToken :one
UPDATE users SET user_token = $1 WHERE user_id = $2
RETURNING user_id, user_name, user_password, user_email, user_phone, user_token
`

type UpdateTokenParams struct {
	UserToken *string `json:"user_token"`
	UserID    int32   `json:"user_id"`
}

func (q *Queries) UpdateToken(ctx context.Context, arg UpdateTokenParams) (*User, error) {
	row := q.db.QueryRow(ctx, updateToken, arg.UserToken, arg.UserID)
	var i User
	err := row.Scan(
		&i.UserID,
		&i.UserName,
		&i.UserPassword,
		&i.UserEmail,
		&i.UserPhone,
		&i.UserToken,
	)
	return &i, err
}

const updateUserName = `-- name: UpdateUserName :one
UPDATE users SET user_name = $1 WHERE user_id = $2
RETURNING user_id, user_name, user_password, user_email, user_phone, user_token
`

type UpdateUserNameParams struct {
	UserName *string `json:"user_name"`
	UserID   int32   `json:"user_id"`
}

func (q *Queries) UpdateUserName(ctx context.Context, arg UpdateUserNameParams) (*User, error) {
	row := q.db.QueryRow(ctx, updateUserName, arg.UserName, arg.UserID)
	var i User
	err := row.Scan(
		&i.UserID,
		&i.UserName,
		&i.UserPassword,
		&i.UserEmail,
		&i.UserPhone,
		&i.UserToken,
	)
	return &i, err
}

const updateUserPhone = `-- name: UpdateUserPhone :one
UPDATE users SET user_phone = $1 WHERE user_id = $2
RETURNING user_id, user_name, user_password, user_email, user_phone, user_token
`

type UpdateUserPhoneParams struct {
	UserPhone *string `json:"user_phone"`
	UserID    int32   `json:"user_id"`
}

func (q *Queries) UpdateUserPhone(ctx context.Context, arg UpdateUserPhoneParams) (*User, error) {
	row := q.db.QueryRow(ctx, updateUserPhone, arg.UserPhone, arg.UserID)
	var i User
	err := row.Scan(
		&i.UserID,
		&i.UserName,
		&i.UserPassword,
		&i.UserEmail,
		&i.UserPhone,
		&i.UserToken,
	)
	return &i, err
}
