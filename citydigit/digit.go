package citydigit

import (
	"module3/wordz"
	. "module3/wordz"
)

func Digit() string {

	Words = []string{"One", "Two", "Three", "Four", "Five"}
	return wordz.Random()

}
