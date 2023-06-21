package types

import (
)

type ID uint

type User struct {
	ID ID
	BadgeIDs []int
}

type Badge struct {
	Name string	
	URL string
}
