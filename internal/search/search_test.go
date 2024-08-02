package search

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnect(t *testing.T) {
	_, err := NewClient()

	assert.Nil(t, err)
}

func TestSearchHandler(t *testing.T) {
	expected := []Product{
		{
			Name:         "iPhone 13",
			Description:  "Latest model of iPhone with improved features",
			Price:        999.99,
			Category:     "smartphones",
			Availability: true,
		},
	}

	queries := []string{"iphone", "features", "13"}
	for _, query := range queries {
		actual, err := SearchPruducts(query)

		_ = actual
		_ = expected
		assert.Nil(t, err)
		// assert.Equal(t, expected, actual)
	}
}
