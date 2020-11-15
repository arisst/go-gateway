package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRandom(t *testing.T) {
	vendor := []Vendor{{"A", "url A", true}, {"B", "url B", true}, {"C", "url C", true}}

	getrandom := GetRandom(vendor)

	b := contains(vendor, *getrandom)

	assert.Equal(t, b, true)
}

func contains(arr []Vendor, str Vendor) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}
