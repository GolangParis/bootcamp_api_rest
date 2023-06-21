package main

import (
	"badges/types"
)

type DataSourceInterface interface {

	Connect()

	CreateUser(types.User) (types.ID, error)
		
	GetUsers() ([]types.User, error)

	GetUserBadges(userID types.ID) ([]types.Badge, error)
	AddUserBadge(userID types.ID, b types.Badge) (types.ID, error)
	UpdateUserBadge(userID types.ID, b types.Badge) error
	DeleteUserBadge(userID types.ID, badgeName string) error
}
