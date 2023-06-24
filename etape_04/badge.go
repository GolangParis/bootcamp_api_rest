package main

import (
	"fmt"
	//"math/rand"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

/*

Améliorations souhaitables
=============================

Remplacer les log à base de fmt.Println par la lib log standard 
ou une lib de log structurés


*/

type User struct {
	ID     int
	Badges []Badge
}

type Badge struct {
	Name string
	Url  string
}


var (

	BadgesMap = map[int]*Badge {
		1 : &Badge{"go",    "https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white"},
		2 : &Badge{"steam", "https://img.shields.io/badge/Steam-000000?style=for-the-badge&logo=steam&logoColor=white"},
		3 : &Badge{"vue",   "https://img.shields.io/badge/Vue.js-35495E?style=for-the-badge&logo=vue.js&logoColor=4FC08D"},
		4 : &Badge{"docker","https://img.shields.io/badge/Docker-00005E?style=for-the-badge&logo=docker&logoColor=4FC08D"},
	}

	UserBadges = StorageInMemory{
		1: []int{ 1, 2 },
		2: []int{ 2, 3 },
		3: []int{ 4},
	}
)

func GetBadgeIDFromName(name string) (int, error) {

	for id, b := range BadgesMap {
		if b.Name == name {
			return id, nil
		}
	}
	return 0, fmt.Errorf("unknown badge name")
}

func GetUserBadges(c *gin.Context) {

	strID := c.Params.ByName("userid")
	vInt64, err := strconv.ParseInt(strID, 10, 32)

	if err != nil {
		fmt.Println("failed to parse userID: ", strID)
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	userID := int(vInt64)

	badges, ok := UserBadges[userID]
	if !ok {
		err := fmt.Errorf("Unknown userID: ", userID)
		c.JSON(http.StatusNotFound, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, badges)
}

// PostBadge ajoute un badge à un utilisateur dont l'ID est userID
func PostBadge(c *gin.Context) {
	userIDAsStr := c.Param("userid")

	valInt64, err := strconv.ParseInt(userIDAsStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	userID := int(valInt64)

	// Amélioration : dans le cas présent, on inspecte Userser.BadgeIDs
	// mais idéalement on devrait avoir un stockage utilisateurs distinct.
	userBadgeIDs, ok := UserBadges[userID]
	if !ok {
		c.JSON(404, gin.H{"error": fmt.Errorf("Unknown user: %v", userID)})
		return
	}
	
	var badge Badge

	if err := c.ShouldBindJSON(&badge); err != nil {
		c.JSON(500, gin.H{
			"status": "error",
			"reason": err.Error(),
		})
		return
	}

	fmt.Printf("Post badge : %v\n", badge)

	badgeID, err := GetBadgeIDFromName(badge.Name)
	if err != nil {
		c.JSON(404, gin.H{"error": err})
		return
	}

	userBadgeIDs = AddToSet(userBadgeIDs, badgeID)

	UserBadges[userID] = userBadgeIDs

	fmt.Println("User badges:", userBadgeIDs)

	c.JSON(200, gin.H{"status": "ok"})
}


func UpdateBadge(c *gin.Context) {
	userID := c.Param("userid")
	fmt.Println("userID", userID)

	var badge Badge

	if err := c.ShouldBindJSON(&badge); err != nil {
		c.JSON(500, gin.H{
			"status": "error",
			"reason": err.Error(),
		})
		return
	}

	fmt.Printf("Update badge : %v\n", badge)

	c.JSON(200, gin.H{"status": "ok"})
}


func DeleteBadge(c *gin.Context) {
	userIDAsStr := c.Param("userid")

	vInt64, err := strconv.ParseInt(userIDAsStr, 10, 64)

	if err != nil {
		fmt.Println("failed to parse userID: ", userIDAsStr)
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	userID := int(vInt64)

	_, ok := UserBadges[userID]
	if !ok {
		c.JSON(404, gin.H{"error": fmt.Errorf("Unknown user: %v", userID)})
		return
	}

	fmt.Println("userID", userID)

	var badge Badge
	
	if err := c.ShouldBindJSON(&badge); err != nil {
		c.JSON(500, gin.H{
			"status": "error",
			"reason": err.Error(),
		})
		return
	}

	fmt.Printf("Delete badge : %v\n", badge)

	// TODO remove badge ID

	c.JSON(200, gin.H{"status": "ok"})
}


func main() {
	vv := []int{ 1, 2}
	bb := RemoveFromSet(vv, 2)
	fmt.Println("bb: ", bb)

	ginRouter := gin.Default()

	ginRouter.GET("/api/badges/:userid", GetUserBadges)
	ginRouter.POST("/api/badge/:userid", PostBadge)
	ginRouter.PATCH("/api/badge/:userid", UpdateBadge)
	ginRouter.DELETE("/api/badge/:userid", DeleteBadge)

	ginRouter.Run(":8090")
}
