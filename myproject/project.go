package main

import (
	"fmt"
	"github.com/jaishri241/mymodule"
)

func main() {
	message := mymodule.Greet("Alice")
	fmt.Println(message)
}
