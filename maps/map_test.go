package maps_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	u "github.com/rjNemo/underscore"
	m "github.com/rjNemo/underscore/maps"
)

func TestMap(t *testing.T) {
	scores := m.M[string, int]{
		"alice": 0,
		"bob":   10,
		"clara": 7,
		"david": 23,
	}

	hasWon := func(key string, value int) m.M[string, bool] { return m.M[string, bool]{key: value > 21} }
	want := m.M[string, bool]{
		"alice": false,
		"bob":   false,
		"clara": false,
		"david": true,
	}
	assert.Equal(t, want, m.Map(scores, hasWon))
}

func TestMapSlices(t *testing.T) {
	scores := []m.M[string, int]{
		{"score": 0},
		{"score": 10},
		{"score": 7},
		{"score": 23},
	}

	hasWon := func(s m.M[string, int]) bool { return s["score"] > 21 }
	want := []bool{false, false, false, true}
	assert.Equal(t, want, u.Map(scores, hasWon))
}
