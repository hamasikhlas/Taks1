package main1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumberHp(t *testing.T) {

	test1 := NumberHp("6286548673212")
	assert.Equal(t, test1, "6286548673212", "sukses")

	test2 := NumberHp("086548673212")
	assert.Equal(t, test2, "6286548673212", "sukses")

	test3 := NumberHp("086548jhjshdu")
	assert.Equal(t, test3, "", "gagal")

	test4 := NumberHp("086548677565659877")
	assert.Equal(t, test4, "", "gagal")

	test5 := NumberHp("086546")
	assert.Equal(t, test5, "", "gagal")

}
