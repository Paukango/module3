package main

import (
	"fmt"
	"module3/citydigit"

	"github.com/huandu/xstrings"
)

func main() {

	city := citydigit.City()
	fmt.Println(city)
	digit := citydigit.Digit()
	fmt.Println(digit)
	fmt.Println(xstrings.Shuffle(city))

}
