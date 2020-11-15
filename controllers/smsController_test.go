package controllers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetVendor(t *testing.T) {
	vendor := getVendor("../vendor.toml")
	assert.NotEmpty(t, vendor)
}
