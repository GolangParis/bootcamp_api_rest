package main

type User struct {
	ID string
	BadgeIDs []int
}

type Badge struct {
	Name string `json:"name"`
	URL string  `json:"url"`
}
