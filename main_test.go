package main

import (
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func wordCount(value string) int {
	re := regexp.MustCompile(`[\S]+`)
	results := re.FindAllString(value, -1)
	return len(results)
}

func TestRandomName(t *testing.T) {

	assert.NotEqual(t, wordCount(fetchRandomName("")), wordCount(fetchRandomName(" test")), "They should not be equal")

}
