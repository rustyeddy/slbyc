package main

import (
	"flag"
)

func main() {
	flag.Parse()

	todoist := newTodoist()
	todoist.GetProjects()

}
