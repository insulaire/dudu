package entity

import "github.com/google/uuid"

type User struct {
	Id   uint32
	Name string
}

func NewUser(name string) User {
	return User{
		Id:   uuid.New().ID(),
		Name: name,
	}
}
