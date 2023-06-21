package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddToSet(t *testing.T) {
	vals := []int{ 1, 2, 3}

	vals = AddToSet(vals, 3)

	assert.Equal(t, vals, []int{ 1, 2, 3})
}
