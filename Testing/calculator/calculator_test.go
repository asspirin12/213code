package calculator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculatePlusTwo(t *testing.T) {
	input := 5
	expected := 7
	assert.Equal(t, expected, CalculatePlusTwo(input))
}
