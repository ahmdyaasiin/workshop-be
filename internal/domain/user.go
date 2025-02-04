package domain

import "github.com/google/uuid"

type UserEntity struct {
	ID       uuid.UUID
	Name     string
	Email    string
	Password string
	RoleID   int
}

type UserRequestRegister struct {
	Name     string
	Email    string
	Password string
}

type RoleEntity struct {
	ID   int
	Name string
}
