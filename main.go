package main

import (
	"flag"
	"fmt"
)

func main() {
	flag.Parse()

	todoist := newTodoist()
	fmt.Printf("Token: %s\n", todoist.getToken())
}
