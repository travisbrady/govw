package main

import (
	"fmt"

	vw "github.com/travisbrady/govw"
)

func main() {
	h := vw.Initialize("-q st --noconstant --invert_hash xor.w --quiet")
	vw.LearnString(h, "1 | 1:0 2:1 3:0")
	vw.LearnString(h, "1 | 1:0 2:1 3:0")
	vw.LearnString(h, "-1 | 1:0 2:0 3:0")
	vw.LearnString(h, "-1 | 1:0 2:0 3:0")
	vw.LearnString(h, "-1 | 1:1 2:1 3:0")
	vw.LearnString(h, "1 | 1:0 2:0 3:1")
	vw.LearnString(h, "1 | 1:0 2:0 3:1")
	vw.LearnString(h, "1 | 1:1 2:0 3:0")
	vw.LearnString(h, "1 | 1:1 2:0 3:0")
	vw.LearnString(h, "-1 | 1:0 2:0 3:0")
	vw.LearnString(h, "-1 | 1:0 2:0 3:0")
	vw.LearnString(h, "-1 | 1:0 2:0 3:0")
	vw.LearnString(h, "-1 | 1:1 2:1 3:1")
	vw.LearnString(h, "1 | 1:1 2:0 3:0")
	vw.LearnString(h, "1 | 1:0 2:1 3:0")
	vw.LearnString(h, "1 | 1:0 2:1 3:0")
	vw.LearnString(h, "1 | 1:0 2:1 3:0")
	vw.LearnString(h, "1 | 1:0 2:1 3:0")
	vw.LearnString(h, "1 | 1:0 2:1 3:0")
	vw.LearnString(h, "1 | 1:1 2:0 3:0")
	vw.LearnString(h, "1 | 1:0 2:1 3:0")
	vw.LearnString(h, "1 | 1:0 2:1 3:0")
	vw.LearnString(h, "1 | 1:0 2:1 3:0")
	vw.LearnString(h, "1 | 1:0 2:1 3:0")
	vw.LearnString(h, "-1 | 1:0 2:1 3:1")
	vw.LearnString(h, "1 | 1:0 2:1 3:0")
	vw.LearnString(h, "1 | 1:0 2:1 3:0")
	vw.LearnString(h, "1 | 1:0 2:1 3:0")
	vw.LearnString(h, "1 | 1:0 2:1 3:0")
	vw.LearnString(h, "1 | 1:0 2:1 3:0")
	vw.LearnString(h, "1 | 1:0 2:1 3:0")
	// Whoa, vw is really sensitive to order with few examples

	//h2 := vw.Initialize("-q st --noconstant -t -i xor.w")
	fmt.Printf("[SHOULD] = 0: %v\n", vw.LearnString(h, "-1 1 | 1:0 2:0 3:0"))
	fmt.Printf("[SHOULD] = 1: %v\n", vw.LearnString(h, "-1 1 | 1:0 2:1 3:0"))
	fmt.Printf("[SHOULD] = 1: %v\n", vw.LearnString(h, "-1 1 | 1:0 2:0 3:1"))
	//vw.Finish(h2)
	vw.Finish(h)
}
