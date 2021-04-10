package main
//entry point

import (
	"fmt"
	"github.com/LivePurposefully/go-project-structure/pkg/maths"
)

func main() {
	fmt.Println("Hello World!")

	fmt.Println(maths.Add(2,3))
	fmt.Println(maths.Sub(2,3))
}