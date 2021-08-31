package main1

import (
	// "fmt"

	"github.com/dongri/phonenumber"
)

func NumberHp(nomorhp string) string {

	number := phonenumber.Parse(nomorhp, "id")
	// fmt.Println(number)

	return number

}
