package main

import (
	"flag"
	"fmt"
)

func main() {
	flag.Parse()

	todoist := newTodoist()

	name := todoist.ProjectName
	proj := todoist.GetProject(name)
	fmt.Printf("Project: %+v\n", proj)
}
