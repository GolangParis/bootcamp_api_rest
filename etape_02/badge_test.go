package main

import (
	"testing"
)

func TestAddToSet(t *testing.T) {
	
	user := User{
		Id:     1234,
		Badges: []Badge{},
	}

	badge := Badge{
		Name: "go",
		Url:  "https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white",
	}

	newbadges := AddtoSet(user.Badges, badge)

	if len(newbadges) != 1 {
		t.Errorf("Unexpected length, got: %d", len(newbadges))
	}

}
