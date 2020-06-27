package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	Echo()
}

func Echo() {
	fmt.Printf("Excuted command: %s\n", os.Args[0])
	fmt.Println("Arguments:")
	for i, v := range os.Args {
		fmt.Printf("arg[%v] = %s\n", i, v)
	}
	fmt.Println("Output:")
	fmt.Println(strings.Join(os.Args[1:], " "))
}
