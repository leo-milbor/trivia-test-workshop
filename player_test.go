package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewPlayer(t *testing.T) {
	p := NewPlayer("leo")

	assert.NotNil(t, p)
}

// func Test_MovePlayer(t *testing.T){

// }
