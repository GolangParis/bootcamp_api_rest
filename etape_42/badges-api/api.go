package main

import (
	"fmt"

	"net/http"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {

	// Utilise la db basique en mémoire (cf. UserMap)
	// TODO à remplacer par un appel à une fonction de base de donnée

	numUsers := len(UsersMap)

	uu := make([]*User, 0, numUsers)

	for _, u := range UsersMap {
		uu = append(uu, u)
	}

	c.JSON(http.StatusOK, gin.H{ "users" : uu})
}

func GetUserBadges(c *gin.Context) {
	userID := c.Param("userid")

	// Utilise la db basique en mémoire (cf. UserMap)
	// TODO à remplacer par un appel à une fonction de base de donnée
	fmt.Println("userID", userID)

	user, ok := UsersMap[userID]
	if !ok {
		c.JSON(404, nil)
		return
	}

	bb := []*Badge{}

	for _, bid := range user.BadgeIDs {
		if b, ok := BadgesMap[bid]; ok {
			bb = append(bb, b)
		} else {
			fmt.Println("GetUserBadges: got unknown badge ID:", bid)
		}
	}

	c.JSON(http.StatusOK, gin.H{ "badges" : bb})
}

func GetBadgeIDFromName(name string) (int, error) {

	// Utilise la db basique en mémoire (cf. UserMap)
	// TODO à remplacer par un appel à une fonction de base de donnée

	for id, b := range BadgesMap {
		if b.Name == name {
			return id, nil
		}
	}
	return 0, fmt.Errorf("unknown badge name")
}

func PostBadge(c *gin.Context) {
	userID := c.Param("userid")

	fmt.Println("userID", userID)

	// Utilise la db basique en mémoire (cf. UserMap)
	// TODO à remplacer par un appel à une fonction de base de donnée

	user, ok := UsersMap[userID]
	if !ok {
		c.JSON(404, gin.H{ "error" : fmt.Errorf("user ID not found")})
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

	id, err := GetBadgeIDFromName(badge.Name)
	if err != nil {
		c.JSON(404, gin.H{"error": err})
		return
	}

	//user.BadgeIDs = append(user.BadgeIDs, id)
	user.BadgeIDs = AddToSet(user.BadgeIDs, id)

	fmt.Println("User badges:", user.BadgeIDs)

	// if _, ok := data.Films[filmID]; !ok { // Bien noter le !ok !
	// 	data.Films[filmID] = json
	c.JSON(200, gin.H{"status": "ok"})
	// } else {
	// 	c.JSON(500, gin.H{"status": "error"})
	// }
}

func DeleteBadge(c *gin.Context) {
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

	fmt.Printf("Delete badge : %v\n", badge)

	// if _, ok := data.Films[filmID]; !ok { // Bien noter le !ok !
	// 	data.Films[filmID] = json
	c.JSON(200, gin.H{"status": "ok"})
	// } else {
	// 	c.JSON(500, gin.H{"status": "error"})
	// }
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

	// if _, ok := data.Films[filmID]; !ok { // Bien noter le !ok !
	// 	data.Films[filmID] = json
	c.JSON(200, gin.H{"status": "ok"})
	// } else {
	// 	c.JSON(500, gin.H{"status": "error"})
	// }
}

func SetupRouter() *gin.Engine {

	// Gin est utilisé pour mettre en place l'API HTTP
	ginRouter := gin.Default()

	// On configure le routeur Gin pour servir le contenu du frontend
	// se trouvant dans le répertoire situé dans le répertoire public adjacent
	ginRouter.Use(static.Serve("/", static.LocalFile("../public", true)))

	ginRouter.GET("/api/users", GetUsers)
	ginRouter.GET("/api/badges/:userid", GetUserBadges)

	ginRouter.POST("/api/badge/:userid", PostBadge)
	ginRouter.DELETE("/api/badge/:userid", DeleteBadge)
	ginRouter.PATCH("/api/badge/:userid", UpdateBadge)

	return ginRouter
}
