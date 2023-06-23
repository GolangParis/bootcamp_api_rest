package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID     int
	Badges []Badge
}

type Badge struct {
	Name string
	Url  string
}

type StorageInMemory map[int][]Badge

var (
	//UsersMap = make(map[string]*User)

	BadgesMap = StorageInMemory{
		1: {
			{
				Name: "go",
				Url:  "https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white",
			},
			{
				Name: "steam",
				Url:  "https://img.shields.io/badge/Steam-000000?style=for-the-badge&logo=steam&logoColor=white",
			},
		},
		2: {
			{
				Name: "vue",
				Url:  "https://img.shields.io/badge/Vue.js-35495E?style=for-the-badge&logo=vue.js&logoColor=4FC08D",
			},
		},
		3: {},
	}
)

func (store *StorageInMemory) Connect() error {
	return nil
}

func (store StorageInMemory) CreateUser(user User) (User, error) {
	user.ID = rand.Int()
	store[user.ID] = user.Badges

	return user, nil
}

func (store StorageInMemory) GetUsers() ([]User, error) {
	users := make([]User, 0)

	for uid, badges := range store {
		users = append(users, User{ID: uid, Badges: badges})
	}

	return users, nil
}

func (store StorageInMemory) GetUserBadges(userID int) ([]Badge, error) {
	badges, ok := store[userID]

	if !ok {
		return nil, fmt.Errorf("userId: %d does not exists", userID)
	}

	return badges, nil
}

func (store StorageInMemory) AddUserBadge(userID int, badge Badge) error {
	// check if user exists
	badges, err := store.GetUserBadges(userID)

	if err != nil {
		return err
	}

	store[userID] = AddtoSet(badges, badge)

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

func (store StorageInMemory) DeleteUserBadge(userID int, badgeName string) error {
	badges, err := store.GetUserBadges(userID)

	if err != nil {
		return err
	}

	for i := range badges {
		if badges[i].Name == badgeName {
			badges = append(badges[:i], badges[i+1:]...)
			store[userID] = badges

			return nil
		}
	}

	return fmt.Errorf("Badge Name=%s does not exists cannot update", badgeName)
}

func GetUserBadges(c *gin.Context) {

	id, err := strconv.ParseInt(c.Params.ByName("userid"), 10, 32)

	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	badges, err := BadgesMap.GetUserBadges(int(id))
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	c.JSON(http.StatusOK, badges)
}

func main() {
	ginRouter := gin.Default()

	ginRouter.GET("/api/badges/:userid", GetUserBadges)
	ginRouter.Run(":8090")
}
