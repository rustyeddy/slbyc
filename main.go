package main

import (
	"flag"
	"fmt"
)

func main() {
	flag.Parse()

	todoist := newTodoist()
	// todoist.GetProjects()

	name := "SlBYC House Tasks"
	proj := todoist.GetProject(name)
	fmt.Printf("Project: %+v\n", proj)
}
