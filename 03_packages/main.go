package main

// import "fmt"
// import "math"

import (
	"fmt"
	"math"

	"github.com/kkhosla7/learn_on_the_go/03_packages/strutil"
)

func main() {
	fmt.Println("Here we go")
	fmt.Println("FLOOR: ", math.Floor(2.7))
	fmt.Println("CEIL: ", math.Ceil(2.7))
	fmt.Println("SQRT: ", math.Sqrt(16))
	fmt.Println("REVERSE: ", strutil.Reverse("olleH"))
}
