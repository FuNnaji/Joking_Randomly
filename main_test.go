package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandomName(t *testing.T) {

	assert.False(t, strings.Contains(fetchRandomName(randomNameURL), errorConstant))
	assert.True(t, strings.Contains(fetchRandomName(""), errorConstant))

}

func TestRandomJoke(t *testing.T) {

	assert.False(t, strings.Contains(fetchRandomJoke(randomJokeURL), errorConstant))
	assert.True(t, strings.Contains(fetchRandomJoke(""), errorConstant))

}
