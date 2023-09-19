package citydigit

import (
	"module3/wordz"
	. "module3/wordz"
)

func City() string {

	Words = []string{"Penza", "Saransk", "Tver", "Moskva", "Samara"}
	result := wordz.Random()
	return result

}
