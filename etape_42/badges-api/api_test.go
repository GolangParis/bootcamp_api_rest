package main

import (
	"testing"

	"net/http"
	"net/http/httptest"

	"github.com/stretchr/testify/assert"
)

func TestGetUsers(t *testing.T) {
	router := SetupRouter()

	// Setup data source

	// Delete all users

	// ginRouter.GET("/api/users", GetUsers)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/users", nil)
	router.ServeHTTP(w, req)

	// should return 200
	assert.Equal(t, 200, w.Code) 

	// and an empty slice
	// assert.Equal(t, "pong", w.Body.String())
}

func TestGetUserBadges(t *testing.T) {
	// GET "/api/badges/:userid"

	// should return 200 + slice of badge IDs

	t.Errorf("TODO")
}

func TestAddUserBadge(t *testing.T) {
	// POST("/api/badge/:userid

	// should return 200 + badgeID

	t.Errorf("TODO")
}

func TestDeleteUserBadge(t *testing.T) {
	// Create user

	// DELETE "/api/badge/:userid"

	// should return 200

	t.Errorf("TODO")
}

func TestPatchUserBadge(t *testing.T) {
	// Create user

	// PATCH "/api/badge/:userid"

	// should return 200

	t.Errorf("TODO")
}
