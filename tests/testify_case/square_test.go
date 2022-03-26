package math

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSquare(t *testing.T) {
	input := 3
	expect := 8
	output := Square(input)
	assert.Equal(t, expect, output)
	assert.NotEqual(t, expect, output)
}