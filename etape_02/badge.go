package main

import "fmt"

type User struct {
	Id     int
	Badges []Badge
}

type Badge struct {
	Name string
	Url  string
}

func main() {
	badge := Badge{
		Name: "go",
		Url:  "https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white",
	}

	user := User{
		Id:     1234,
		Badges: []Badge{},
	}

	user.Badges = AddtoSet(user.Badges, badge)

	fmt.Println("User ID=", user.Id, " Badges=", user.Badges)
}
