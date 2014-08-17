package main

import (
	"fmt"

	vw "github.com/travisbrady/go-vw"
)

func main() {
	h := vw.Initialize("-q st --noconstant")

	iters := 20
	for i := 0; i < iters; i++ {
		good_score := vw.LearnString(h, "1 | good")
		bad_score := vw.LearnString(h, "-1 | bad")
		fmt.Printf("I: %2d GoodScore: %6.3f BadScore: %6.3f\n", i, good_score, bad_score)
	}

	fmt.Printf("Should be good (1): %v\n", vw.LearnString(h, "-1 1 | good"))
	fmt.Printf("Should be bad (-1): %v\n", vw.LearnString(h, "-1 1 | bad"))

	vw.Finish(h)
}
