package main

import (
	"fmt"
	"math/rand"
)

type StorageInMemory map[int][]int


func (store *StorageInMemory) Connect() error {
	return nil
}

func (store StorageInMemory) CreateUser(user User) (User, error) {
	user.ID = rand.Int()

	fmt.Println("TODO: pass a slice of Badges IDs instead")
	// store[user.ID] = user.Badges

	return user, nil
}

func (store StorageInMemory) GetUsers() ([]User, error) {
	users := make([]User, 0)

	for ID, badgeIDs := range store {
		// TODO 
		bb := []Badge{}
		for _, badgeID := range badgeIDs {
			badge, ok := BadgesMap[badgeID]
			if ok {
				bb = append(bb, *badge)
			}
		}

		users = append(users, User{ID: ID, Badges: bb})
	}

	return users, nil
}

func (store StorageInMemory) GetUserBadges(userID int) ([]Badge, error) {
	badgeIDs, ok := store[userID]

	if !ok {
		return nil, fmt.Errorf("userId: %d does not exists", userID)
	}

	bb := []Badge{}
	for _, badgeID := range badgeIDs {
		badge, ok := BadgesMap[badgeID]
		if ok {
			bb = append(bb, *badge)
		}
	}

	return bb, nil
}

func (store StorageInMemory) AddUserBadge(userID int, badge Badge) error {
	// check if user exists
	badgeIDs, ok := UserBadges[userID]
	if !ok {
		return fmt.Errorf("userID: %v does not exists", userID)
	}

	newBadgeID, err := GetBadgeIDFromName(badge.Name)
	if err != nil {
		return fmt.Errorf("badge name unknown : %v", badge.Name)
	}

	store[userID] = AddToSet(badgeIDs, newBadgeID)

	return nil
}

func (store StorageInMemory) UpdteUserBadge(userID int, badge Badge) error {
	// check if user exists
	badges, err := store.GetUserBadges(userID)

	if err != nil {
		return err
	}

	for i := range badges {
		if badges[i].Name == badge.Name {
			badges[i] = badge
			return nil
		}
	}

	return fmt.Errorf("Badge Name=%s does not exists cannot update", badge.Name)
}

func (store StorageInMemory) DeleteUserBadge(userID int, badgeID int) error {
	badgeIDs, ok := UserBadges[userID]
	if !ok {
		return fmt.Errorf("userID: %v does not exists", userID)
	}

	store[userID] = RemoveFromSet(badgeIDs, badgeID)

	return nil
}

