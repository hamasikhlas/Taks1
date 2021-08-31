package main1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	bener = true
	salah = false
	// sukses = "Email benar"
	gagal = "Email salah"
)

func TestEmail(t *testing.T) {
	a := Email("hamas@gmail.com")
	assert.Equal(t, a, bener, gagal)

	b := Email("hamas@yahoo.com")
	assert.Equal(t, b, bener, gagal)

	c := Email("hamasgmail.com")
	assert.Equal(t, c, salah, gagal)

	d := Email("hamasyahoo.")
	assert.Equal(t, d, salah, gagal)
}
