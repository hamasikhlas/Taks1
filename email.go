package main1

import (
	"regexp"
)

func Email(email string) bool {

	emailRegex := regexp.MustCompile(`^[a-z0-9A-Z._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)

	return emailRegex.MatchString(email)
}

// func main() {
// 	fmt.Println(Email("hamas@gmail.com"))
// }
