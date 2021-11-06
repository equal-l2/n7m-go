package main

import (
	"fmt"
	"os"
    "github.com/equal-l2/n7m"
)

func main() {
	args := os.Args
	if len(args) > 1 {
		ret := n7m.N7m(args[1])
		fmt.Println(ret)
	}
}
