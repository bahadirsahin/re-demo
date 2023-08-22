package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSmallestPack(t *testing.T) {
	// case 1
	assert.Equal(t,
		GetSmallestPack(1, []int{250, 500, 1000, 2000, 5000}),
		250,
	)

	// case 2
	assert.Equal(t,
		GetSmallestPack(250, []int{250, 500, 1000, 2000, 5000}),
		250,
	)

	// case 3
	assert.Equal(t,
		GetSmallestPack(251, []int{250, 500, 1000, 2000, 5000}),
		500,
	)

	// case 4
	assert.Equal(t,
		GetSmallestPack(500, []int{250, 500, 1000, 2000, 5000}),
		500,
	)

	// case 5
	assert.Equal(t,
		GetSmallestPack(501, []int{250, 500, 1000, 2000, 5000}),
		1000,
	)

	// case 6
	assert.Equal(t,
		GetSmallestPack(1000, []int{250, 500, 1000, 2000, 5000}),
		1000,
	)

	// case 7
	assert.Equal(t,
		GetSmallestPack(1001, []int{250, 500, 1000, 2000, 5000}),
		2000,
	)

	// case 8
	assert.Equal(t,
		GetSmallestPack(2000, []int{250, 500, 1000, 2000, 5000}),
		2000,
	)

	// case 9
	assert.Equal(t,
		GetSmallestPack(2001, []int{250, 500, 1000, 2000, 5000}),
		5000,
	)

	// case 10
	assert.Equal(t,
		GetSmallestPack(5000, []int{250, 500, 1000, 2000, 5000}),
		5000,
	)
}
